package crypto

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
)

// Load loads a params file from disk and unmarshals it as a crypto.Params
// pointer.
func Load(paramsFile string) (*Params, error) {
	if _, err := os.Stat(paramsFile); err != nil {
		return nil, fmt.Errorf("Failed to stat params file: %w", err)
	}

	paramsJSON, err := ioutil.ReadFile(paramsFile)
	if err != nil {
		return nil, fmt.Errorf("Failed to read params file: %w", err)
	}

	params := &Params{}
	if err := json.Unmarshal(paramsJSON, params); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal params file as JSON: %w", err)
	}

	return params, nil
}

// GenerateC generates a c-value for the challenge step of the login process. It
// is kept at a 32-bit value to avoid overflow when converted to int64.
func GenerateC() (*big.Int, error) {
	return RandomInt(32)
}

// ComputeR1AndR2 computes r1 and r2 provided a random-k and the public param
// input:
//
//   r1 = g^k mod p
//   r2 = h^k mod p
func ComputeR1AndR2(params *Params) (*big.Int, *big.Int, *big.Int, error) {
	// Generate k.
	k, err := rand.Prime(rand.Reader, 256)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("Failed to generate k: %w", err)
	}

	r1 := new(big.Int).Exp(params.G(), k, params.P())
	r2 := new(big.Int).Exp(params.H(), k, params.P())
	return r1, r2, k, nil
}

// ComputeS computes an s-value for the response/answer step of the login
// process.
//
//   s = k - c * x (mod q)
func ComputeS(params *Params, x, k, c *big.Int) (*big.Int, error) {
	s := new(big.Int)
	s.Mul(c, x)
	s.Sub(k, s)
	s.Mod(s, params.Q())

	return s, nil
}

// ComputeY1AndY2 computes y1 and y2-values for the registration process.
//
//   y1 = g^x mod p
//   y2 = h^x mod p
func ComputeY1AndY2(params *Params, x *big.Int) (*big.Int, *big.Int, error) {
	y1 := new(big.Int).Exp(params.G(), x, params.P())
	y2 := new(big.Int).Exp(params.H(), x, params.P())
	return y1, y2, nil
}

// RandomInt generates a cryptographically secure random number of an arbitrary
// bit length.
func RandomInt(bits int) (*big.Int, error) {
	buffer := make([]byte, bits/8)
	_, err := rand.Read(buffer)
	if err != nil {
		return nil, fmt.Errorf("Failed to read random number into buffer: %w", err)
	}

	return new(big.Int).SetBytes(buffer), nil
}

// VerifyR1AndR2 verifies r1 and r2 by ensuring:
//
//   r1 = g^s * y1^c mod p
//   r2 = h^s * y2^c mod p
func VerifyR1AndR2(params *Params, r1, r2, s, c, y1, y2 *big.Int) bool {
	// Compute g^s * y1^c.
	term1 := new(big.Int).Exp(params.G(), s, params.P())
	term2 := new(big.Int).Exp(y1, c, params.P())
	r1Cmp := new(big.Int).Mul(term1, term2)
	r1Cmp.Mod(r1Cmp, params.P())

	// Compute h^s * y2^c.
	term1 = new(big.Int).Exp(params.H(), s, params.P())
	term2 = new(big.Int).Exp(y2, c, params.P())
	r2Cmp := new(big.Int).Mul(term1, term2)
	r2Cmp.Mod(r2Cmp, params.P())

	// Compare them against r1 and r2, respectively.
	return r1.Cmp(r1Cmp) == 0 && r2.Cmp(r2Cmp) == 0
}
