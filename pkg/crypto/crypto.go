package crypto

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var (
	// For now, hard-coded. Later, could be injected at run-time.
	p = big.NewInt(55903)
	g = big.NewInt(35)
	h = big.NewInt(1225)
	q = big.NewInt(11)
)

func GenerateC() (*big.Int, error) {
	return RandomInt(8)
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

func RandomInt(bits int) (*big.Int, error) {
	buffer := make([]byte, bits/8)
	_, err := rand.Read(buffer)
	if err != nil {
		return nil, fmt.Errorf("Failed to read random number into buffer: %w", err)
	}

	return new(big.Int).SetBytes(buffer), nil
}

func VerifyR1AndR2(r1, r2, s, c, y1, y2 *big.Int) bool {
	fmt.Printf("r1: %d, r2: %d, s: %d, c: %d, y1: %d, y2: %d\n", r1, r2, s, c, y1, y2)

	term1 := new(big.Int).Exp(g, s, p)
	term2 := new(big.Int).Exp(y1, c, p)
	r1Cmp := new(big.Int).Mul(term1, term2)
	r1Cmp.Mod(r1Cmp, p)

	term1 = new(big.Int).Exp(h, s, p)
	term2 = new(big.Int).Exp(y2, c, p)
	r2Cmp := new(big.Int).Mul(term1, term2)
	r2Cmp.Mod(r2Cmp, p)

	fmt.Printf("r1Cmp: %d, r2Cmp: %d\n", r1Cmp, r2Cmp)

	return r1.Cmp(r1Cmp) == 0 &&
		r2.Cmp(r2Cmp) == 0
}
