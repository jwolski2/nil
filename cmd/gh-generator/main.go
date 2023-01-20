package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

func main() {
	randBytes := make([]byte, 32)
	rand.Read(randBytes)

	// Generate g.
	gChecksum := sha256.Sum256(randBytes)
	g := new(big.Int).SetBytes(gChecksum[:])

	// Generate h.
	hChecksum := sha256.Sum256(gChecksum[:])
	h := new(big.Int).SetBytes(hChecksum[:])

	fmt.Printf("g: %s\n", g.String())
	fmt.Printf("h: %s\n", h.String())
}
