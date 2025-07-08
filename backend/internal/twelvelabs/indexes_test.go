package twelvelabs

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestManageIndex(t *testing.T) {
	require.NotNil(t, client, "TwelveLabs client should be initialized")

	indexName := "test-manage-index"
	models := []string{"pegasus1.2"}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var indexId string

	defer func() {
		cleanup(t, indexId)
	}()

	t.Logf("Creating index %s...", indexName)
	createResp, err := client.CreateIndex(ctx, indexName, models)
	require.NoError(t, err, "Failed to create index")

	indexId = createResp.GetId()
	require.NotEmpty(t, indexId, "Index ID should not be empty")
	t.Logf("Successfully created index %s with ID %s", indexName, indexId)

	t.Logf("Retrieving index %s...", indexId)
	index, err := client.RetrieveIndex(ctx, indexId)
	require.NoError(t, err, "Failed to retrieve index")
	t.Logf("Successfully retrieved index %s: %+v", indexId, index)
}

func TestListIndexes(t *testing.T) {
	require.NotNil(t, client, "TwelveLabs client should be initialized")

	indexName := "test-list-indexes"
	models := []string{"pegasus1.2"}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var indexId string
	defer func() {
		cleanup(t, indexId)
	}()

	t.Logf("Creating index %s...", indexName)
	createResp, err := client.CreateIndex(ctx, indexName, models)
	require.NoError(t, err, "Failed to create index")
	indexId = createResp.GetId()
	t.Logf("Successfully created index %s with ID %s", indexName, indexId)

	t.Log("Listing indexes with default query...")
	query := NewListIndexesQuery()
	defaultListResp, err := client.ListIndexes(ctx, query)
	require.NoError(t, err, "Failed to list indexes")
	assert.GreaterOrEqual(t, len(defaultListResp.GetData()), 1, "Should return at least one index")

	indexFound := false
	for _, index := range defaultListResp.GetData() {
		if index.GetId() == indexId {
			indexFound = true
			break
		}
	}
	assert.True(t, indexFound, "Created index should be in the list")

	t.Log("Listing indexes with custom query...")
	query = NewListIndexesQuery(WithIndexName(indexName), WithModelFamily("pegasus"))
	customListResp, err := client.ListIndexes(ctx, query)
	require.NoError(t, err, "Failed to list indexes with custom query")
	assert.Equal(t, 1, len(customListResp.GetData()), "Should return exactly one index with the specified name and model family")
	assert.Equal(t, indexId, customListResp.GetData()[0].GetId(), "Returned index ID should match the created index ID")

	t.Log("Successfully listed indexes")
}

func cleanup(t *testing.T, indexId string) {
	if indexId != "" {
		t.Logf("Cleaning up index %s...", indexId)
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
}
