package twelvelabs

import (
	"context"
	"testing"

	"github.com/xyberii4/lec-assist/backend/internal/utils"
	"go.uber.org/zap"
)

func TestMain(m *testing.M) {
	env, err := utils.InitTest(m)
	if err != nil {
		zap.L().Error("Failed to load config for tests", zap.Error(err))
	}

	// Initialize the TwelveLabs client with the test configuration
	InitClient(env.Config.TwelveLabsAPIKey)

	// Run the tests
	m.Run()
}

func TestCreateIndex(t *testing.T) {
	client := GetClient()
	if client == nil {
		t.Fatal("TwelveLabs client is not initialized")
	}

	resp, err := client.CreateIndex(context.Background(), "test", []string{"pegasus1.2"})
	if err != nil {
		t.Fatalf("Failed to create index: %v", err)
	}

	if resp == nil {
		t.Fatal("CreateIndex response is nil")
	}

	t.Logf("CreateIndex response: %+v", resp)
}
