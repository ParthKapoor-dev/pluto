package inference

import (
	"context"
	"fmt"

	"google.golang.org/genai"
)

type geminiClient struct {
	client *genai.Client
}

func newGeminiClient(ctx context.Context, apiKey string) (*geminiClient, error) {

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		return nil, fmt.Errorf("initiating gemini client: %w", err)
	}

	return &geminiClient{client}, nil

}

func (gc *geminiClient) Call(ctx context.Context, model string, userPrompt string) (string, error) {

	// Inference call
	result, err := gc.client.Models.GenerateContent(
		ctx,
		// "gemini-2.5-flash",
		model,
		genai.Text(userPrompt),
		nil,
	)

	if err != nil {
		return "", fmt.Errorf("at generating result: %w", err)
	}

	// response builder
	if len(result.Candidates) > 0 && len(result.Candidates[0].Content.Parts) > 0 {
		return result.Candidates[0].Content.Parts[0].Text, nil
	}

	return "", fmt.Errorf("No text returned from the model.")

}
