package twelvelabs

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUploadVideo(t *testing.T) {
	require.NotNil(t, client, "TwelveLabs client should be initialized")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	godotenv.Load("../../.env")

	// Avoid running upload tests unless INTEGRATION_TESTS is set to true to save indexing usage
	if os.Getenv("INTEGRATION_TESTS") != "true" {
		t.Skip("Skipping integration test for video upload")
	}

	indexId := os.Getenv("TEST_INDEX_ID")
	videoUrl := os.Getenv("TEST_VIDEO_URL")
	videoFileName := os.Getenv("TEST_VIDEO_FILE")
	videoFile, err := os.OpenFile(videoFileName, os.O_RDONLY, 0o644)
	require.NoError(t, err, "Failed to open video file")

	fileReq := NewUploadVideoRequest(indexId, WithVideoFile(videoFile))
	urlReq := NewUploadVideoRequest(indexId, WithVideoUrl(videoUrl))

	t.Log("Uploading from file to index...")
	fileResp, err := client.UploadVideo(ctx, fileReq)
	require.NoError(t, err, "UploadVideo should not return an error")
	assert.NotEmpty(t, fileResp.GetId(), "Task ID should not be empty")
	assert.NotEmpty(t, fileResp.GetVideoId(), "Video ID should not be empty")

	t.Log("Uploading from URL to index...")
	urlResp, err := client.UploadVideo(ctx, urlReq)
	require.NoError(t, err, "UploadVideo should not return an error")
	assert.NotEmpty(t, urlResp.GetId(), "Task ID should not be empty")
	assert.NotEmpty(t, urlResp.GetVideoId(), "Video ID should not be empty")
}

func TestListUploadTasks(t *testing.T) {
	require.NotNil(t, client, "TwelveLabs client should be initialized")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	t.Log("Listing video upload tasks with default query...")
	req := NewListUploadTasksQuery()
	resp, err := client.ListUploadTasks(ctx, req)
	require.NoError(t, err, "ListUploadTasks should not return an error")
	// Assuming videos have been uploaded before this test
	assert.NotNil(t, resp.GetData(), "Tasks list should not be nil")
	for _, task := range resp.GetData() {
		t.Logf("Task ID: %s, Video ID: %s, Status: %s", task.GetId(), task.GetVideoId(), task.GetStatus())
		assert.NotEmpty(t, task.GetId(), "Task ID should not be empty")
		assert.NotEmpty(t, task.GetVideoId(), "Video ID should not be empty")
	}
}
