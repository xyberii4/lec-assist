package twelvelabs

import (
	"os"

	"go.uber.org/zap"
)

// -- Common Query Parameters for List Operations --

type commonListQueryParameters struct {
	Page       int32
	PageLimit  int32  // Max: 50
	SortBy     string // created_at/updated_at
	SortOption string // asc/desc
	CreatedAt  string // optional, RFC3339 format
	UpdatedAt  string // optional, RFC3339 format
}

type CommonListQueryOption func(*commonListQueryParameters)

func defaultCommonQueryParameters() commonListQueryParameters {
	return commonListQueryParameters{
		Page:       1,
		PageLimit:  10,
		SortBy:     "created_at",
		SortOption: "desc",
	}
}

func WithPage(page int32) func(*commonListQueryParameters) {
	return func(q *commonListQueryParameters) {
		q.Page = page
	}
}

func WithPageLimit(pageLimit int32) func(*commonListQueryParameters) {
	return func(q *commonListQueryParameters) {
		q.PageLimit = pageLimit
	}
}

func WithSortBy(sortBy string) func(*commonListQueryParameters) {
	return func(q *commonListQueryParameters) {
		q.SortBy = sortBy
	}
}

func WithSortOption(sortOption string) func(*commonListQueryParameters) {
	return func(q *commonListQueryParameters) {
		q.SortOption = sortOption
	}
}

func WithCreatedAt(createdAt string) func(*commonListQueryParameters) {
	return func(q *commonListQueryParameters) {
		q.CreatedAt = createdAt
	}
}

func WithUpdatedAt(updatedAt string) func(*commonListQueryParameters) {
	return func(q *commonListQueryParameters) {
		q.UpdatedAt = updatedAt
	}
}

// --Index Listing Query Parameters--

type listIndexesQuery struct {
	commonListQueryParameters

	IndexName    string
	ModelOptions string // audio/visual, multiple options must be comma-separated
	ModelFamily  string // marengo/pegasus
}

type IndexesParameter (func(*listIndexesQuery))

func DefaultListIndexesQuery() *listIndexesQuery {
	return &listIndexesQuery{
		commonListQueryParameters: defaultCommonQueryParameters(),
	}
}

func NewListIndexesQuery(opts ...interface{}) *listIndexesQuery {
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
	return func(q *listIndexesQuery) {
		q.IndexName = indexName
	}
}

func WithModelOptions(modelOptions string) IndexesParameter {
	return func(q *listIndexesQuery) {
		q.ModelOptions = modelOptions
	}
}

func WithModelFamily(modelFamily string) IndexesParameter {
	return func(q *listIndexesQuery) {
		q.ModelFamily = modelFamily
	}
}

// -- Upload Tasks Listing Query Parameters --

type listUploadTasksQuery struct {
	commonListQueryParameters

	IndexId  string
	Status   []string
	Filename string
	Duration float32
	Width    int32
	Height   int32
}

type listUploadTasksOption (func(*listUploadTasksQuery))

func DefaultListUploadTasksQuery() *listUploadTasksQuery {
	return &listUploadTasksQuery{
		commonListQueryParameters: defaultCommonQueryParameters(),
	}
}

func NewListUploadTasksQuery(opts ...interface{}) *listUploadTasksQuery {
	query := DefaultListUploadTasksQuery()
	for _, opt := range opts {
		switch o := opt.(type) {
		case listUploadTasksOption:
			o(query)
		case CommonListQueryOption:
			o(&query.commonListQueryParameters)
		default:
			zap.L().Warn("Unknown option type for NewListUploadTasksQuery", zap.Any("option", o))
		}
	}
	return query
}

func WithIndexId(indexId string) listUploadTasksOption {
	return func(q *listUploadTasksQuery) {
		q.IndexId = indexId
	}
}

func WithStatus(status []string) listUploadTasksOption {
	return func(q *listUploadTasksQuery) {
		q.Status = status
	}
}

func WithFilename(filename string) listUploadTasksOption {
	return func(q *listUploadTasksQuery) {
		q.Filename = filename
	}
}

func WithDuration(duration float32) listUploadTasksOption {
	return func(q *listUploadTasksQuery) {
		q.Duration = duration
	}
}

func WithWidth(width int32) listUploadTasksOption {
	return func(q *listUploadTasksQuery) {
		q.Width = width
	}
}

func WithHeight(height int32) listUploadTasksOption {
	return func(q *listUploadTasksQuery) {
		q.Height = height
	}
}

// -- Video Upload Request Parameters --

type uploadVideoRequest struct {
	IndexId   string
	VideoFile *os.File
	VideoUrl  string
}

type VideoOption (func(*uploadVideoRequest))

func DefaultUploadVideoRequest(indexId string) *uploadVideoRequest {
	return &uploadVideoRequest{
		IndexId:   indexId,
		VideoFile: nil,
		VideoUrl:  "",
	}
}

func NewUploadVideoRequest(indexId string, opts ...VideoOption) *uploadVideoRequest {
	req := DefaultUploadVideoRequest(indexId)
	for _, opt := range opts {
		opt(req)
	}
	return req
}

func WithVideoFile(videoFile *os.File) VideoOption {
	return func(req *uploadVideoRequest) {
		req.VideoFile = videoFile
	}
}

func WithVideoUrl(videoUrl string) VideoOption {
	return func(req *uploadVideoRequest) {
		req.VideoUrl = videoUrl
	}
}
