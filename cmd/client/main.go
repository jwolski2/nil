package main

import (
	"errors"
	"fmt"
	"math/big"
	"os"

	"github.com/jwolski2/nil-extended/pkg/client"
	"github.com/jwolski2/nil-extended/pkg/crypto"
	"github.com/urfave/cli/v2"
)

const (
	defaultHostname = "nil-server"
	defaultPort     = 9999
)

func login(ctx *cli.Context) error {
	// Get positional args.
	user, secret, err := validatePositionalArgs(ctx)
	if err != nil {
		return fmt.Errorf("error: could not validate positional args: %w", err)
	}

	// Validate options.
	params, err := crypto.Load(ctx.String("params-file"))
	if err != nil {
		return fmt.Errorf("Failed to load params from disk: %w", err)
	}

	// Login user.
	sessionID, err := client.Login(
		ctx.String("hostname"),
		ctx.Uint("port"),
		params,
		user,
		secret,
	)
	if err != nil {
		return errors.New("Login unsuccessful.")
	}

	fmt.Println(fmt.Sprintf("Login successful. Session ID is %s.", sessionID))

	return nil
}

func register(ctx *cli.Context) error {
	// Get positional args.
	user, secret, err := validatePositionalArgs(ctx)
	if err != nil {
		return fmt.Errorf("error: could not validate positional args: %w", err)
	}

	// Validate options.
	params, err := crypto.Load(ctx.String("params-file"))
	if err != nil {
		return fmt.Errorf("Failed to load params from disk: %w", err)
	}

	// Register user.
	err = client.Register(
		ctx.String("hostname"),
		ctx.Uint("port"),
		params,
		user,
		secret,
	)
	if err != nil {
		return fmt.Errorf("Failed to register user: %w", err)
	}

	fmt.Println("User has been registered!")

	return nil
}

func validatePositionalArgs(ctx *cli.Context) (string, *big.Int, error) {
	args := ctx.Args()

	if args.Len() != 2 {
		return "", nil, errors.New("missing positional args")
	}

	user := args.Get(0)
	if len(user) == 0 {
		return "", nil, errors.New("user cannot be empty")
	}

	secret := args.Get(1)
	if len(secret) == 0 {
		return "", nil, errors.New("secret cannot be empty")
	}

	secretBigInt := new(big.Int)
	secretBigInt, ok := secretBigInt.SetString(secret, 10)
	if !ok {
		return "", nil, errors.New("secret cannot be converted to bigint")
	}

	return user, secretBigInt, nil
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
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "params-file",
						Value: crypto.DefaultParamsFile,
						Usage: "Public params to use between client/server",
					},
				},
			},
			{
				Name:      "register",
				ArgsUsage: "user secret",
				Usage:     "Register with the Nil server",
				Action:    register,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "params-file",
						Value: crypto.DefaultParamsFile,
						Usage: "Public params to use between client/server",
					},
				},
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "hostname",
				Value: defaultHostname,
				Usage: "Hostname for nil-server",
			},
			&cli.UintFlag{
				Name:  "port",
				Value: defaultPort,
				Usage: "Port for nil-server",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
