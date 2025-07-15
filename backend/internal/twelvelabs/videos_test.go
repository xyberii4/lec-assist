package twelvelabs

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUploadVideo(t *testing.T) {
	require.NotNil(t, client, "TwelveLabs client should be initialized")

	ctx := context.Background()

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
	fileTaskId := fileResp.TaskId
	fileVideoId := fileResp.VideoId
	require.NoError(t, err, "UploadVideo should not return an error")
	assert.NotEmpty(t, fileTaskId, "Task ID should not be empty")
	assert.NotEmpty(t, fileVideoId, "Video ID should not be empty")

	t.Log("Retrieving video upload task for file upload...")

	fileRetrieveTaskReq := &RetrieveUploadTaskRequest{
		TaskId: fileTaskId,
	}

	fileTaskResp, err := client.RetrieveUploadTask(ctx, fileRetrieveTaskReq)
	require.NoError(t, err, "GetUploadTask should not return an error")
	assert.Equal(t, fileTaskId, fileTaskResp.TaskId, "Task ID should match")
	assert.Equal(t, fileVideoId, fileTaskResp.VideoId, "Video ID should match")

	t.Log("Uploading from URL to index...")
	urlResp, err := client.UploadVideo(ctx, urlReq)
	urlTaskId := urlResp.TaskId
	urlVideoId := urlResp.VideoId
	require.NoError(t, err, "UploadVideo should not return an error")
	assert.NotEmpty(t, urlResp.TaskId, "Task ID should not be empty")
	assert.NotEmpty(t, urlResp.VideoId, "Video ID should not be empty")

	t.Log("Retrieving video upload task for URL upload...")

	urlRetrieveTaskReq := &RetrieveUploadTaskRequest{
		TaskId: urlTaskId,
	}

	urlTaskResp, err := client.RetrieveUploadTask(ctx, urlRetrieveTaskReq)
	require.NoError(t, err, "GetUploadTask should not return an error")
	assert.Equal(t, urlTaskId, urlTaskResp.TaskId, "Task ID should match")
	assert.Equal(t, urlVideoId, urlTaskResp.VideoId, "Video ID should match")
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
	for _, task := range resp {
		fmt.Println(task)
		t.Logf("Task ID: %s, Video ID: %s, Status: %s", task.TaskId, task.VideoId, task.Status)
		assert.NotEmpty(t, task.TaskId, "Task ID should not be empty")
		assert.NotEmpty(t, task.VideoId, "Video ID should not be empty")
	}
}

func TestListVideos(t *testing.T) {
	require.NotNil(t, client, "TwelveLabs client should be initialized")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	godotenv.Load("../../.env")
	indexId := os.Getenv("TEST_INDEX_ID")

	t.Log("Listing videos with default query...")
	req := NewListVideosQuery(indexId)
	resp, err := client.ListVideos(ctx, req)
	require.NoError(t, err, "ListVideos should not return an error")
	assert.NotNil(t, resp, "Videos list should not be nil")
	for _, video := range resp {
		id := video.VideoId
		filename := video.Filename
		t.Logf("Video ID: %s, Filename: %s", id, filename)
		assert.NotEmpty(t, video.VideoId, "Video ID should not be empty")
		assert.NotEmpty(t, filename, "Filename should not be empty")
	}
}
