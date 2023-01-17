package main

import (
	"fmt"
	"log"
	"net"

	"github.com/jwolski2/zkp-extended/pkg/server"
	"github.com/jwolski2/zkp-extended/proto"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9999))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	proto.RegisterAuthServer(grpcServer, &server.AuthServer{})
	grpcServer.Serve(listener)
}
