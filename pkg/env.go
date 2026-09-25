package pkg

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const DEFAULT_ENV = ".env.local"

func LoadEnv() error {

	err := godotenv.Load(DEFAULT_ENV)

	if err != nil {
		return fmt.Errorf("loading the env: %w", err)
	}

	return nil
}

func GetEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("no variable variable with key %s exists", key)
	}
	return value, nil
}
