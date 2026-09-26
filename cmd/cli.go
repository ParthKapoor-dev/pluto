package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/parthkapoor-dev/pluto/internal/inference"
	"github.com/parthkapoor-dev/pluto/pkg"
)

type Cli struct {
}

func NewCli() *Cli {
	return &Cli{}
}

func (c *Cli) Run(args []string) error {

	if len(args) != 1 {
		panic("invalid arguments for the cli")
	}

	ctx := context.Background()

	if err := pkg.LoadEnv(); err != nil {
		return err
	}

	// userPrompt := args[1]

	reader := bufio.NewReader(os.Stdin)

	infClient, err := inference.NewInferenceClient(ctx, inference.ProviderGemini)
	if err != nil {
		return err
	}

	for {
		fmt.Print("USER: ")

		currUserPrompt, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("error reading user prompt: %w", err)
		}

		if currUserPrompt == "exit\n" {
			return nil
		}

		if err := infClient.Stream(ctx, "gemini-2.5-flash", currUserPrompt); err != nil {
			return err
		}

	}

	// return nil
}
