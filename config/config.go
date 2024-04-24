package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to parse .env: %w", err)
	}

	return &Config{
		Port: os.Getenv("PORT"),
	}, nil
}
