package twelvelabs

import (
	"os"

	sdk "github.com/xyberii4/lec-assist/backend/pkg/twelvelabs"
	"go.uber.org/zap"
)

// -- Retrieve Upload Task Request Body --

type RetrieveUploadTaskRequest struct {
	TaskId string `json:"task_id"`
}

// -- Video Upload Request Parameters --

type UploadVideoRequest struct {
	IndexId   string   `json:"index_id"`
	VideoFile *os.File `json:"-"`
	VideoUrl  string   `json:"video_url,omitempty"` // Optional, either video_file or video_url must be provided
}

type VideoOption (func(*UploadVideoRequest))

func DefaultUploadVideoRequest(indexId string) *UploadVideoRequest {
	return &UploadVideoRequest{
		IndexId:   indexId,
		VideoFile: nil,
		VideoUrl:  "",
	}
}

func NewUploadVideoRequest(indexId string, opts ...VideoOption) *UploadVideoRequest {
	req := DefaultUploadVideoRequest(indexId)
	for _, opt := range opts {
		opt(req)
	}
	return req
}

func WithVideoFile(videoFile *os.File) VideoOption {
	return func(req *UploadVideoRequest) {
		req.VideoFile = videoFile
	}
}

func WithVideoUrl(videoUrl string) VideoOption {
	return func(req *UploadVideoRequest) {
		req.VideoUrl = videoUrl
	}
}

type TaskDetails struct {
	TaskId    string
	VideoId   string
	CreatedAt string
	UpdatedAt string
	Status    string
	IndexId   string
}

// -- Upload Tasks Listing Query Parameters --

type ListUploadTasksQuery struct {
	commonListQueryParameters

	IndexId  string   `json:"index_id,omitempty"`
	Status   []string `json:"status,omitempty"`
	Filename string   `json:"filename,omitempty"`
	Duration float32  `json:"duration,omitempty"` // seconds
	Width    int32    `json:"width,omitempty"`    // pixels
	Height   int32    `json:"height,omitempty"`   // pixels
}

type ListUploadTasksOption (func(*ListUploadTasksQuery))

func DefaultListUploadTasksQuery() *ListUploadTasksQuery {
	return &ListUploadTasksQuery{
		commonListQueryParameters: defaultCommonQueryParameters(),
	}
}

func NewListUploadTasksQuery(opts ...interface{}) *ListUploadTasksQuery {
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
	return func(q *ListUploadTasksQuery) {
		q.IndexId = indexId
	}
}

func WithListUploadTasksStatus(status []string) ListUploadTasksOption {
	return func(q *ListUploadTasksQuery) {
		q.Status = status
	}
}

func WithListUploadTasksFilename(filename string) ListUploadTasksOption {
	return func(q *ListUploadTasksQuery) {
		q.Filename = filename
	}
}

func WithListUploadTasksDuration(duration float32) ListUploadTasksOption {
	return func(q *ListUploadTasksQuery) {
		q.Duration = duration
	}
}

func WithListUploadTasksWidth(width int32) ListUploadTasksOption {
	return func(q *ListUploadTasksQuery) {
		q.Width = width
	}
}

func WithListUploadTasksHeight(height int32) ListUploadTasksOption {
	return func(q *ListUploadTasksQuery) {
		q.Height = height
	}
}

type VideoDetails struct {
	VideoId   string
	CreatedAt string
	UpdatedAt string
	IndexId   string
	Filename  string
	Duration  float32
	Fps       float32
	Width     int32
	Height    int32
	Size      float32
}

// -- List Videos Query Parameters --
type ListVideosQuery struct {
	commonListQueryParameters

	IndexId      string  `json:"index_id",omitempty`
	Filename     string  `json:"filename,omitempty"`
	Duration     float32 `json:"duration,omitempty"` // seconds
	Fps          float32 `json:"fps,omitempty"`      // frames per second
	Width        float32 `json:"width,omitempty"`    // pixels
	Height       int32   `json:"height,omitempty"`   // pixels
	Size         float32 `json:"size,omitempty"`     // bytes
	UserMetadata map[string]sdk.ListVideosUserMetadataParameterValue
}

type ListVideosOption func(*ListVideosQuery)

func DefaultListVideosQuery(indexId string) *ListVideosQuery {
	return &ListVideosQuery{
		commonListQueryParameters: defaultCommonQueryParameters(),
		IndexId:                   indexId,
	}
}

func NewListVideosQuery(indexId string, opts ...interface{}) *ListVideosQuery {
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
	return func(q *ListVideosQuery) {
		q.Filename = filename
	}
}

func WithListVideosDuration(duration float32) ListVideosOption {
	return func(q *ListVideosQuery) {
		q.Duration = duration
	}
}

func WithListVideosFps(fps float32) ListVideosOption {
	return func(q *ListVideosQuery) {
		q.Fps = fps
	}
}

func WithListVideosWidth(width float32) ListVideosOption {
	return func(q *ListVideosQuery) {
		q.Width = width
	}
}

func WithListVideosHeight(height int32) ListVideosOption {
	return func(q *ListVideosQuery) {
		q.Height = height
	}
}

func WithListVideosSize(size float32) ListVideosOption {
	return func(q *ListVideosQuery) {
		q.Size = size
	}
}

func WithListVideosUserMetadata(userMetadata map[string]sdk.ListVideosUserMetadataParameterValue) ListVideosOption {
	return func(q *ListVideosQuery) {
		q.UserMetadata = userMetadata
	}
}
