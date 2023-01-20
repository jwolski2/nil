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
	p, _ = new(big.Int).SetString("109803578352529518575151709482869916174120725793701284668325439525881006011871", 10)
)

func GenerateR1AndR2() (*big.Int, *big.Int, error) {
	// Generate k.
	k, err := rand.Prime(rand.Reader, 256)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to generate k: %w", err)
	}

	// Generate r1: g^k mod p
	r1 := new(big.Int).Mod(new(big.Int).Exp(g, k, nil), k)

	// Generate r2: h^k mod p
	r2 := new(big.Int).Mod(new(big.Int).Exp(h, k, nil), k)

	return r1, r2, nil
}

func GenerateY1AndY2(x *big.Int) (*big.Int, *big.Int, error) {
	// Generate y1: g^x mod p
	y1 := new(big.Int).Mod(new(big.Int).Exp(g, x, nil), p)

	// Generate y2: h^x mod p
	y2 := new(big.Int).Mod(new(big.Int).Exp(h, x, nil), p)

	return y1, y2, nil
}
