package twelvelabs

import (
	"context"

	sdk "github.com/xyberii4/lec-assist/backend/pkg/twelvelabs"
	"go.uber.org/zap"
)

// Creates index if it does not exist.
func (c *twelvelabsClient) CreateIndex(ctx context.Context, req *sdk.CreateIndexRequest) (*sdk.InlineObject9, error) {
	resp, r, err := c.apiClient.ManageIndexesAPI.CreateIndex(ctx).
		XApiKey(c.getDefaultHeader("x-api-key")).
		ContentType(c.getDefaultHeader("Content-Type")).
		CreateIndexRequest(*req).
		Execute()
	if err != nil {
		return nil, c.handleHttpError(r, err, "CreateIndex")
	}
	defer r.Body.Close()

	zap.L().Info("Index created successfully",
		zap.String("index_id", *resp.Id),
		zap.String("index_name", req.IndexName))
	return resp, nil
}

// Deletes an index by its ID
func (c *twelvelabsClient) DeleteIndex(ctx context.Context, req *DeleteIndexRequest) error {
	// Create and execute request
	r, err := c.apiClient.ManageIndexesAPI.DeleteIndex(ctx, req.IndexId).
		XApiKey(c.getDefaultHeader("x-api-key")).
		ContentType(c.getDefaultHeader("Content-Type")).
		Execute()
	if err != nil {
		return c.handleHttpError(r, err, "DeleteIndex")
	}
	defer r.Body.Close()

	zap.L().Info("Index deleted successfully",
		zap.String("index_id", req.IndexId))
	return nil
}

func (c *twelvelabsClient) ListIndexes(ctx context.Context, query *ListIndexesQuery) (*sdk.InlineObject7, error) {
	req := c.apiClient.ManageIndexesAPI.ListIndexes(ctx).
		XApiKey(c.getDefaultHeader("x-api-key")).
		ContentType(c.getDefaultHeader("Content-Type")).
		Page(query.Page).
		PageLimit(query.PageLimit).
		SortBy(query.SortBy).
		SortOption(query.SortOption)

	if query.IndexName != "" {
		req = req.IndexName(query.IndexName)
	}
	if query.ModelFamily != "" {
		req = req.ModelFamily(query.ModelFamily)
	}
	if query.ModelOptions != "" {
		req = req.ModelOptions(query.ModelOptions)
	}
	if query.CreatedAt != "" {
		req = req.CreatedAt(query.CreatedAt)
	}
	if query.UpdatedAt != "" {
		req = req.UpdatedAt(query.UpdatedAt)
	}

	resp, r, err := req.Execute()
	if err != nil {
		return nil, c.handleHttpError(r, err, "ListIndexes")
	}
	defer r.Body.Close()

	zap.L().Info("Indexes listed successfully")

	return resp, nil
}

func (c *twelvelabsClient) RetrieveIndex(ctx context.Context, req *RetrieveIndexRequest) (*sdk.Index, error) {
	resp, r, err := c.apiClient.ManageIndexesAPI.RetrieveIndex(ctx, req.IndexId).
		XApiKey(c.getDefaultHeader("x-api-key")).
		ContentType(c.getDefaultHeader("Content-Type")).
		Execute()
	if err != nil {
		return nil, c.handleHttpError(r, err, "RetrieveIndex")
	}
	defer r.Body.Close()

	zap.L().Info("Index retrieved successfully",
		zap.String("index_id", req.IndexId),
		zap.String("index_name", *resp.IndexName))

	return resp, nil
}
