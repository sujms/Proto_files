package server

import (
	"log"
	"net"
	"testing"

	pb "github.com/senyosimpson/tutorials/grokkingrpc/helloworld"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

func TestServer(t *testing.T) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterHelloWorldServer(server, &helloWorldServer{})
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)

	}
}
