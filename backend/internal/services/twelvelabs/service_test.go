package tlservice_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	tlservice "github.com/xyberii4/lec-assist/backend/internal/services/twelvelabs"
	"github.com/xyberii4/lec-assist/backend/internal/twelvelabs"
	"github.com/xyberii4/lec-assist/backend/internal/utils"
	"go.uber.org/zap"
)

var (
	testService          tlservice.Service
	mockTwelveLabsServer *httptest.Server
	mockHandler          http.HandlerFunc
	handlerMu            sync.Mutex
)

func TestMain(m *testing.M) {
	// Initialize the logger
	env, err := utils.InitTest(m, "../../../.env")
	if err != nil {
		zap.L().Error("Failed to initialize test environment", zap.Error(err))
		return
	}

	mockTwelveLabsServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlerMu.Lock()
		handler := mockHandler
		handlerMu.Unlock()

		if handler != nil {
			handler(w, r)
		} else {
			zap.L().Error("No mock handler set for TwelveLabs server")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"message": "handler not set"})
		}
	}))

	mockApiClient := twelvelabs.NewTestClient(
		mockTwelveLabsServer.Client(),
		mockTwelveLabsServer.URL,
		env.Config.TwelveLabsAPIKey,
	)

	testService = tlservice.NewService(mockApiClient)

	// Run the tests
	m.Run()

	// Clean up and exit
	zap.L().Info("Tests completed, shutting down mock server")
	mockTwelveLabsServer.Close()
	zap.L().Sync()
}

func setMockHandler(handler http.HandlerFunc) {
	handlerMu.Lock()
	defer handlerMu.Unlock()
	mockHandler = handler
}
