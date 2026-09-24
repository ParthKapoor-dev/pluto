package main

import (
	"os"

	"github.com/parthkapoor-dev/pluto/cmd"
)

func main() {
	if err := cmd.NewCli().Run(os.Args); err != nil {
		panic(err)
	}
}
