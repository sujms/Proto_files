package server

import (
	"log"
	"net"
	"os"
	"testing"

	pb "github.com/senyosimpson/tutorials/grokkingrpc/helloworld"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func testserver() {
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
func TestMain(m *testing.M) {
	go testserver()
	os.Exit(m.Run())
}
