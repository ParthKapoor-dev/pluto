package pkg

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const DEFAULT_ENV = ".env.local"

type Env struct {
}

func NewEnv() (*Env, error) {

	err := godotenv.Load(DEFAULT_ENV)
	if err != nil {
		return nil, fmt.Errorf("loading the env: %w", err)
	}

	return &Env{}, nil
}

func (e *Env) Get(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("no variable variable with key %s exists", key)
	}
	return value, nil
}
