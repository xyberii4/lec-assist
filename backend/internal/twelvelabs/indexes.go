package twelvelabs

import (
	"context"

	sdk "github.com/xyberii4/lec-assist/backend/pkg/twelvelabs"
	"go.uber.org/zap"
)

// Creates index if it does not exist.
func (c *twelvelabsClient) CreateIndex(ctx context.Context, req *sdk.CreateIndexRequest) (string, error) {
	resp, r, err := c.apiClient.ManageIndexesAPI.CreateIndex(ctx).
		XApiKey(c.getDefaultHeader("x-api-key")).
		ContentType(c.getDefaultHeader("Content-Type")).
		CreateIndexRequest(*req).
		Execute()

	if r != nil {
		defer r.Body.Close()
	}

	if err != nil {
		return "", c.handleHttpError(r, err, createIndexOperation)
	}

	zap.L().Info("Index created successfully",
		zap.String("index_id", *resp.Id),
		zap.String("index_name", req.IndexName))
	return *resp.Id, nil
}

// Deletes an index by its ID
func (c *twelvelabsClient) DeleteIndex(ctx context.Context, req *DeleteIndexRequest) error {
	// Create and execute request
	r, err := c.apiClient.ManageIndexesAPI.DeleteIndex(ctx, req.IndexId).
		XApiKey(c.getDefaultHeader("x-api-key")).
		ContentType(c.getDefaultHeader("Content-Type")).
		Execute()

	if r != nil {
		defer r.Body.Close()
	}

	if err != nil {
		return c.handleHttpError(r, err, deleteIndexOperation)
	}

	zap.L().Info("Index deleted successfully",
		zap.String("index_id", req.IndexId))
	return nil
}

func (c *twelvelabsClient) ListIndexes(ctx context.Context, query *ListIndexesQuery) ([]*IndexDetails, error) {
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

	if r != nil {
		defer r.Body.Close()
	}

	if err != nil {
		return nil, c.handleHttpError(r, err, listIndexesOperation)
	}

	// Extract index details from response
	indexDetails := resp.GetData()

	indexes := make([]*IndexDetails, 0, len(indexDetails))

	for _, index := range indexDetails {
		indexes = append(indexes, &IndexDetails{
			IndexId:    index.GetId(),
			IndexName:  index.GetIndexName(),
			CreatedAt:  index.GetCreatedAt(),
			UpdatedAt:  index.GetUpdatedAt(),
			ExpiresAt:  index.GetExpiresAt(),
			VideoCount: int32(index.GetVideoCount()),
		})
	}

	zap.L().Info("Indexes listed successfully",
		zap.Int("count", len(indexes)))

	return indexes, nil
}

func (c *twelvelabsClient) RetrieveIndex(ctx context.Context, req *RetrieveIndexRequest) (*IndexDetails, error) {
	resp, r, err := c.apiClient.ManageIndexesAPI.RetrieveIndex(ctx, req.IndexId).
		XApiKey(c.getDefaultHeader("x-api-key")).
		ContentType(c.getDefaultHeader("Content-Type")).
		Execute()

	if r != nil {
		defer r.Body.Close()
	}

	if err != nil {
		return nil, c.handleHttpError(r, err, retrieveIndexOperation)
	}

	index := &IndexDetails{
		IndexId:    resp.GetId(),
		IndexName:  resp.GetIndexName(),
		CreatedAt:  resp.GetCreatedAt(),
		UpdatedAt:  resp.GetUpdatedAt(),
		ExpiresAt:  resp.GetExpiresAt(),
		VideoCount: int32(resp.GetVideoCount()),
	}

	zap.L().Info("Index retrieved successfully",
		zap.String("index_id", index.IndexId),
		zap.String("index_name", index.IndexName))

	return index, nil
}
