package server

import (
	"fmt"
	"math/big"
)

type authID string
type userID string

type challenge struct {
	authID *big.Int
	c      *big.Int
	r1     int64
	r2     int64
	user   string
}

func (ch *challenge) authIDStr() string {
	return fmt.Sprintf("%x", ch.authID)
}

type ys struct {
	y1 int64
	y2 int64
}
