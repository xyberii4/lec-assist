package twelvelabs

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	sdk "github.com/xyberii4/lec-assist/backend/pkg/twelvelabs"
	"go.uber.org/zap"
)

var (
	globalClient Client
	once         sync.Once
)

type APIError struct {
	StatusCode int
	Code       string
	Operation  string
	Message    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("TwelveLabs API error for %s (%d): %s - %s", e.Operation, e.StatusCode, e.Code, e.Message)
}

type Client interface {
	CreateIndex(ctx context.Context, req *sdk.CreateIndexRequest) (*sdk.InlineObject9, error)
	DeleteIndex(ctx context.Context, req *DeleteIndexRequest) error
	ListIndexes(ctx context.Context, query *ListIndexesQuery) (*sdk.InlineObject7, error)
	RetrieveIndex(ctx context.Context, req *RetrieveIndexRequest) (*sdk.Index, error)

	UploadVideo(ctx context.Context, req *UploadVideoRequest) (*sdk.InlineObject8, error)
	ListUploadTasks(ctx context.Context, query *ListUploadTasksQuery) (*sdk.InlineObject5, error)
	RetrieveUploadTask(ctx context.Context, req *RetrieveUploadTaskRequest) (*sdk.InlineObject6, error)
	ListVideos(ctx context.Context, query *ListVideosQuery) (*sdk.InlineObject3, error)

	Analyze(ctx context.Context, req *sdk.AnalyzeRequest) (*sdk.Analyze200Response, error)
	Gist(ctx context.Context, req *sdk.GistRequest) (*sdk.Gist, error)
	Summarise(ctx context.Context, req *sdk.SummarizeRequest) (*sdk.InlineObject19, error)
}

type twelvelabsClient struct {
	apiClient *sdk.APIClient
}

// Initialize TwelveLabs client only once
func InitClient(authKey string) {
	once.Do(func() {
		cfg := sdk.NewConfiguration()
		cfg.AddDefaultHeader("x-api-key", authKey)
		cfg.AddDefaultHeader("Content-Type", "application/json")

		apiClient := sdk.NewAPIClient(cfg)

		globalClient = &twelvelabsClient{
			apiClient: apiClient,
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

func (c *twelvelabsClient) getDefaultHeader(header string) string {
	return c.apiClient.GetConfig().DefaultHeader[header]
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
		apiErr.StatusCode = http.StatusBadGateway
		apiErr.Message = fmt.Sprintf("No HTTP response received: %s", err.Error())
		zap.L().Error("TwelveLabs API call failed (no HTTP response)",
			zap.String("operation", operation),
			zap.Error(err),
		)
	}
	return apiErr
}
