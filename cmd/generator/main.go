package main

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path"

	"github.com/jwolski2/nil-extended/pkg/crypto"
	"github.com/urfave/cli/v2"
)

func genParams() *crypto.Params {
	var p, g, h, q int64

	for {
		// Generate p. 32-bits is small enough to make computationally
		// practical, but certainly not large enough for production purposes.
		bigP, err := rand.Prime(rand.Reader, 32)
		if err != nil {
			fmt.Fprintln(os.Stderr, fmt.Errorf("Failed to generate p: %w", err))
			os.Exit(1)
		}

		p = bigP.Int64()

		// Find possible qs.
		dividesEvenlyAndPrime := []int64{}
		for i := int64(2); i < p-1; i++ {
			if (p-1)%i == 0 && big.NewInt(i).ProbablyPrime(10) {
				dividesEvenlyAndPrime = append(dividesEvenlyAndPrime, i)
			}
		}

		if len(dividesEvenlyAndPrime) == 0 {
			continue
		}

		// Find g.
		for i := int64(2); i < p; i++ {
			for _, j := range dividesEvenlyAndPrime {
				// i^j mod p == 1
				if new(big.Int).Exp(big.NewInt(i), big.NewInt(j), big.NewInt(p)).Cmp(big.NewInt(1)) == 0 &&
					// i is relatively prime to p.
					new(big.Int).GCD(nil, nil, big.NewInt(i), big.NewInt(p)).Cmp(big.NewInt(1)) == 0 {
					g = i
					q = j
					goto findh
				}
			}
		}

	findh:
		// Find h.
		for i := int64(2); i < p; i++ {
			// i^q mod p == 1
			if new(big.Int).Exp(big.NewInt(i), big.NewInt(q), big.NewInt(p)).Cmp(big.NewInt(1)) == 0 &&
				// h is relatively prime to p.
				new(big.Int).GCD(nil, nil, big.NewInt(i), big.NewInt(p)).Cmp(big.NewInt(1)) == 0 &&
				// h is not equal to g.
				i != g {
				h = i
				goto exit
			}
		}

	exit:
		if g != 0 && h != 0 && q != 0 {
			break
		}

		// Reset.
		g = 0
		h = 0
		q = 0
	}

	return &crypto.Params{p, g, h, q}
}

func gen(ctx *cli.Context) error {
	count := ctx.Uint("count")
	outDir := ctx.String("out-dir")

	// Validate out-dir by ensuring the directory exists.
	fileInfo, err := os.Stat(outDir)
	if err != nil {
		return fmt.Errorf("Failed to stat out-dir: %w", err)
	}

	if !fileInfo.IsDir() {
		return errors.New("error: out-dir is not a directory")
	}

	// Start generating params sets by writing JSON files to, e.g.:
	//
	//   data/
	//     params1.json
	//     params2.json
	//     paramsN.json
	for i := uint(1); i <= count; i++ {
		params := genParams()

		// Convert params to JSON.
		paramsAsJson, err := json.MarshalIndent(params, "", " ")
		if err != nil {
			return fmt.Errorf("Failed to marshal JSON data: %w", err)
		}

		// Write params to disk.
		filename := path.Join(outDir, fmt.Sprintf("params%d.json", i))
		if err := ioutil.WriteFile(filename, paramsAsJson, 0644); err != nil {
			return fmt.Errorf("Failed to write params file: %w", err)
		}
	}

	return nil
}

func main() {
	app := &cli.App{
		Name:  "nil-generator",
		Usage: "A Nil generator",
		Commands: []*cli.Command{
			{
				Name:   "gen",
				Usage:  "Generates public params: p, g, h, q",
				Action: gen,
				Flags: []cli.Flag{
					&cli.UintFlag{
						Name:  "count",
						Value: 5,
						Usage: "Number of param files generated",
					},
					&cli.StringFlag{
						Name:  "out-dir",
						Value: "data",
						Usage: "Data dir for param files",
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
