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
	CreateIndex(ctx context.Context, indexName string, model []string) (*sdk.InlineObject9, error)
	DeleteIndex(ctx context.Context, indexId string) error
	RetrieveIndex(ctx context.Context, indexId string) (*sdk.Index, error)
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

// Creates index if it does not exist. model is list of models to use for the index (Marengo/Pegasus).
func (c *twelvelabsClient) CreateIndex(ctx context.Context, indexName string, models []string) (*sdk.InlineObject9, error) {
	indexModels := make([]sdk.CreateIndexRequestModelsInner, len(models))
	for i, m := range models {
		indexModels[i] = *sdk.NewCreateIndexRequestModelsInner(m, []string{"visual", "audio"})
	}

	// Create and execute request
	request := *sdk.NewCreateIndexRequest(indexName, indexModels)
	resp, r, err := c.apiClient.ManageIndexesAPI.CreateIndex(ctx).
		XApiKey(c.getDefaultHeader("x-api-key")).
		ContentType(c.getDefaultHeader("Content-Type")).
		CreateIndexRequest(request).
		Execute()
	if err != nil {
		return nil, c.handleHttpError(r, err, "CreateIndex")
	}
	defer r.Body.Close()

	zap.L().Info("Index created successfully",
		zap.String("index_id", *resp.Id),
		zap.String("index_name", indexName))
	return resp, nil
}

// Deletes an index by its ID
func (c *twelvelabsClient) DeleteIndex(ctx context.Context, indexId string) error {
	// Create and execute request
	r, err := c.apiClient.ManageIndexesAPI.DeleteIndex(ctx, indexId).
		XApiKey(c.getDefaultHeader("x-api-key")).
		ContentType(c.getDefaultHeader("Content-Type")).
		Execute()
	if err != nil {
		return c.handleHttpError(r, err, "DeleteIndex")
	}
	defer r.Body.Close()

	zap.L().Info("Index deleted successfully",
		zap.String("index_id", indexId))
	return nil
}

func (c *twelvelabsClient) RetrieveIndex(ctx context.Context, indexId string) (*sdk.Index, error) {
	resp, r, err := c.apiClient.ManageIndexesAPI.RetrieveIndex(ctx, indexId).
		XApiKey(c.getDefaultHeader("x-api-key")).
		ContentType(c.getDefaultHeader("Content-Type")).
		Execute()
	if err != nil {
		return nil, c.handleHttpError(r, err, "RetrieveIndex")
	}
	defer r.Body.Close()

	zap.L().Info("Index retrieved successfully",
		zap.String("index_id", indexId),
		zap.String("index_name", *resp.IndexName))

	return resp, nil
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
