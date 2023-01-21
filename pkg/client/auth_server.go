package client

import (
	"context"
	"fmt"

	"github.com/jwolski2/nil-extended/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RemoteAuthServer struct {
	proto.UnimplementedAuthServer

	authClient proto.AuthClient
	conn       *grpc.ClientConn
}

func NewRemoteAuthServer(hostname string, port uint) (*RemoteAuthServer, error) {
	// Instantiate client.
	conn, err := grpc.Dial(
		// Server address.
		fmt.Sprintf("%s:%d", hostname, port),

		// Client options.
		[]grpc.DialOption{
			// Insecure TLS. Sorry, it was a time-saver :|
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		}...,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to dial gRPC server: %w", err)
	}

	return &RemoteAuthServer{
		authClient: proto.NewAuthClient(conn),
		conn:       conn,
	}, nil
}

func (s *RemoteAuthServer) Close() {
	s.conn.Close()
}

func (s *RemoteAuthServer) CreateAuthenticationChallenge(ctx context.Context, req *proto.AuthenticationChallengeRequest) (*proto.AuthenticationChallengeResponse, error) {
	return s.authClient.CreateAuthenticationChallenge(ctx, req, []grpc.CallOption{}...)
}

func (s *RemoteAuthServer) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	return s.authClient.Register(ctx, req, []grpc.CallOption{}...)
}

func (s *RemoteAuthServer) VerifyAuthentication(ctx context.Context, req *proto.AuthenticationAnswerRequest) (*proto.AuthenticationAnswerResponse, error) {
	return s.authClient.VerifyAuthentication(ctx, req, []grpc.CallOption{}...)
}
