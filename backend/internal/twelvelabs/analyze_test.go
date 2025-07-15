package twelvelabs

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	sdk "github.com/xyberii4/lec-assist/backend/pkg/twelvelabs"
)

func TestAnalyze(t *testing.T) {
	godotenv.Load("../../.env")

	// Avoid running upload tests unless INTEGRATION_TESTS is set to true to save tokens
	if os.Getenv("INTEGRATION_TESTS") != "true" {
		t.Skip("Skipping integration test for video upload")
	}

	assert.NotNil(t, client, "TwelveLabs client should be initialized")

	context, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	videoId := os.Getenv("TEST_VIDEO2_ID")
	prompt := "Create defintions for keywords mentioned in the video"

	req := sdk.NewAnalyzeRequest(videoId, prompt)

	resp, err := client.Analyze(context, req)
	assert.NoError(t, err, "Analyze should not return an error")

	t.Logf("Analyze response: %v", resp)
}

func TestGist(t *testing.T) {
	godotenv.Load("../../.env")

	// Avoid running upload tests unless INTEGRATION_TESTS is set to true to save tokens
	if os.Getenv("INTEGRATION_TESTS") != "true" {
		t.Skip("Skipping integration test for video upload")
	}

	require.NotNil(t, client, "TwelveLabs client should be initialized")

	context, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	videoId := os.Getenv("TEST_VIDEO1_ID")
	outputTypes := []string{"title", "topic"}
	gistReq := sdk.NewGistRequest(videoId, outputTypes)

	resp, err := client.Gist(context, gistReq)
	require.NoError(t, err, "Gist should not return an error")
	title := resp.Title
	topics := resp.Topics
	t.Logf("Gist response: title=%s, topics=%v", title, topics)
}

func TestSummarise(t *testing.T) {
	godotenv.Load("../../.env")

	// Avoid running upload tests unless INTEGRATION_TESTS is set to true to save tokens
	if os.Getenv("INTEGRATION_TESTS") != "true" {
		t.Skip("Skipping integration test for video upload")
	}

	require.NotNil(t, client, "TwelveLabs client should be initialized")

	context, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	videoId := os.Getenv("TEST_VIDEO2_ID")
	outputType := "summary"
	summaryReq := sdk.NewSummarizeRequest(videoId, outputType)

	resp, err := client.Summarise(context, summaryReq)
	require.NoError(t, err, "Summarise should not return an error")

	summary := resp.Summary
	highlights := resp.Highlights
	assert.NotEmpty(t, summary, "Summary should not be empty")
	assert.Empty(t, highlights, "Highlights should be empty for summary request")
}
