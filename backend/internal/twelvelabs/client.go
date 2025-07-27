package twelvelabs

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"

	sdk "github.com/xyberii4/lec-assist/backend/pkg/twelvelabs"
	"go.uber.org/zap"
)

const (
	createIndexOperation   = "CreateIndex"
	deleteIndexOperation   = "DeleteIndex"
	listIndexesOperation   = "ListIndexes"
	retrieveIndexOperation = "RetrieveIndex"

	uploadVideoOperation        = "UploadVideo"
	listUploadTasksOperation    = "ListUploadTasks"
	retrieveUploadTaskOperation = "RetrieveUploadTask"
	listVideosOperation         = "ListVideos"

	analyzeOperation   = "Analyze"
	gistOperation      = "Gist"
	summariseOperation = "Summarise"
)

var (
	globalClient Client
	once         sync.Once
)

type Client interface {
	CreateIndex(ctx context.Context, req *sdk.CreateIndexRequest) (string, error) // Returns index ID
	DeleteIndex(ctx context.Context, req *DeleteIndexRequest) error
	ListIndexes(ctx context.Context, query *ListIndexesQuery) ([]*IndexDetails, error)
	RetrieveIndex(ctx context.Context, req *RetrieveIndexRequest) (*IndexDetails, error)

	UploadVideo(ctx context.Context, req *UploadVideoRequest) (*TaskDetails, error) // Returns task ID and video ID
	ListUploadTasks(ctx context.Context, query *ListUploadTasksQuery) ([]*TaskDetails, error)
	RetrieveUploadTask(ctx context.Context, req *RetrieveUploadTaskRequest) (*TaskDetails, error)
	ListVideos(ctx context.Context, query *ListVideosQuery) ([]*VideoDetails, error)

	Analyze(ctx context.Context, req *sdk.AnalyzeRequest) (string, error) // Returns analysis result
	Gist(ctx context.Context, req *sdk.GistRequest) (*GistResult, error)
	Summarise(ctx context.Context, req *sdk.SummarizeRequest) (*SummariseResult, error)
}

type twelvelabsClient struct {
	apiClient       *sdk.APIClient
	mu              sync.Mutex
	rateLimitStates map[string]*RateLimit
}

// Initialize TwelveLabs client only once
func InitClient(authKey string) {
	once.Do(func() {
		cfg := sdk.NewConfiguration()
		cfg.AddDefaultHeader("x-api-key", authKey)
		cfg.AddDefaultHeader("Content-Type", "application/json")

		apiClient := sdk.NewAPIClient(cfg)

		globalClient = &twelvelabsClient{
			apiClient:       apiClient,
			rateLimitStates: make(map[string]*RateLimit),
		}

		zap.L().Info("TwelveLabs client initialized")
	})
}

func GetClient() Client {
	if globalClient == nil {
		zap.L().Fatal("TwelveLabs client not initialized. Call InitWrapper first.")
	}
	return globalClient
}

// Mock Client for tests
func NewTestClient(hc *http.Client, baseurl string, authKey string) Client {
	cfg := sdk.NewConfiguration()
	cfg.HTTPClient = hc
	cfg.Servers = []sdk.ServerConfiguration{
		{URL: baseurl},
	}
	cfg.AddDefaultHeader("x-api-key", authKey)
	cfg.AddDefaultHeader("Content-Type", "application/json")

	apiClient := sdk.NewAPIClient(cfg)

	return &twelvelabsClient{
		apiClient:       apiClient,
		mu:              sync.Mutex{},
		rateLimitStates: make(map[string]*RateLimit),
	}
}

func (c *twelvelabsClient) getDefaultHeader(header string) string {
	return c.apiClient.GetConfig().DefaultHeader[header]
}

// Get rate limit values from HTTP response
func (c *twelvelabsClient) getRateLimits(hr *http.Response) *RateLimit {
	rateLimit := &RateLimit{}

	// Convert header values to int
	parseIntHeader := func(header string) (int, bool) {
		valStr := hr.Header.Get(header)
		if valStr == "" {
			return 0, false
		}

		val, err := strconv.ParseInt(valStr, 10, 32)
		if err != nil {
			zap.L().Warn("Failed to parse rate limit header",
				zap.String("header", header),
				zap.String("value", valStr))

			return 0, false
		}
		return int(val), true
	}

	// Convert header values to int64
	parseInt64Header := func(header string) (int64, bool) {
		valStr := hr.Header.Get(header)
		if valStr == "" {
			return 0, false
		}

		val, err := strconv.ParseInt(valStr, 10, 64)
		if err != nil {
			zap.L().Warn("Failed to parse rate limit header",
				zap.String("header", header),
				zap.String("value", valStr))

			return 0, false
		}

		return val, true
	}

	// Extract rate limit values from headers
	if v, ok := parseIntHeader("X-RateLimit-Limit"); ok {
		rateLimit.Limit = v
	}

	if v, ok := parseIntHeader("X-RateLimit-Remaining"); ok {
		rateLimit.Remaining = v
	}

	if v, ok := parseIntHeader("X-RateLimit-Used"); ok {
		rateLimit.Used = v
	}

	if v, ok := parseInt64Header("X-RateLimit-Reset"); ok {
		rateLimit.Reset = v
	}

	return rateLimit
}

