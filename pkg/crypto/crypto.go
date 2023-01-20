package crypto

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

func GenerateY1AndY2(x *big.Int) (*big.Int, *big.Int, error) {
	randBytes := make([]byte, 32)
	rand.Read(randBytes)

	// Generate g.
	gChecksum := sha256.Sum256(randBytes)
	g := new(big.Int).SetBytes(gChecksum[:])

	// Generate h.
	hChecksum := sha256.Sum256(gChecksum[:])
	h := new(big.Int).SetBytes(hChecksum[:])

	// Generate p.
	p, err := rand.Prime(rand.Reader, 16)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to generate p: %w", err)
	}

	// Generate y1: g^x mod p
	y1 := new(big.Int).Mod(new(big.Int).Exp(g, x, nil), p)

	// Generate y2: h^x mod p
	y2 := new(big.Int).Mod(new(big.Int).Exp(h, x, nil), p)

	return y1, y2, nil
}
