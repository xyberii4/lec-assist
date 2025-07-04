package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TwelveLabsAPIKey string
	Production       bool
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}

	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	cfg.Production = os.Getenv("PRODUCTION") == "true"

	cfg.TwelveLabsAPIKey = os.Getenv("TWELVELABS_API_KEY")
	if cfg.TwelveLabsAPIKey == "" {
		return nil, fmt.Errorf("TWELVELABS_API_KEY variable is required")
	}

	return cfg, nil
}
