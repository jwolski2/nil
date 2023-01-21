package client

import (
	"context"
	"fmt"
	"math/big"

	"github.com/jwolski2/nil-extended/pkg/crypto"
	"github.com/jwolski2/nil-extended/pkg/proto"
)

type Client struct {
	AuthServer proto.AuthServer
	Params     *crypto.Params
}

func (c *Client) Login(user string, secret *big.Int) (string, error) {
	r1, r2, k, err := crypto.GenerateR1AndR2(c.Params)
	if err != nil {
		return "", fmt.Errorf("Failed to generate r1 and r2: %w", err)
	}

	// Send commitment.
	challengeResp, err := c.AuthServer.CreateAuthenticationChallenge(context.Background(), &proto.AuthenticationChallengeRequest{
		User: user,
		R1:   r1.Int64(),
		R2:   r2.Int64(),
	})
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
	answerResp, err := c.AuthServer.VerifyAuthentication(context.Background(), &proto.AuthenticationAnswerRequest{
		AuthId: challengeResp.AuthId,
		S:      s.Int64(),
	})
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

	// Send registration request.
	_, err = c.AuthServer.Register(context.Background(), &proto.RegisterRequest{
		User: user,
		Y1:   y1.Int64(),
		Y2:   y2.Int64(),
	})
	if err != nil {
		return fmt.Errorf("Failed to register with auth server: %w", err)
	}

	return nil
}
