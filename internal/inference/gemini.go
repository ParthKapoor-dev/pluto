package inference

import (
	"context"
	"fmt"
	"os"
	"strings"

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
		&genai.GenerateContentConfig{
			SystemInstruction: &genai.Content{Parts: []*genai.Part{{Text: "If you want to list files in current directory, respond with <TOOL-CALL>LIST_FILES<TOOL-CALL>"}}},
		},
	)

	if err != nil {
		return "", fmt.Errorf("at generating result: %w", err)
	}

	reply := result.Text()

	gc.history = append(gc.history, genai.NewContentFromText(reply, genai.RoleModel))

	return reply, nil

}

func (gc *geminiClient) Stream(ctx context.Context, model string, userPrompt string) error {

	prompt := userPrompt

	for {

		gc.history = append(gc.history, genai.NewContentFromText(prompt, genai.RoleUser))

		var reply []string

		fmt.Print("AGENT: ")

		// Inference call
		for result, err := range gc.client.Models.GenerateContentStream(
			ctx,
			model,
			gc.history,
			&genai.GenerateContentConfig{
				SystemInstruction: &genai.Content{Parts: []*genai.Part{
					{Text: "If you want to list files in current directory, respond with <TOOL-CALL>LIST_FILES</TOOL-CALL>"}}},
			},
		) {
			if err != nil {
				return fmt.Errorf("at generating result: %w", err)
			}

			fmt.Print(result.Text(), " ")
			reply = append(reply, result.Text())
		}

		gc.history = append(gc.history, genai.NewContentFromText(strings.Join(reply, " "), genai.RoleModel))

		fmt.Println()

		var err error
		prompt, err = gc.toolCall()
		if err != nil {
			return err
		}

		if prompt == "" {
			return nil
		}

	}

}

func (gc *geminiClient) toolCall() (string, error) {

	response := gc.history[len(gc.history)-1].Parts[0].Text

	if strings.TrimSpace(response) != "<TOOL-CALL>LIST_FILES</TOOL-CALL>" {
		return "", nil
	}

	files, err := os.ReadDir(".")
	if err != nil {
		return "", err
	}

	var fileNames []string

	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	return strings.Join(fileNames, " "), nil
}
