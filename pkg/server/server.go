package server

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"net"

	"github.com/jwolski2/nil-extended/pkg/crypto"
	"github.com/jwolski2/nil-extended/pkg/proto"
	"google.golang.org/grpc"
)

func Start(port uint, params *crypto.Params) error {
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		return fmt.Errorf("Failed to create listener on port %d: %w", port, err)
	}

	// Create server.
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	proto.RegisterAuthServer(grpcServer, NewAuthServer(params))

	// Start listening/accepting requests.
	err = grpcServer.Serve(listener)
	if err != nil {
		return fmt.Errorf("Failed to start serving requests: %w", err)
	}

	return nil
}

type AuthServer struct {
	proto.UnimplementedAuthServer
	params  *crypto.Params
	storage *storage
}

func NewAuthServer(params *crypto.Params) *AuthServer {
	return &AuthServer{
		params:  params,
		storage: newStorage(),
	}
}

func (s *AuthServer) CreateAuthenticationChallenge(ctx context.Context, req *proto.AuthenticationChallengeRequest) (*proto.AuthenticationChallengeResponse, error) {
	if !s.storage.hasUser(req.User) {
		return nil, errors.New("user does not exist")
	}

	c, err := crypto.GenerateC()
	if err != nil {
		return nil, fmt.Errorf("Failed to generate c: %w", err)
	}

	authID, err := crypto.RandomInt(256)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate auth ID: %w", err)
	}

	// We've now got all the data to record the challenge.
	//
	// TODO: Challenges should ultimately expire...
	ch := &challenge{
		authID: authID,
		c:      c,
		r1:     req.R1,
		r2:     req.R2,
		user:   req.User,
	}
	err = s.storage.storeChallenge(ch)
	if err != nil {
		return nil, fmt.Errorf("Failed to record challenge: %w", err)
	}

	return &proto.AuthenticationChallengeResponse{
		AuthId: ch.authIDStr(),
		C:      ch.c.Int64(),
	}, nil
}

func (s *AuthServer) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	// Attempt to store user in in-memory storage.
	err := s.storage.store(req.User, &ys{req.Y1, req.Y2})
	if err != nil {
		return nil, fmt.Errorf("Failed to register: %w", err)
	}

	return &proto.RegisterResponse{}, nil
}

func (s *AuthServer) VerifyAuthentication(ctx context.Context, req *proto.AuthenticationAnswerRequest) (*proto.AuthenticationAnswerResponse, error) {
	ch, ok := s.storage.activeChallenges[authID(req.AuthId)]
	if !ok {
		return nil, errors.New("Active challenge not found")
	}

	ys, ok := s.storage.registeredUsers[userID(ch.user)]
	if !ok {
		return nil, errors.New("Registered user not found for active challenge")
	}

	isVerified := crypto.VerifyR1AndR2(
		s.params,
		big.NewInt(ch.r1),
		big.NewInt(ch.r2),
		big.NewInt(req.S),
		ch.c,
		big.NewInt(ys.y1),
		big.NewInt(ys.y2),
	)

	if !isVerified {
		return nil, errors.New("r1 and r2 could not be verified")
	}

	// Generate session ID.
	id, err := crypto.RandomInt(128)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate sessionID: %w", err)
	}

	idStr := fmt.Sprintf("%x", id)

	if err := s.storage.storeSession(userID(ch.user), sessionID(idStr)); err != nil {
		return nil, fmt.Errorf("Failed to store session: %w", err)
	}

	return &proto.AuthenticationAnswerResponse{
		SessionId: idStr,
	}, nil
}
