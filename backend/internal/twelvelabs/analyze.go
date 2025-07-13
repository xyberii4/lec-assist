package twelvelabs

import (
	"context"

	sdk "github.com/xyberii4/lec-assist/backend/pkg/twelvelabs"
	"go.uber.org/zap"
)

func (c *twelvelabsClient) Analyze(ctx context.Context, videoId string, prompt string, temp float32, stream bool) (*sdk.Analyze200Response, error) {
	analyzeReq := *sdk.NewAnalyzeRequest(videoId, prompt)

	if temp != 0.2 {
		analyzeReq.SetTemperature(temp)
	}
	if !(stream) {
		analyzeReq.SetStream(stream)
	}

	resp, r, err := c.apiClient.AnalyzeVideosAPI.Analyze(ctx).
		XApiKey(c.getDefaultHeader("x-api-key")).
		ContentType(c.getDefaultHeader("Content-Type")).
		AnalyzeRequest(analyzeReq).
		Execute()
	if err != nil {
		return nil, c.handleHttpError(r, err, "Analyze")
	}

	zap.L().Info("Successfully analyzed video",
		zap.String("video_id", videoId))

	return resp, nil
}

// Generates title, topics or hashtags from video
func (c *twelvelabsClient) Gist(ctx context.Context, videoId string, outputTypes []string) (*sdk.Gist, error) {
	gistReq := *sdk.NewGistRequest(videoId, outputTypes)
	resp, r, err := c.apiClient.AnalyzeVideosAPI.Gist(ctx).
		XApiKey(c.getDefaultHeader("x-api-key")).
		ContentType(c.getDefaultHeader("Content-Type")).
		GistRequest(gistReq).
		Execute()
	if err != nil {
		return nil, c.handleHttpError(r, err, "Gist")
	}

	zap.L().Info("Successfully generated gist",
		zap.String("video_id", videoId))

	return resp, nil
}

func (c *twelvelabsClient) Summarise(ctx context.Context, videoId string, outputType string, prompt string, temp float32) (*sdk.InlineObject19, error) {
	summariseReq := *sdk.NewSummarizeRequest(videoId, outputType)

	if prompt != "" {
		summariseReq.SetPrompt(prompt)
	}
	if temp != 0.2 {
		summariseReq.SetTemperature(temp)
	}

	resp, r, err := c.apiClient.AnalyzeVideosAPI.Summarize(ctx).
		XApiKey(c.getDefaultHeader("x-api-key")).
		ContentType(c.getDefaultHeader("Content-Type")).
		SummarizeRequest(summariseReq).
		Execute()
	if err != nil {
		return nil, c.handleHttpError(r, err, "Summarise")
	}

	zap.L().Info("Successfully summarised video",
		zap.String("video_id", videoId))

	return resp, nil
}
