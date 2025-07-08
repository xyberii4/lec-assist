package twelvelabs

import (
	"context"

	sdk "github.com/xyberii4/lec-assist/backend/pkg/twelvelabs"
	"go.uber.org/zap"
)

type listIndexesQuery struct {
	SortBy       string // The field to sort on.
	SortOption   string // The sorting direction.
	IndexName    string // Filter by the name of an index.
	ModelFamily  string // Filter by the model family.
	ModelOptions string // Filter by the model options.
	CreatedAt    string // Filter indexes by the creation date and time.
	UpdatedAt    string // Filter indexes by the last update date and time.
	Page         int32  // A number that identifies the page to retrieve.
	PageLimit    int32  // The number of items to return on each page.
}

type Parameter (func(*listIndexesQuery))

func DefaultListIndexesQuery() *listIndexesQuery {
	return &listIndexesQuery{
		SortBy:       "created_at", // created_at/updated_at
		SortOption:   "desc",       // asc/desc
		IndexName:    "",           // optional
		ModelFamily:  "",           // optional, Marengo/Pegasus
		ModelOptions: "",           // optional, comma-separated list of model options
		CreatedAt:    "",           // optional, RFC3339 format
		UpdatedAt:    "",           // optional, RFC3339 format
		Page:         1,
		PageLimit:    10,
	}
}

func WithSortBy(sortBy string) Parameter {
	return func(q *listIndexesQuery) {
		q.SortBy = sortBy
	}
}

func WithSortOption(sortOption string) Parameter {
	return func(q *listIndexesQuery) {
		q.SortOption = sortOption
	}
}

func WithIndexName(indexName string) Parameter {
	return func(q *listIndexesQuery) {
		q.IndexName = indexName
	}
}

func WithModelFamily(modelFamily string) Parameter {
	return func(q *listIndexesQuery) {
		q.ModelFamily = modelFamily
	}
}

func WithModelOptions(modelOptions string) Parameter {
	return func(q *listIndexesQuery) {
		q.ModelOptions = modelOptions
	}
}

func WithCreatedAt(createdAt string) Parameter {
	return func(q *listIndexesQuery) {
		q.CreatedAt = createdAt
	}
}

func WithUpdatedAt(updatedAt string) Parameter {
	return func(q *listIndexesQuery) {
		q.UpdatedAt = updatedAt
	}
}

func WithPage(page int32) Parameter {
	return func(q *listIndexesQuery) {
		q.Page = page
	}
}

func WithPageLimit(pageLimit int32) Parameter {
	return func(q *listIndexesQuery) {
		q.PageLimit = pageLimit
	}
}

func NewListIndexesQuery(params ...Parameter) *listIndexesQuery {
	query := DefaultListIndexesQuery()
	for _, param := range params {
		param(query)
	}
	return query
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

func (c *twelvelabsClient) ListIndexes(ctx context.Context, query *listIndexesQuery) (*sdk.InlineObject7, error) {
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
