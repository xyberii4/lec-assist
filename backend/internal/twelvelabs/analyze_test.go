package twelvelabs

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAnalyze(t *testing.T) {
	assert.NotNil(t, client, "TwelveLabs client should be initialized")

	context, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	videoId := os.Getenv("TEST_VIDEO1_ID")
	prompt := "Create defintions for keywords mentioned in the video"
	stream := false

	resp, err := client.Analyze(context, videoId, prompt, 0.2, stream)
	assert.NoError(t, err, "Analyze should not return an error")

	data := resp.NonStreamGenerateResponse.GetData()
	t.Logf("Analyze response: %v", data)
}

func TestGist(t *testing.T) {
	require.NotNil(t, client, "TwelveLabs client should be initialized")

	context, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	videoId := os.Getenv("TEST_VIDEO1_ID")
	outputTypes := []string{"title", "topic"}
	resp, err := client.Gist(context, videoId, outputTypes)
	require.NoError(t, err, "Gist should not return an error")
	title := resp.GetTitle()
	topics := resp.GetTopics()
	t.Logf("Gist response: title=%s, topics=%v", title, topics)
}

func TestSummarise(t *testing.T) {
	require.NotNil(t, client, "TwelveLabs client should be initialized")

	context, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	videoId := os.Getenv("TEST_VIDEO2_ID")
	outputType := "summary"

	resp, err := client.Summarise(context, videoId, outputType, "", 0.2)
	require.NoError(t, err, "Summarise should not return an error")

	summary := *resp.Summary.Summary
	t.Logf("Summary response: %s", summary)
}
