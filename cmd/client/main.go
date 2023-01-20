package main

import (
	"errors"
	"fmt"
	"math/big"
	"os"

	"github.com/jwolski2/nil-extended/pkg/client"
	"github.com/urfave/cli/v2"
)

const (
	serverHostname = "nil-server"
	serverPort     = 9999
)

func login(ctx *cli.Context) error {
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
	err := client.Login(serverHostname, serverPort, user, secretBigInt)
	if err != nil {
		return fmt.Errorf("Failed to register user: %w", err)
	}

	fmt.Println("Successfully registered with auth server!")

	return nil
}

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
	err := client.Register(serverHostname, serverPort, user, secretBigInt)
	if err != nil {
		return fmt.Errorf("Failed to register user: %w", err)
	}

	fmt.Println("Successfully registered with auth server!")

	return nil
}

func main() {
	app := &cli.App{
		Name:  "nil-client",
		Usage: "A CLI for the Nil server",
		Commands: []*cli.Command{
			{
				Name:      "login",
				ArgsUsage: "user secret",
				Usage:     "Login with the Nil server",
				Action:    login,
			},
			{
				Name:      "register",
				ArgsUsage: "user secret",
				Usage:     "Register with the Nil server",
				Action:    register,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
