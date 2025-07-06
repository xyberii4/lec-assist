package utils

import (
	"testing"

	"github.com/xyberii4/lec-assist/backend/internal/config"
	"github.com/xyberii4/lec-assist/backend/internal/log"
	"go.uber.org/zap"
)

type TestEnv struct {
	Config *config.Config
}

func InitTest(m *testing.M) (*TestEnv, error) {
	// Initialize any necessary test configurations or mock services here.
	// This function can be used to set up the environment before running tests.

	log.InitLogger(false)

	cfg, err := config.LoadConfig("../../.env")
	if err != nil {
		return nil, err
	}
	zap.L().Info("Test configuration loaded successfully")

	return &TestEnv{
		Config: cfg,
	}, nil
}
