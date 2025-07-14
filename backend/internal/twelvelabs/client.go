package twelvelabs

import (
	"context"
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
	// Handles
	if httpResp != nil {
		bodyBytes, readErr := io.ReadAll(httpResp.Body)
		if readErr != nil {
			zap.L().Error("Failed to read response body",
				zap.String("operation", operation),
				zap.Error(readErr),
			)
		}
		zap.L().Error("TwelveLabs API call failed",
			zap.String("operation", operation),
			zap.Int("status_code", httpResp.StatusCode),
			zap.String("response_body", string(bodyBytes)),
			zap.Error(err),
		)

		return fmt.Errorf("TwelveLabs API call failed for operation %s (status %d): %s - %w", operation, httpResp.StatusCode, string(bodyBytes), err)
	}
	zap.L().Error("TwelveLabs API call failed (no HTTP response)",
		zap.String("operation", operation),
		zap.Error(err),
	)
	return fmt.Errorf("TwelveLabs API call failed for operation %s: %w", operation, err)
}
