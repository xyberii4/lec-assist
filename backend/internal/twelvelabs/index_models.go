package twelvelabs

import "go.uber.org/zap"

// -- Index Deletion Request Body --

type DeleteIndexRequest struct {
	IndexId string `json:"index_id"`
}

// -- Index Retrieval Request Body --

type RetrieveIndexRequest struct {
	IndexId string `json:"index_id"`
}

// --Index Listing Query Parameters--

type ListIndexesQuery struct {
	commonListQueryParameters

	IndexName    string `json:"index_name,omitempty"`    // Filter by index name
	ModelOptions string `json:"model_options,omitempty"` // audio/visual, multiple options must be comma-separated
	ModelFamily  string `json:"model_family,omitempty"`  // marengo/pegasus
}

type IndexesParameter (func(*ListIndexesQuery))

func DefaultListIndexesQuery() *ListIndexesQuery {
	return &ListIndexesQuery{
		commonListQueryParameters: defaultCommonQueryParameters(),
	}
}

func NewListIndexesQuery(opts ...interface{}) *ListIndexesQuery {
	query := DefaultListIndexesQuery()
	for _, opt := range opts {
		switch o := opt.(type) {
		case IndexesParameter:
			o(query)
		case CommonListQueryOption:
			o(&query.commonListQueryParameters)
		default:
			zap.L().Warn("Unknown option type for NewListIndexesQuery", zap.Any("option", o))
		}
	}
	return query
}

func WithIndexName(indexName string) IndexesParameter {
	return func(q *ListIndexesQuery) {
		q.IndexName = indexName
	}
}

func WithModelOptions(modelOptions string) IndexesParameter {
	return func(q *ListIndexesQuery) {
		q.ModelOptions = modelOptions
	}
}

func WithModelFamily(modelFamily string) IndexesParameter {
	return func(q *ListIndexesQuery) {
		q.ModelFamily = modelFamily
	}
}

type IndexDetails struct {
	IndexId    string
	IndexName  string
	CreatedAt  string
	UpdatedAt  string
	ExpiresAt  string
	VideoCount int32
}
