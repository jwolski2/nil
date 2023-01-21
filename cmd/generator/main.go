package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

func main() {
	var p, g, h, q int64

	for {
		// Generate p.
		bigP, err := rand.Prime(rand.Reader, 16)
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
				if new(big.Int).Exp(big.NewInt(i), big.NewInt(j), big.NewInt(p)).Cmp(big.NewInt(1)) == 0 {
					g = i
					q = j
					goto findh
				}
			}
		}

	findh:
		// Find h.
		for i := int64(2); i < p; i++ {
			for _, j := range dividesEvenlyAndPrime {
				if new(big.Int).Exp(big.NewInt(i), big.NewInt(j), big.NewInt(p)).Cmp(big.NewInt(1)) == 0 && i != g {
					h = i
					goto exit
				}
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

	fmt.Printf("p: %d\n", p)
	fmt.Printf("g: %d\n", g)
	fmt.Printf("h: %d\n", h)
	fmt.Printf("q: %d\n", q)
}
