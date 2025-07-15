package twelvelabs

import (
	"context"

	sdk "github.com/xyberii4/lec-assist/backend/pkg/twelvelabs"
	"go.uber.org/zap"
)

func (c *twelvelabsClient) Analyze(ctx context.Context, req *sdk.AnalyzeRequest) (string, error) {
	req.SetStream(false) // Streaming disabled as OpenAPI generation does not support NDJSON streams
	resp, r, err := c.apiClient.AnalyzeVideosAPI.GenerateTextRepresentation(ctx).
		XApiKey(c.getDefaultHeader("x-api-key")).
		ContentType(c.getDefaultHeader("Content-Type")).
		AnalyzeRequest(*req).
		Execute()
	if err != nil {
		return "", c.handleHttpError(r, err, "Analyze")
	}
	defer r.Body.Close()

	respData := resp.NonStreamGenerateResponse.GetData()

	zap.L().Info("Successfully analyzed video",
		zap.String("video_id", req.VideoId))

	return respData, nil
}

// Generates title, topics or hashtags from video
func (c *twelvelabsClient) Gist(ctx context.Context, req *sdk.GistRequest) (*GistResult, error) {
	resp, r, err := c.apiClient.AnalyzeVideosAPI.Gist(ctx).
		XApiKey(c.getDefaultHeader("x-api-key")).
		ContentType(c.getDefaultHeader("Content-Type")).
		GistRequest(*req).
		Execute()
	if err != nil {
		return nil, c.handleHttpError(r, err, "Gist")
	}
	defer r.Body.Close()

	newGist := &GistResult{
		VideoId:  req.GetVideoId(),
		Topics:   []string{},
		Hashtags: []string{},
	}

	if resp.GetTitle() != "" {
		newGist.Title = resp.GetTitle()
	}

	if resp.GetTopics() != nil {
		newGist.Topics = resp.GetTopics()
	}

	if resp.GetHashtags() != nil {
		newGist.Hashtags = resp.GetHashtags()
	}

	zap.L().Info("Successfully generated gist",
		zap.String("video_id", req.VideoId))

	return newGist, nil
}

// Video ID, output type (chapter, highlight, summary) and prompt
func (c *twelvelabsClient) Summarise(ctx context.Context, req *sdk.SummarizeRequest) (*SummariseResult, error) {
	resp, r, err := c.apiClient.AnalyzeVideosAPI.Summarize(ctx).
		XApiKey(c.getDefaultHeader("x-api-key")).
		ContentType(c.getDefaultHeader("Content-Type")).
		SummarizeRequest(*req).
		Execute()
	if err != nil {
		return nil, c.handleHttpError(r, err, "Summarise")
	}
	defer r.Body.Close()

	// Extract chapter, highlight, or summary details based on the response
	newSummarise := &SummariseResult{
		VideoId:    req.GetVideoId(),
		Chapter:    []*ChapterDetails{},
		Highlights: []*HighlightDetails{},
	}

	if resp.Chapter != nil {
		for _, chapter := range resp.Chapter.GetChapters() {
			newChapterDetail := &ChapterDetails{
				ChapterNumber: int32(chapter.GetChapterNumber()),
				Start:         int32(chapter.GetStart()),
				End:           int32(chapter.GetEnd()),
				Title:         chapter.GetChapterTitle(),
				Description:   chapter.GetChapterSummary(),
			}

			newSummarise.Chapter = append(newSummarise.Chapter, newChapterDetail)
		}
	} else if resp.Highlight != nil {
		for _, highlight := range resp.Highlight.GetHighlights() {
			newHighlightDetail := &HighlightDetails{
				Start:       int32(highlight.GetStart()),
				End:         int32(highlight.GetEnd()),
				Title:       highlight.GetHighlight(),
				Description: highlight.GetHighlightSummary(),
			}
			newSummarise.Highlights = append(newSummarise.Highlights, newHighlightDetail)
		}
	} else if resp.Summary != nil {
		newSummarise.Summary = resp.Summary.GetSummary()
	}

	zap.L().Info("Successfully summarised video",
		zap.String("video_id", req.GetVideoId()),
		zap.String("output_type", req.GetType()))

	return newSummarise, nil
}
