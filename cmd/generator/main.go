package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

func try() (*big.Int, *big.Int, *big.Int, *big.Int) {
	// Generate p.
	p, err := rand.Prime(rand.Reader, 16)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("Failed to generate p: %w", err))
		os.Exit(1)
	}

	// Generate g and h.
	g, _ := rand.Int(rand.Reader, p)
	h, _ := rand.Int(rand.Reader, p)

	gClone := new(big.Int).SetBytes(g.Bytes())
	hClone := new(big.Int).SetBytes(h.Bytes())

	// Find the primer order of g and h.
	q := big.NewInt(1)

	for ; gClone.Cmp(big.NewInt(1)) != 0 && hClone.Cmp(big.NewInt(1)) != 0; q.Add(q, big.NewInt(1)) {
		gClone.Exp(g, q, p)
		hClone.Exp(h, q, p)
	}

	return p, g, h, q
}

func main() {
	var p, g, h, q *big.Int

	for {
		p, g, h, q = try()

		// p cannot equal q.
		if p.Cmp(q) == 0 {
			continue
		}

		// q can't be less than g or h.
		if q.Cmp(g) == -1 || q.Cmp(h) == -1 {
			continue
		}

		// q must divide p - 1 evenly.
		if new(big.Int).Mod(new(big.Int).Sub(p, big.NewInt(1)), q).Cmp(big.NewInt(0)) != 0 {
			continue
		}

		// q is prime.
		if q.ProbablyPrime(10) {
			break
		}
	}

	fmt.Printf("p: %s\n", p.String())
	fmt.Printf("g: %s\n", g.String())
	fmt.Printf("h: %s\n", h.String())
	fmt.Printf("q: %s\n", q.String())
}
