package server

import (
	"context"

	"github.com/jwolski2/zkp-extended/proto"
)

type AuthServer struct {
	proto.UnimplementedAuthServer
}

func (s *AuthServer) CreateAuthenticationChallenge(ctx context.Context, req *proto.AuthenticationChallengeRequest) (*proto.AuthenticationChallengeResponse, error) {
	return nil, nil
}

func (s *AuthServer) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	return nil, nil
}

func (s *AuthServer) VerifyAuthentication(ctx context.Context, req *proto.AuthenticationAnswerRequest) (*proto.AuthenticationAnswerResponse, error) {
	return nil, nil
}
