package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func main() {

	// cli
	args := os.Args

	if len(args) != 2 {
		panic("invalid arguments for the cli")
	}

	userPrompt := args[1]

	// env
	err := godotenv.Load(".env.local")
	if err != nil {
		panic(fmt.Errorf("loading the env: %w", err))
	}

	GEMINI_API_KEY := os.Getenv("GEMINI_API_KEY")
	ctx := context.Background()

	if GEMINI_API_KEY == "" {
		panic("No gemini api key")
	}

	// gemini client initiation
	geminiClient, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  GEMINI_API_KEY,
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		panic(fmt.Errorf("initiating gemini client: %w", err))
	}

	// Inference call
	result, err := geminiClient.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		genai.Text(userPrompt),
		nil,
	)

	if err != nil {
		panic(fmt.Errorf("at generating result: %w", err))
	}

	// response builder
	if len(result.Candidates) > 0 && len(result.Candidates[0].Content.Parts) > 0 {
		fmt.Println("GEMINI RESPONSE: ", result.Candidates[0].Content.Parts[0].Text)
	} else {
		fmt.Println("No text returned from the model.")
	}
}
