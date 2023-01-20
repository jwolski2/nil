package main

import (
	"fmt"
	"os"

	"github.com/jwolski2/nil-extended/pkg/server"
	"github.com/urfave/cli/v2"
)

const (
	port = 9999
)

func start(ctx *cli.Context) error {
	err := server.Start(port)
	if err != nil {
		return fmt.Errorf("Failed to start server: %w", err)
	}

	return nil
}

func main() {
	app := &cli.App{
		Name:  "nil-server",
		Usage: "A Nil server",
		Commands: []*cli.Command{
			{
				Name:   "start",
				Usage:  "Start the server",
				Action: start,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
