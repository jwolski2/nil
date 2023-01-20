package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"
)

func main() {
	// Generate p.
	p, err := rand.Prime(rand.Reader, 256)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("Failed to generate p: %w", err))
		os.Exit(1)
	}

	randBytes := make([]byte, 32)
	rand.Read(randBytes)

	// Generate g.
	gChecksum := sha256.Sum256(randBytes)
	g := new(big.Int).SetBytes(gChecksum[:])

	// Generate h.
	hChecksum := sha256.Sum256(gChecksum[:])
	h := new(big.Int).SetBytes(hChecksum[:])

	fmt.Printf("p: %s\n", p.String())
	fmt.Printf("g: %s\n", g.String())
	fmt.Printf("h: %s\n", h.String())
}
