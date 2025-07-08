package twelvelabs

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestManageIndex(t *testing.T) {
	require.NotNil(t, client, "TwelveLabs client should be initialized")

	indexName := "test-index"
	models := []string{"pegasus1.2"}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var indexId string

	defer func() {
		if indexId != "" {
			t.Logf("Cleaning up index %s", indexId)
			deleteCtx, deleteCancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer deleteCancel()

			deleteErr := client.DeleteIndex(deleteCtx, indexId)
			if deleteErr != nil {
				t.Errorf("Failed to delete index %s: %v", indexId, deleteErr)
			} else {
				t.Logf("Successfully deleted index %s", indexId)
			}
		} else {
			t.Log("indexId is empty, no cleanup needed")
		}
	}()

	t.Logf("Creating index %s:", indexName)
	createResp, err := client.CreateIndex(ctx, indexName, models)
	require.NoError(t, err, "Failed to create index")

	indexId = createResp.GetId()
	require.NotEmpty(t, indexId, "Index ID should not be empty")
	t.Logf("Successfully created index %s with ID %s", indexName, indexId)

	t.Logf("Retrieving index %s:", indexId)
	index, err := client.RetrieveIndex(ctx, indexId)
	require.NoError(t, err, "Failed to retrieve index")
	t.Logf("Successfully retrieved index %s: %+v", indexId, index)
}
