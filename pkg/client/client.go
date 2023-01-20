package client

import (
	"context"
	"fmt"
	"math/big"

	"github.com/jwolski2/nil-extended/pkg/crypto"
	"github.com/jwolski2/nil-extended/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Login(serverHostname string, serverPort int, user string, secret *big.Int) error {
	return nil
}

func Register(serverHostname string, serverPort int, user string, secret *big.Int) error {
	y1, y2, err := crypto.GenerateY1AndY2(secret)
	if err != nil {
		return fmt.Errorf("Failed to generate y1 and y2: %w", err)
	}

	// Instantiate client.
	conn, err := grpc.Dial(
		// Server address.
		fmt.Sprintf("%s:%d", serverHostname, serverPort),

		// Client options.
		[]grpc.DialOption{
			// Insecure TLS. Sorry, it was a time-saver :|
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		}...,
	)
	if err != nil {
		return fmt.Errorf("Failed to dial gRPC server: %w", err)
	}

	// Remember to close open connection.
	defer conn.Close()

	// Instantiate client.
	authClient := proto.NewAuthClient(conn)

	// Send registration request.
	_, err = authClient.Register(context.Background(), &proto.RegisterRequest{
		User: user,
		Y1:   y1.Int64(),
		Y2:   y2.Int64(),
	}, []grpc.CallOption{}...)
	if err != nil {
		return fmt.Errorf("Failed to register with auth server: %w", err)
	}

	return nil
}
