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

type Client struct {
	Hostname string
	Port     uint
	Params   *crypto.Params
}

func (c *Client) Login(user string, secret *big.Int) (string, error) {
	r1, r2, k, err := crypto.GenerateR1AndR2(c.Params)
	if err != nil {
		return "", fmt.Errorf("Failed to generate r1 and r2: %w", err)
	}

	// Instantiate client.
	conn, err := grpc.Dial(
		// Server address.
		fmt.Sprintf("%s:%d", c.Hostname, c.Port),

		// Client options.
		[]grpc.DialOption{
			// Insecure TLS. Sorry, it was a time-saver :|
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		}...,
	)
	if err != nil {
		return "", fmt.Errorf("Failed to dial gRPC server: %w", err)
	}

	// Remember to close open connection.
	defer conn.Close()

	// Instantiate client.
	authClient := proto.NewAuthClient(conn)

	// Send commitment.
	challengeResp, err := authClient.CreateAuthenticationChallenge(context.Background(), &proto.AuthenticationChallengeRequest{
		User: user,
		R1:   r1.Int64(),
		R2:   r2.Int64(),
	}, []grpc.CallOption{}...)
	if err != nil {
		return "", fmt.Errorf("Failed to create authentication challenge: %w", err)
	}

	// Generate s for verification.
	s, err := crypto.GenerateS(
		c.Params,
		secret,                      // x
		k,                           // from r1/r2 calc,
		big.NewInt(challengeResp.C), // challenge computed by server
	)
	if err != nil {
		return "", fmt.Errorf("Failed to generate s: %w", err)
	}

	// Send challenge answer.
	answerResp, err := authClient.VerifyAuthentication(context.Background(), &proto.AuthenticationAnswerRequest{
		AuthId: challengeResp.AuthId,
		S:      s.Int64(),
	}, []grpc.CallOption{}...)
	if err != nil {
		return "", fmt.Errorf("Failed to verify authentication: %w", err)
	}

	return answerResp.SessionId, nil
}

func (c *Client) Register(user string, secret *big.Int) error {
	y1, y2, err := crypto.GenerateY1AndY2(c.Params, secret)
	if err != nil {
		return fmt.Errorf("Failed to generate y1 and y2: %w", err)
	}

	// Instantiate client.
	conn, err := grpc.Dial(
		// Server address.
		fmt.Sprintf("%s:%d", c.Hostname, c.Port),

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
