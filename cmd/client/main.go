package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func login(ctx *cli.Context) error {
	os.Exit(1)
	return nil
}

func main() {
	app := &cli.App{
		Name:  "zkp-client",
		Usage: "A CLI for the ZKP server",
		Commands: []cli.Command{
			{
				Name:   "login",
				Usage:  "Login to the ZKP server",
				Action: login,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
