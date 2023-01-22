package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/jwolski2/nil-extended/pkg/crypto"
	"github.com/jwolski2/nil-extended/pkg/server"
	"github.com/urfave/cli/v2"
)

const (
	defaultPort = 9999
)

func start(ctx *cli.Context) error {
	port := ctx.Uint("port")

	paramsFile := ctx.String("params-file")
	if paramsFile == "" {
		return errors.New("error: params-file cannot be empty")
	}

	params, err := crypto.Load(paramsFile)
	if err != nil {
		return fmt.Errorf("Failed to load params from disk: %w", err)
	}

	if err := server.Start(port, params); err != nil {
		return fmt.Errorf("Server failed: %w", err)
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
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "params-file",
						Value: crypto.DefaultParamsFile,
						Usage: "Public params to use between client/server",
					},
					&cli.UintFlag{
						Name:  "port",
						Value: defaultPort,
						Usage: "Listening port for gRPC server",
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
