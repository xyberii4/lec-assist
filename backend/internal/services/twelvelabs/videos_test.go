package tlservice_test

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	tlservice "github.com/xyberii4/lec-assist/backend/internal/services/twelvelabs"
	"github.com/xyberii4/lec-assist/backend/internal/twelvelabs"
)

// retrive and list tests not included as they are similar to the ones for indexes
func createTempFile(t *testing.T, filename string, content string) *os.File {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, filename)
	file, err := os.Create(filePath)
	require.NoError(t, err)
	_, err = file.WriteString(content)
	require.NoError(t, err)
	_, err = file.Seek(0, io.SeekStart) // Rewind to the beginning
	return file
}

func TestUploadVideo(t *testing.T) {
	testCases := []struct {
		name          string
		request       *twelvelabs.UploadVideoRequest
		mockHandler   http.HandlerFunc
		expectedError *tlservice.ServiceError
	}{
		{
			name:    "Valid video upload with file",
			request: &twelvelabs.UploadVideoRequest{IndexId: "test_index", VideoFile: createTempFile(t, "test_video.mp4", "dummy content")},
			mockHandler: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(map[string]string{
					"_id":      "62a1ec6d9ea24f052b971a0f",
					"video_id": "62a1ec6d9ea24f052b971a0f",
				})
			},
			expectedError: nil,
		},
		{
			name:    "Valid video upload with URL",
			request: &twelvelabs.UploadVideoRequest{IndexId: "test_index", VideoUrl: "http://example.com/video.mp4"},
			mockHandler: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(map[string]string{
					"_id":      "62a1ec6d9ea24f052b971a0f",
					"video_id": "62a1ec6d9ea24f052b971a0f",
				})
			},
			expectedError: nil,
		},
		{
			name:        "Invalid request with no video file or URL",
			request:     &twelvelabs.UploadVideoRequest{IndexId: "test_index"},
			mockHandler: nil,
			expectedError: &tlservice.ServiceError{
				Code:    tlservice.ErrCodeInvalidArg,
				Message: "either video file or video URL must be provided",
				Err:     nil,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			setMockHandler(tc.mockHandler)
			_, err := testService.UploadVideo(ctx, tc.request)
			if tc.expectedError != nil {
				require.Error(t, err, "Expected error but got none")
				var svcErr *tlservice.ServiceError
				require.True(t, errors.As(err, &svcErr), "Expected error type does not match")
				assert.Equal(t, tc.expectedError.Code, svcErr.Code, "Expected error code does not match")
				assert.Equal(t, tc.expectedError.Message, svcErr.Message, "Expected error code does not match")
				assert.Equal(t, tc.expectedError.Err, svcErr.Err, "Expected underlying error does not match")
			} else {
				require.NoError(t, err, "Expected no error but got one")
			}
		})
	}
}