func (c *twelvelabsClient) updateRateLimitState(operation string, hr *http.Response) {
	newRateLimitState := c.getRateLimits(hr)
	if newRateLimitState == nil {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	currentRateLimitState, exists := c.rateLimitStates[operation]

	// Update the rate limit state only if it doesn't exist or if the new state is more restrictive
	if !exists || newRateLimitState.Reset > currentRateLimitState.Reset || newRateLimitState.Remaining < currentRateLimitState.Remaining {
		c.rateLimitStates[operation] = newRateLimitState
		zap.L().Debug("Updated rate limit state",
			zap.String("operation", operation),
			zap.Int("limit", newRateLimitState.Limit),
			zap.Int("remaining", newRateLimitState.Remaining))
	} else {
		zap.L().Debug("Rate limit state unchanged",
			zap.String("operation", operation))
	}
}

// check if rate limit is exceeded for the given operation
func (c *twelvelabsClient) checkRateLimitState(ctx context.Context, operation string) error {
	c.mu.Lock()
	currentRateLimitState, exists := c.rateLimitStates[operation]
	c.mu.Unlock()

	if exists && currentRateLimitState.Remaining <= 0 {
		// get time to reset
		resetTime := time.Unix(currentRateLimitState.Reset, 0).UTC()
		now := time.Now().UTC()

		// if before reset time, wait until reset
		if now.Before(resetTime) {
			sleepDuration := resetTime.Sub(now)
			zap.L().Warn("Rate limit exceeded, waiting for reset",
				zap.String("operation", operation),
				zap.Duration("time_to_reset", sleepDuration),
				zap.String("reset_time", resetTime.Format(time.RFC3339)))

			select {
			case <-time.After(sleepDuration):
				zap.L().Info("Rate limit reset, resuming operation",
					zap.String("operation", operation))

				c.mu.Lock()
				if rl, ok := c.rateLimitStates[operation]; ok {
					rl.Remaining = rl.Limit // Reset remaining count
					rl.Used = 0             // Reset used count
				}
				c.mu.Unlock()
			// Check if context is cancelled while waiting
			case <-ctx.Done():
				zap.L().Warn("Context cancelled while waiting for rate limit reset",
					zap.String("operation", operation))

				return ctx.Err()
			}

			// Already past reset time, reset rate limit state
		} else {
			zap.L().Info("Rate limit exceeded but reset time has passed, resuming operation",
				zap.String("operation", operation))

			c.mu.Lock()
			if rl, ok := c.rateLimitStates[operation]; ok {
				rl.Remaining = rl.Limit // Reset remaining count
				rl.Used = 0             // Reset used count
			}
			c.mu.Unlock()
		}
	}

	return nil
}

// Processes HTTP error responses from TwelveLabs API calls
func (c *twelvelabsClient) handleHttpError(httpResp *http.Response, err error, operation string) error {
	apiErr := &APIError{
		StatusCode: http.StatusInternalServerError,
		Operation:  operation,
		Message:    fmt.Sprintf("Unknown error occurred: %s", err.Error()),
	}

	if httpResp != nil {
		// twelvelabs error response structure
		type twelvelabsError struct {
			Code    string `json:"code"`
			Message string `json:"message"`
			DocsUrl string `json:"docs_url"`
		}

		apiErr.StatusCode = httpResp.StatusCode

		// Read the response body
		bodyBytes, readErr := io.ReadAll(httpResp.Body)
		if readErr != nil {
			zap.L().Error("Failed to read response body",
				zap.String("operation", operation),
				zap.Error(readErr),
			)
			apiErr.Message = fmt.Sprintf("Failed to read response body: %s", readErr.Error())
			apiErr.Code = err.Error()

			return apiErr
		}

		// Extract JSON error details
		var apiResponse twelvelabsError
		if jsonErr := json.Unmarshal(bodyBytes, &apiResponse); jsonErr != nil {
			if len(bodyBytes) > 0 {
				apiErr.Message = fmt.Sprintf("Failed to parse error response: %s", string(bodyBytes))
			}
		} else {
			apiErr.Code = apiResponse.Code
			apiErr.Message = apiResponse.Message
		}

		zap.L().Error("TwelveLabs API call failed",
			zap.String("operation", operation),
			zap.Int("status_code", httpResp.StatusCode),
			zap.String("response_body", string(bodyBytes)),
			zap.Error(err),
		)

	} else {
		// if no HTTP response
		apiErr.StatusCode = http.StatusServiceUnavailable
		apiErr.Message = fmt.Sprintf("No HTTP response received: %s", err.Error())
		zap.L().Error("TwelveLabs API call failed (no HTTP response)",
			zap.String("operation", operation),
			zap.Error(err),
		)
	}
	return apiErr
}
