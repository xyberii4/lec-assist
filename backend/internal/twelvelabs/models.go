package twelvelabs

import (
	"os"

	sdk "github.com/xyberii4/lec-assist/backend/pkg/twelvelabs"
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

type ListUploadTasksOption (func(*listUploadTasksQuery))

func DefaultListUploadTasksQuery() *listUploadTasksQuery {
	return &listUploadTasksQuery{
		commonListQueryParameters: defaultCommonQueryParameters(),
	}
}

func NewListUploadTasksQuery(opts ...interface{}) *listUploadTasksQuery {
	query := DefaultListUploadTasksQuery()
	for _, opt := range opts {
		switch o := opt.(type) {
		case ListUploadTasksOption:
			o(query)
		case CommonListQueryOption:
			o(&query.commonListQueryParameters)
		default:
			zap.L().Warn("Unknown option type for ListUploadTasksQuery", zap.Any("option", o))
		}
	}
	return query
}

func WithListUploadTasksIndexId(indexId string) ListUploadTasksOption {
	return func(q *listUploadTasksQuery) {
		q.IndexId = indexId
	}
}

func WithListUploadTasksStatus(status []string) ListUploadTasksOption {
	return func(q *listUploadTasksQuery) {
		q.Status = status
	}
}

func WithListUploadTasksFilename(filename string) ListUploadTasksOption {
	return func(q *listUploadTasksQuery) {
		q.Filename = filename
	}
}

func WithListUploadTasksDuration(duration float32) ListUploadTasksOption {
	return func(q *listUploadTasksQuery) {
		q.Duration = duration
	}
}

func WithListUploadTasksWidth(width int32) ListUploadTasksOption {
	return func(q *listUploadTasksQuery) {
		q.Width = width
	}
}

func WithListUploadTasksHeight(height int32) ListUploadTasksOption {
	return func(q *listUploadTasksQuery) {
		q.Height = height
	}
}

// -- List Videos Query Parameters --
type listVideosQuery struct {
	commonListQueryParameters

	IndexId      string
	Filename     string
	Duration     float32
	Fps          float32
	Width        float32
	Height       int32
	Size         float32 // bytes
	UserMetadata map[string]sdk.ListVideosUserMetadataParameterValue
}

type ListVideosOption func(*listVideosQuery)

func DefaultListVideosQuery(indexId string) *listVideosQuery {
	return &listVideosQuery{
		commonListQueryParameters: defaultCommonQueryParameters(),
		IndexId:                   indexId,
	}
}

func NewListVideosQuery(indexId string, opts ...interface{}) *listVideosQuery {
	query := DefaultListVideosQuery(indexId)
	for _, opt := range opts {
		switch o := opt.(type) {
		case ListVideosOption:
			o(query)
		case CommonListQueryOption:
			o(&query.commonListQueryParameters)
		default:
			zap.L().Warn("Unknown option type for ListVideosQuery", zap.Any("option", o))
		}
	}
	return query
}

func WithListVideosFilename(filename string) ListVideosOption {
	return func(q *listVideosQuery) {
		q.Filename = filename
	}
}

func WithListVideosDuration(duration float32) ListVideosOption {
	return func(q *listVideosQuery) {
		q.Duration = duration
	}
}

func WithListVideosFps(fps float32) ListVideosOption {
	return func(q *listVideosQuery) {
		q.Fps = fps
	}
}

func WithListVideosWidth(width float32) ListVideosOption {
	return func(q *listVideosQuery) {
		q.Width = width
	}
}

func WithListVideosHeight(height int32) ListVideosOption {
	return func(q *listVideosQuery) {
		q.Height = height
	}
}

func WithListVideosSize(size float32) ListVideosOption {
	return func(q *listVideosQuery) {
		q.Size = size
	}
}

func WithListVideosUserMetadata(userMetadata map[string]sdk.ListVideosUserMetadataParameterValue) ListVideosOption {
	return func(q *listVideosQuery) {
		q.UserMetadata = userMetadata
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
