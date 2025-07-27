package tlservice_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	tlservice "github.com/xyberii4/lec-assist/backend/internal/services/twelvelabs"
)

var (
	mockTime        = time.Date(2024, time.January, 9, 11, 11, 20, 463000000, time.UTC).Truncate(time.Second)
	mockUpdatedTime = time.Date(2024, time.January, 9, 11, 17, 15, 296000000, time.UTC)
	mockExpiresTime = mockTime.Add(time.Hour * 24 * 90).Truncate(time.Second) // Roughly 3 months later
)

func TestCreateIndex(t *testing.T) {
	testCases := []struct {
		name          string
		indexName     string
		models        []string
		mockHandler   http.HandlerFunc
		expectedError *tlservice.ServiceError
	}{
		{
			name:      "Valid index creation",
			indexName: "test_service_create_index",
			models:    []string{"pegasus1.2"},
			mockHandler: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(map[string]string{"_id": "12345"})
			},
			expectedError: nil,
		},
		{
			name:        "Invalid models",
			indexName:   "test_service_create_index",
			models:      []string{"invalid_model"},
			mockHandler: nil,
			expectedError: &tlservice.ServiceError{
				Code:    tlservice.ErrCodeInvalidArg,
				Message: "invalid model: invalid_model",
				Err:     nil,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf("Running test case: %s", tc.name)
			setMockHandler(tc.mockHandler)

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			_, err := testService.CreateIndex(ctx, tc.indexName, tc.models)
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

func TestDeleteIndex(t *testing.T) {
	testCases := []struct {
		name          string
		indexId       string
		mockHandler   http.HandlerFunc
		expectedError *tlservice.ServiceError
	}{
		{
			name:    "Valid index deletion",
			indexId: "12345",
			mockHandler: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNoContent)
			},
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf("Running test case: %s", tc.name)
			setMockHandler(tc.mockHandler)

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			err := testService.DeleteIndex(ctx, tc.indexId)
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

func TestListIndexes(t *testing.T) {
	mockResponse := map[string]interface{}{
		"data": []map[string]interface{}{
			{
				"_id":            "62d85fe7340fa665f1fda5dd",
				"created_at":     mockTime.Format(time.RFC3339Nano), // Use RFC3339Nano for milliseconds
				"updated_at":     mockTime.Format(time.RFC3339Nano),
				"expires_at":     mockExpiresTime.Format(time.RFC3339Nano),
				"index_name":     "index-01",
				"total_duration": 1363.76,
				"video_count":    2,
				"models": []map[string]interface{}{
					{
						"model_name":    "marengo2.7",
						"model_options": []string{"visual", "audio"},
						"addons":        []string{"thumbnail"},
					},
					{
						"model_name":    "pegasus1.2",
						"model_options": []string{"visual", "audio"},
					},
				},
			},
			{
				"_id":            "62d858ee340fa665f1fda5d8",
				"created_at":     mockTime.Format(time.RFC3339Nano),
				"updated_at":     mockTime.Format(time.RFC3339Nano),
				"expires_at":     mockExpiresTime.Format(time.RFC3339Nano),
				"index_name":     "index-02",
				"total_duration": 579.12,
				"video_count":    1,
				"models": []map[string]interface{}{
					{
						"model_name":    "marengo2.7",
						"model_options": []string{"visual"},
						"addons":        []string{"thumbnail"},
					},
					{
						"model_name":    "pegasus1.2",
						"model_options": []string{"visual", "audio"},
					},
				},
			},
		},
		"page_info": map[string]interface{}{
			"limit_per_page": 10,
			"page":           1,
			"total_page":     1,
			"total_results":  2,
		},
	}
	invalidPage := int32(-1)
	testCases := []struct {
		name          string
		req           *tlservice.ListIndexesQuery
		mockHandler   http.HandlerFunc
		expectedError *tlservice.ServiceError
	}{
		{
			name: "Default query",
			req:  &tlservice.ListIndexesQuery{CommonListQuery: &tlservice.CommonListQuery{}},
			mockHandler: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(mockResponse)
			},
			expectedError: nil,
		},
		{
			name: "Invalid page number",
			req: &tlservice.ListIndexesQuery{
				CommonListQuery: &tlservice.CommonListQuery{
					Page: &invalidPage,
				},
			},
			mockHandler:   nil,
			expectedError: tlservice.NewServiceError(tlservice.ErrCodeInvalidArg, "page must be greater than 0", nil),
		},
		{
			name: "Invalid model options",
			req: &tlservice.ListIndexesQuery{
				CommonListQuery: &tlservice.CommonListQuery{},
				ModelOptions:    "visual audio",
			},
			mockHandler:   nil,
			expectedError: tlservice.NewServiceError(tlservice.ErrCodeInvalidArg, "invalid model option: visual audio", nil),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf("Running test case: %s", tc.name)
			setMockHandler(tc.mockHandler)

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			indexes, err := testService.ListIndexes(ctx, tc.req)
			if tc.expectedError != nil {
				require.Error(t, err, "Expected error but got none")
				var svcErr *tlservice.ServiceError
				require.True(t, errors.As(err, &svcErr), "Expected error type does not match")
				assert.Equal(t, tc.expectedError.Code, svcErr.Code, "Expected error code does not match")
				assert.Equal(t, tc.expectedError.Message, svcErr.Message, "Expected error code does not match")
				assert.Equal(t, tc.expectedError.Err, svcErr.Err, "Expected underlying error does not match")
			} else {
				require.NoError(t, err, "Expected no error but got one")
				assert.NotEmpty(t, indexes, "Expected non-empty index list")
			}
		})
	}
}

func TestRetrieveIndex(t *testing.T) {
	mockRetrieveIndexBody := map[string]interface{}{
		"_id":            "62d9bafa90077fc60af827a0",
		"created_at":     mockTime.Format(time.RFC3339Nano), // Use RFC3339Nano for milliseconds
		"updated_at":     mockUpdatedTime.Format(time.RFC3339Nano),
		"expires_at":     mockExpiresTime.Format(time.RFC3339Nano),
		"index_name":     "myIndex",
		"total_duration": 8716.8969,
		"video_count":    12,
		"models": []map[string]interface{}{
			{
				"model_name":    "marengo2.7",
				"model_options": []string{"visual", "audio"},
				"addons":        []string{"thumbnail"},
			},
			{
				"model_name":    "pegasus1.2",
				"model_options": []string{"visual", "audio"},
				// "addons" field is optional, so it can be omitted if not present in the API response
			},
		},
	}
	testCases := []struct {
		name          string
		indexId       string
		mockHandler   http.HandlerFunc
		expectedError *tlservice.ServiceError
	}{
		{
			name:    "Valid index retrieval",
			indexId: "12345",
			mockHandler: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(mockRetrieveIndexBody)
			},
			expectedError: nil,
		},
		{
			name:          "Invalid index ID",
			indexId:       "",
			mockHandler:   nil,
			expectedError: tlservice.NewServiceError(tlservice.ErrCodeInvalidArg, "index id cannot be empty", nil),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			setMockHandler(tc.mockHandler)

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			indexDetails, err := testService.RetrieveIndex(ctx, tc.indexId)
			if tc.expectedError != nil {
				if tc.expectedError != nil {
					require.Error(t, err, "Expected error but got none")
					var svcErr *tlservice.ServiceError
					require.True(t, errors.As(err, &svcErr), "Expected error type does not match")
					assert.Equal(t, tc.expectedError.Code, svcErr.Code, "Expected error code does not match")
					assert.Equal(t, tc.expectedError.Message, svcErr.Message, "Expected error code does not match")
					assert.Equal(t, tc.expectedError.Err, svcErr.Err, "Expected underlying error does not match")
				} else {
					require.NoError(t, err, "Expected no error but got one")
					assert.NotEmpty(t, indexDetails, "Expected non-empty index list")
				}
			}
		})
	}
}
