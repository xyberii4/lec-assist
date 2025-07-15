package twelvelabs

import (
	"context"

	"go.uber.org/zap"
)

// Create video uploading task
func (c *twelvelabsClient) UploadVideo(ctx context.Context, req *UploadVideoRequest) (*TaskDetails, error) {
	apiReq := c.apiClient.UploadVideosAPI.CreateVideoIndexingTask(ctx).
		XApiKey(c.getDefaultHeader("x-api-key")).
		ContentType("multipart/form-data").
		IndexId(req.IndexId).
		EnableVideoStream(false)

	if req.VideoFile != nil {
		apiReq = apiReq.VideoFile(req.VideoFile)
		zap.L().Info("Uploading video file",
			zap.String("index_id", req.IndexId),
			zap.String("video_filename", req.VideoFile.Name()))
	} else if req.VideoUrl != "" {
		apiReq = apiReq.VideoUrl(req.VideoUrl)
		zap.L().Info("Uploading video from URL",
			zap.String("index_id", req.IndexId),
			zap.String("video_url", req.VideoUrl))
	}

	resp, r, err := apiReq.Execute()

	if r != nil {
		defer r.Body.Close()
	}

	zap.L().Debug("Rate Limits:",
		zap.String("X-RateLimit-Limit", r.Header.Get("X-RateLimit-Limit")),
		zap.String("X-RateLimit-Remaining", r.Header.Get("X-RateLimit-Remaining")),
		zap.String("X-RateLimit-Reset", r.Header.Get("X-RateLimit-Reset")))
	if err != nil {
		return nil, c.handleHttpError(r, err, "UploadVideo")
	}

	task := &TaskDetails{
		TaskId:  resp.GetId(),
		VideoId: resp.GetVideoId(),
	}

	zap.L().Info("Video upload initiated successfully",
		zap.String("index_id", req.IndexId),
		zap.String("task_id", resp.GetId()),
		zap.String("video_id", resp.GetVideoId()))

	return task, nil
}

// Get list of videos and their upload status
func (c *twelvelabsClient) ListUploadTasks(ctx context.Context, query *ListUploadTasksQuery) ([]*TaskDetails, error) {
	req := c.apiClient.UploadVideosAPI.ListVideoIndexingTasks(ctx).
		XApiKey(c.getDefaultHeader("x-api-key")).
		ContentType(c.getDefaultHeader("Content-Type")).
		Page(query.Page).PageLimit(query.PageLimit).SortBy(query.SortBy).SortOption(query.SortOption)

	if query.CreatedAt != "" {
		req = req.CreatedAt(query.CreatedAt)
	}
	if query.UpdatedAt != "" {
		req = req.UpdatedAt(query.UpdatedAt)
	}
	if query.IndexId != "" {
		req = req.IndexId(query.IndexId)
	}
	if len(query.Status) != 0 {
		req = req.Status(query.Status)
	}
	if query.Filename != "" {
		req = req.Filename(query.Filename)
	}
	if query.Duration > 0 {
		req = req.Duration(query.Duration)
	}
	if query.Width > 0 {
		req = req.Width(query.Width)
	}
	if query.Height > 0 {
		req = req.Height(query.Height)
	}

	resp, r, err := req.Execute()

	if r != nil {
		defer r.Body.Close()
	}

	if err != nil {
		return nil, c.handleHttpError(r, err, "ListUploadTasks")
	}

	uploads := make([]*TaskDetails, 0, len(resp.GetData()))

	for _, task := range resp.GetData() {
		uploads = append(uploads, &TaskDetails{
			TaskId:    task.GetId(),
			VideoId:   task.GetVideoId(),
			CreatedAt: task.GetCreatedAt(),
			UpdatedAt: task.GetUpdatedAt(),
			Status:    task.GetStatus(),
			IndexId:   task.GetIndexId(),
		})
	}

	zap.L().Info("Upload tasks listed successfully",
		zap.Int("count", len(uploads)))

	return uploads, nil
}

// Get details of a specific upload task
func (c *twelvelabsClient) RetrieveUploadTask(ctx context.Context, req *RetrieveUploadTaskRequest) (*TaskDetails, error) {
	apiReq := c.apiClient.UploadVideosAPI.RetrieveVideoIndexingTask(ctx, req.TaskId).
		XApiKey(c.getDefaultHeader("x-api-key")).
		ContentType(c.getDefaultHeader("Content-Type"))

	resp, r, err := apiReq.Execute()

	if r != nil {
		defer r.Body.Close()
	}

	if err != nil {
		return nil, c.handleHttpError(r, err, "RetrieveUploadTask")
	}

	task := &TaskDetails{
		TaskId:    resp.GetId(),
		VideoId:   resp.GetVideoId(),
		CreatedAt: resp.GetCreatedAt(),
		UpdatedAt: resp.GetUpdatedAt(),
		Status:    resp.GetStatus(),
		IndexId:   resp.GetIndexId(),
	}

	zap.L().Info("Upload task retrieved successfully",
		zap.String("task_id", req.TaskId),
		zap.String("status", task.Status))

	return task, nil
}

// List videos in an index
func (c *twelvelabsClient) ListVideos(ctx context.Context, query *ListVideosQuery) ([]*VideoDetails, error) {
	req := c.apiClient.ManageVideosAPI.ListVideos(ctx, query.IndexId).
		XApiKey(c.getDefaultHeader("x-api-key")).
		ContentType(c.getDefaultHeader("Content-Type")).
		Page(query.Page).PageLimit(query.PageLimit).SortBy(query.SortBy).SortOption(query.SortOption)

	if query.CreatedAt != "" {
		req = req.CreatedAt(query.CreatedAt)
	}
	if query.UpdatedAt != "" {
		req = req.UpdatedAt(query.UpdatedAt)
	}
	if query.Filename != "" {
		req = req.Filename(query.Filename)
	}
	if query.Duration > 0 {
		req = req.Duration(query.Duration)
	}
	if query.Width > 0 {
		req = req.Width(query.Width)
	}
	if query.Height > 0 {
		req = req.Height(query.Height)
	}

	resp, r, err := req.Execute()

	if r != nil {
		defer r.Body.Close()
	}

	if err != nil {
		return nil, c.handleHttpError(r, err, "ListVideos")
	}

	// Extract video details
	videos := make([]*VideoDetails, 0, len(resp.GetData()))

	for _, video := range resp.GetData() {
		metadata := video.GetSystemMetadata()
		videos = append(videos, &VideoDetails{
			VideoId:   video.GetId(),
			CreatedAt: video.GetCreatedAt(),
			UpdatedAt: video.GetUpdatedAt(),
			IndexId:   video.GetIndexedAt(),
			Filename:  metadata.GetFilename(),
			Duration:  metadata.GetDuration(),
			Fps:       metadata.GetFps(),
			Width:     metadata.GetWidth(),
			Height:    metadata.GetHeight(),
			Size:      metadata.GetSize(),
		})
	}

	zap.L().Info("Videos listed successfully",
		zap.String("index_id", query.IndexId),
		zap.Int("count", len(videos)))

	return videos, nil
}
