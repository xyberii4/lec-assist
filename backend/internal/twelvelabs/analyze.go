package twelvelabs

import (
	"context"

	sdk "github.com/xyberii4/lec-assist/backend/pkg/twelvelabs"
	"go.uber.org/zap"
)

func (c *twelvelabsClient) Analyze(ctx context.Context, req *sdk.AnalyzeRequest) (*sdk.Analyze200Response, error) {
	req.SetStream(false) // Streaming disabled as OpenAPI generation does not support NDJSON streams
	resp, r, err := c.apiClient.AnalyzeVideosAPI.GenerateTextRepresentation(ctx).
		XApiKey(c.getDefaultHeader("x-api-key")).
		ContentType(c.getDefaultHeader("Content-Type")).
		AnalyzeRequest(*req).
		Execute()
	if err != nil {
		return nil, c.handleHttpError(r, err, "Analyze")
	}
	defer r.Body.Close()

	zap.L().Info("Successfully analyzed video",
		zap.String("video_id", req.VideoId))

	return resp, nil
}

// Generates title, topics or hashtags from video
func (c *twelvelabsClient) Gist(ctx context.Context, req *sdk.GistRequest) (*sdk.Gist, error) {
	resp, r, err := c.apiClient.AnalyzeVideosAPI.Gist(ctx).
		XApiKey(c.getDefaultHeader("x-api-key")).
		ContentType(c.getDefaultHeader("Content-Type")).
		GistRequest(*req).
		Execute()
	if err != nil {
		return nil, c.handleHttpError(r, err, "Gist")
	}
	defer r.Body.Close()

	zap.L().Info("Successfully generated gist",
		zap.String("video_id", req.VideoId))

	return resp, nil
}

func (c *twelvelabsClient) Summarise(ctx context.Context, req *sdk.SummarizeRequest) (*sdk.InlineObject19, error) {
	resp, r, err := c.apiClient.AnalyzeVideosAPI.Summarize(ctx).
		XApiKey(c.getDefaultHeader("x-api-key")).
		ContentType(c.getDefaultHeader("Content-Type")).
		SummarizeRequest(*req).
		Execute()
	if err != nil {
		return nil, c.handleHttpError(r, err, "Summarise")
	}
	defer r.Body.Close()

	zap.L().Info("Successfully summarised video",
		zap.String("video_id", req.VideoId))

	return resp, nil
}
