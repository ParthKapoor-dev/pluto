package inference

import (
	"context"
	"fmt"
)

type Provider int

const (
	ProviderGemini = iota
)

type Inference interface {
	Call(ctx context.Context, model string, prompt string) (string, error)
}

func NewInferenceClient(ctx context.Context, provider Provider) (Inference, error) {

	if provider == ProviderGemini {
		return newGeminiClient(ctx)
	}

	return nil, fmt.Errorf("invalid provider name")
}
