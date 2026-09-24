package cmd

import (
	"context"
	"fmt"

	"github.com/parthkapoor-dev/pluto/internal/inference"
	"github.com/parthkapoor-dev/pluto/pkg"
)

type Cli struct {
}

func NewCli() *Cli {
	return &Cli{}
}

func (c *Cli) Run(args []string) error {

	if len(args) != 2 {
		panic("invalid arguments for the cli")
	}

	ctx := context.Background()

	userPrompt := args[1]

	env, err := pkg.NewEnv()
	if err != nil {

	}

	GEMINI_API_KEY, err := env.Get("GEMINI_API_KEY")
	if err != nil {
		return err
	}

	infClient, err := inference.NewInferenceClient(ctx, inference.ProviderGemini, GEMINI_API_KEY)
	if err != nil {
		return err
	}

	response, err := infClient.Call(ctx, "gemini-2.5-flash", userPrompt)
	if err != nil {
		return err
	}

	fmt.Println("AGENT RESPONSE: ", response)

	return nil
}
