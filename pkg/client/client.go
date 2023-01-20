package client

import (
	"fmt"
	"math/big"

	"github.com/jwolski2/nil-extended/pkg/crypto"
)

func Register(user string, secret *big.Int) error {
	_, _, err := crypto.GenerateY1AndY2(secret)
	if err != nil {
		return fmt.Errorf("Failed to generate y1 and y2: %w", err)
	}

	// TODO: Instantiate gRPC client here.

	return nil
}
