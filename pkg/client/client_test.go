package client

import (
	"testing"

	"github.com/jwolski2/nil-extended/pkg/server"
)

func Test1(t *testing.T) {
	go server.Start(0, params)
}
