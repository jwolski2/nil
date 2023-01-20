package crypto

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var (
	// For now, hard-coded. Later, could be injected at run-time.
	g = big.NewInt(2)
	h = big.NewInt(3)
	p = big.NewInt(11)
	q = big.NewInt(7)
)

func GenerateC() (*big.Int, error) {
	return Random256Bit()
}

func GenerateR1AndR2() (*big.Int, *big.Int, *big.Int, error) {
	// Generate k.
	k, err := rand.Prime(rand.Reader, 256)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("Failed to generate k: %w", err)
	}

	// Generate r1: g^k mod p
	r1 := new(big.Int).Exp(g, k, p)

	// Generate r2: h^k mod p
	r2 := new(big.Int).Exp(h, k, p)

	return r1, r2, k, nil
}

func GenerateS(x, k, c *big.Int) (*big.Int, error) {
	s := new(big.Int)
	s.Mul(c, x)
	s.Sub(k, s)
	s.Mod(s, q)

	return s, nil
}

func GenerateY1AndY2(x *big.Int) (*big.Int, *big.Int, error) {
	// Generate y1: g^x mod p
	y1 := new(big.Int).Exp(g, x, p)

	// Generate y2: h^x mod p
	y2 := new(big.Int).Exp(h, x, p)

	return y1, y2, nil
}

func Random256Bit() (*big.Int, error) {
	buffer := make([]byte, 32)
	_, err := rand.Read(buffer)
	if err != nil {
		return nil, fmt.Errorf("Failed to read random 256-bit number into buffer: %w", err)
	}

	return new(big.Int).SetBytes(buffer), nil
}
