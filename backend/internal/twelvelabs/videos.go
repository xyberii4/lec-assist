package twelvelabs

import (
	"context"
	"fmt"

	sdk "github.com/xyberii4/lec-assist/backend/pkg/twelvelabs"
	"go.uber.org/zap"
)

// Create video uploading task
func (c *twelvelabsClient) UploadVideo(ctx context.Context, req *uploadVideoRequest) (*sdk.InlineObject8, error) {
	if req.IndexId == "" {
		return nil, fmt.Errorf("IndexId is required")
	}
	if (req.VideoFile == nil && req.VideoUrl == "") || (req.VideoFile != nil && req.VideoUrl != "") {
		return nil, fmt.Errorf("either VideoFile or VideoUrl must be provided, but not both")
	}

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
	zap.L().Debug("Rate Limits:",
		zap.String("X-RateLimit-Limit", r.Header.Get("X-RateLimit-Limit")),
		zap.String("X-RateLimit-Remaining", r.Header.Get("X-RateLimit-Remaining")),
		zap.String("X-RateLimit-Reset", r.Header.Get("X-RateLimit-Reset")))
	if err != nil {
		return nil, c.handleHttpError(r, err, "UploadVideo")
	}
	defer r.Body.Close()

	zap.L().Info("Video upload initiated successfully",
		zap.String("index_id", req.IndexId),
		zap.String("task_id", resp.GetId()),
		zap.String("video_id", resp.GetVideoId()))

	return resp, nil
}

// Get list of videos and their upload status
func (c *twelvelabsClient) ListUploadTasks(ctx context.Context, query *listUploadTasksQuery) (*sdk.InlineObject5, error) {
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
	if err != nil {
		return nil, c.handleHttpError(r, err, "ListUploadTasks")
	}

	pageInfo := resp.GetPageInfo()
	zap.L().Info("Upload tasks listed successfully",
		zap.Int32("count", pageInfo.GetTotalResults()))

	return resp, nil
}
