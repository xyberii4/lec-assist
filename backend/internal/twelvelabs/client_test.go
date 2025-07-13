package twelvelabs

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/xyberii4/lec-assist/backend/internal/utils"
	"go.uber.org/zap"
)

var client Client

func TestMain(m *testing.M) {
	env, err := utils.InitTest(m)
	if err != nil {
		zap.L().Error("Failed to load config for tests", zap.Error(err))
	}

	godotenv.Load("../../.env")

	// Initialize the TwelveLabs client with the test configuration
	InitClient(env.Config.TwelveLabsAPIKey)
	client = GetClient()

	// Run the tests
	m.Run()
}
