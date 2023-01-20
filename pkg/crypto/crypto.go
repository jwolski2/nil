package crypto

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var (
	// For now, hard-coded. Later, could be injected at run-time.
	g, _ = new(big.Int).SetString("27389535407172064518385090712503604205081709236897638041171802294068595118554", 10)
	h, _ = new(big.Int).SetString("82790680804658874823236941014605029406785332510930358145770704050898063577918", 10)
)

func GenerateY1AndY2(x *big.Int) (*big.Int, *big.Int, error) {
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
