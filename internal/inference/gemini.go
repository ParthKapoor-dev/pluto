package inference

import (
	"context"
	"fmt"

	"github.com/parthkapoor-dev/pluto/pkg"
	"google.golang.org/genai"
)

type geminiClient struct {
	client  *genai.Client
	history []*genai.Content
}

func newGeminiClient(ctx context.Context) (*geminiClient, error) {

	GEMINI_API_KEY, err := pkg.GetEnv("GEMINI_API_KEY")
	if err != nil {
		return nil, err
	}

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  GEMINI_API_KEY,
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		return nil, fmt.Errorf("initiating gemini client: %w", err)
	}

	history := make([]*genai.Content, 0)

	return &geminiClient{client, history}, nil

}

func (gc *geminiClient) Call(ctx context.Context, model string, userPrompt string) (string, error) {

	gc.history = append(gc.history, genai.NewContentFromText(userPrompt, genai.RoleUser))

	// Inference call
	result, err := gc.client.Models.GenerateContent(
		ctx,
		model,
		gc.history,
		nil,
	)

	if err != nil {
		return "", fmt.Errorf("at generating result: %w", err)
	}

	reply := result.Text()

	gc.history = append(gc.history, genai.NewContentFromText(reply, genai.RoleModel))

	return reply, nil

}
