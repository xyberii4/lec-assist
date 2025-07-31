package tlservice

import (
	"context"
	"errors"
	"time"

	"github.com/xyberii4/lec-assist/backend/internal/twelvelabs"
	"go.uber.org/zap"
)

func (s *service) UploadVideo(ctx context.Context, req *twelvelabs.UploadVideoRequest) (*twelvelabs.TaskDetails, error) {
	zap.L().Debug("Service: Attempting to upload video...", zap.Any("user_id", ctx.Value("userId")))

	if req == nil {
		return nil, NewServiceError(ErrCodeInvalidArg, "upload video request cannot be nil", nil)
	}
	// TODO: validate index ID
	if req.IndexId == "" {
		return nil, NewServiceError(ErrCodeInvalidArg, "index ID cannot be empty", nil)
	}

	if req.VideoFile == nil && req.VideoUrl == "" {
		return nil, NewServiceError(ErrCodeInvalidArg, "either video file or video URL must be provided", nil)
	}

	if req.VideoFile != nil && req.VideoUrl != "" {
		return nil, NewServiceError(ErrCodeInvalidArg, "only one of video file or video URL can be provided", nil)
	}

	resp, err := s.client.UploadVideo(ctx, req)
	if err != nil {
		zap.L().Error("Service: Failed to upload video",
			zap.Any("user_id", ctx.Value("userId")),
			zap.String("index_id", req.IndexId))
		var apiErr *twelvelabs.APIError
		if errors.As(err, &apiErr) {
			switch apiErr.Code {
			case "RateLimitExceeded":
				return nil, NewServiceError(ErrCodeRateLimitExceeded, "rate limit exceeded for video upload", err)
			default:
				return nil, NewServiceError(ErrCodeInvalidArg, "failed to upload video: "+apiErr.Message, err)
			}
		}

		return nil, NewServiceError(ErrCodeExternalAPIError, "failed to upload video", err)
	}

	return resp, nil
}

func (s *service) ListUploads(ctx context.Context, req *ListUploadsQuery) ([]*UploadDetails, error) {
	zap.L().Debug("Service: Attempting to list video uploads...", zap.Any("user_id", ctx.Value("userId")))

	if req == nil || req.CommonListQuery == nil {
		return nil, NewServiceError(ErrCodeInvalidArg, "request cannot be nil", nil)
	}

	queryOpts, err := s.parseCommonListQuery(req.CommonListQuery)
	if err != nil {
		return nil, err
	}

	// TODO: validate index ID
	if req.IndexId != "" {
		queryOpts = append(queryOpts, twelvelabs.WithListUploadTasksIndexId(req.IndexId))
	}

	if req.Status != nil && len(req.Status) > 0 {
		queryOpts = append(queryOpts, twelvelabs.WithListUploadTasksStatus(req.Status))
	}

	if req.Filename != "" {
		queryOpts = append(queryOpts, twelvelabs.WithListUploadTasksFilename(req.Filename))
	}

	if req.Duration > 0 {
		queryOpts = append(queryOpts, twelvelabs.WithListUploadTasksDuration(req.Duration))
	}

	if req.Width > 0 {
		queryOpts = append(queryOpts, twelvelabs.WithListUploadTasksWidth(req.Width))
	}

	if req.Height > 0 {
		queryOpts = append(queryOpts, twelvelabs.WithListUploadTasksHeight(req.Height))
	}

	query := twelvelabs.NewListUploadTasksQuery(queryOpts...)

	resp, err := s.client.ListUploadTasks(ctx, query)
	if err != nil {
		zap.L().Error("Service: Failed to list video uploads",
			zap.Any("user_id", ctx.Value("userId")),
			zap.Error(err))

		var apiErr *twelvelabs.APIError
		if errors.As(err, &apiErr) {
			return nil, NewServiceError(ErrCodeExternalAPIError, apiErr.Message, err)
		}

		return nil, NewServiceError(ErrCodeExternalAPIError, "failed to list video uploads", err)
	}

	var uploads []*UploadDetails

	for _, task := range resp {
		upload := s.parseTaskDetails(task)
		uploads = append(uploads, upload)
	}

	return uploads, nil
}

