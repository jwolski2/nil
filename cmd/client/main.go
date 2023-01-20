package main

import (
	"errors"
	"fmt"
	"math/big"
	"os"

	"github.com/jwolski2/zkp-extended/pkg/client"
	"github.com/urfave/cli/v2"
)

func register(ctx *cli.Context) error {
	args := ctx.Args()

	// Validate inputs.
	if args.Len() != 2 {
		return errors.New("error: missing arguments. See --help.")
	}

	user := args.Get(0)
	if len(user) == 0 {
		return errors.New("error: user cannot be empty")
	}

	secret := args.Get(1)
	if len(secret) == 0 {
		return errors.New("error: secret cannot be empty")
	}

	secretBigInt := new(big.Int)
	secretBigInt, ok := secretBigInt.SetString(secret, 10)
	if !ok {
		return errors.New("error: secret cannot be converted to bigint")
	}

	// Register user.
	err := client.Register(user, secretBigInt)
	if err != nil {
		return fmt.Errorf("Failed to register user: %w", err)
	}

	return nil
}

func main() {
	app := &cli.App{
		Name:  "zkp-client",
		Usage: "A CLI for the ZKP server",
		Commands: []*cli.Command{
			{
				Name:      "register",
				ArgsUsage: "user secret",
				Usage:     "Register with the ZKP server",
				Action:    register,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
