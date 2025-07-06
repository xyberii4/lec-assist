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

func LoadConfig(envPath string) (*Config, error) {
	cfg := &Config{}

	var err error
	if envPath != "" {
		err = godotenv.Load(envPath)
	} else {
		err = godotenv.Load()
	}

	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	cfg.Production = os.Getenv("PRODUCTION") == "true"

	cfg.TwelveLabsAPIKey = os.Getenv("TWELVELABS_API_KEY")
	if cfg.TwelveLabsAPIKey == "" {
		return nil, fmt.Errorf("TWELVELABS_API_KEY variable is required")
	}

	return cfg, nil
}