func (s *service) RetrieveUpload(ctx context.Context, taskId string) (*UploadDetails, error) {
	zap.L().Debug("Service: Attempting to retrieve video upload details...", zap.Any("user_id", ctx.Value("userId")), zap.String("task_id", taskId))

	// TODO: validate task ID
	if taskId == "" {
		return nil, NewServiceError(ErrCodeInvalidArg, "task ID cannot be empty", nil)
	}

	req := &twelvelabs.RetrieveUploadTaskRequest{TaskId: taskId}
	resp, err := s.client.RetrieveUploadTask(ctx, req)
	if err != nil {
		zap.L().Error("Service: Failed to retrieve video upload details",
			zap.Any("user_id", ctx.Value("userId")),
			zap.String("task_id", taskId),
			zap.Error(err))

		var apiErr *twelvelabs.APIError
		if errors.As(err, &apiErr) {
			return nil, NewServiceError(ErrCodeExternalAPIError, apiErr.Message, err)
		}

		return nil, NewServiceError(ErrCodeExternalAPIError, "failed to retrieve video upload details", err)
	}

	return s.parseTaskDetails(resp), nil
}

func (s *service) ListVideos(ctx context.Context, req *ListVideosQuery) ([]*VideoDetails, error) {
	zap.L().Debug("Service: Attempting to list videos...", zap.Any("user_id", ctx.Value("userId")))

	if req == nil || req.CommonListQuery == nil {
		return nil, NewServiceError(ErrCodeInvalidArg, "request cannot be nil", nil)
	}

	queryOpts, err := s.parseCommonListQuery(req.CommonListQuery)
	if err != nil {
		return nil, err
	}

	// TODO: validate index ID
	if req.IndexId != "" {
		return nil, NewServiceError(ErrCodeInvalidArg, "index ID cannot be empty", nil)
	}

	if req.Filename != "" {
		return nil, NewServiceError(ErrCodeInvalidArg, "filename cannot be empty", nil)
	}

	if req.Duration > 0 {
		return nil, NewServiceError(ErrCodeInvalidArg, "duration must be greater than 0", nil)
	}

	if req.Fps > 0 {
		return nil, NewServiceError(ErrCodeInvalidArg, "fps must be greater than 0", nil)
	}

	if req.Width > 0 {
		return nil, NewServiceError(ErrCodeInvalidArg, "width must be greater than 0", nil)
	}

	if req.Height > 0 {
		return nil, NewServiceError(ErrCodeInvalidArg, "height must be greater than 0", nil)
	}

	if req.Size > 0 {
		return nil, NewServiceError(ErrCodeInvalidArg, "size must be greater than 0", nil)
	}

	query := twelvelabs.NewListVideosQuery(req.IndexId, queryOpts...)

	resp, err := s.client.ListVideos(ctx, query)
	if err != nil {
		zap.L().Error("Service: Failed to list videos",
			zap.Any("user_id", ctx.Value("userId")),
			zap.Error(err))

		var apiErr *twelvelabs.APIError
		if errors.As(err, &apiErr) {
			return nil, NewServiceError(ErrCodeExternalAPIError, apiErr.Message, err)
		}

		return nil, NewServiceError(ErrCodeExternalAPIError, "failed to list videos", err)
	}

	var videos []*VideoDetails

	for _, video := range resp {
		videos = append(videos, s.parseVideoDetails(video))
	}

	return videos, nil
}

func (s *service) parseTaskDetails(t *twelvelabs.TaskDetails) *UploadDetails {
	if t == nil {
		return nil
	}

	createdAt, _ := time.Parse(time.RFC1123, t.CreatedAt)
	updatedAt, _ := time.Parse(time.RFC1123, t.UpdatedAt)

	return &UploadDetails{
		TaskId:    t.TaskId,
		VideoId:   t.VideoId,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Status:    t.Status,
		IndexId:   t.IndexId,
	}
}

func (s *service) parseVideoDetails(v *twelvelabs.VideoDetails) *VideoDetails {
	if v == nil {
		return nil
	}

	createdAt, _ := time.Parse(time.RFC1123, v.CreatedAt)
	updatedAt, _ := time.Parse(time.RFC1123, v.UpdatedAt)

	return &VideoDetails{
		VideoId:   v.VideoId,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		IndexId:   v.IndexId,
		Filename:  v.Filename,
		Duration:  v.Duration,
		Fps:       v.Fps,
		Width:     v.Width,
		Height:    v.Height,
		Size:      v.Size,
	}
}
