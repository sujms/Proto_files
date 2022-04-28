package client

import (
	"context"
	"testing"

	pb "github.com/senyosimpson/tutorials/grokkingrpc/helloworld"
	"google.golang.org/grpc"
)

func clienttest(t *testing.T) {

	// Set up a connection to the Server.
	const address = "localhost:8080"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewHelloWorldClient(conn)

	// Test greet
	t.Run("greet", func(t *testing.T) {
		name := "sujith"
		r, err := client.Greet(context.Background(), &pb.HelloRequest{Name: name})
		if err != nil {
			t.Fatalf("could not greet: %v", err)
		}
		t.Logf("Greeting: %s", r.Message)
		if r.Message != "Hello "+name {
			t.Error("Expected 'Hello sujith', got ", r.Message)
		}

	})
}
