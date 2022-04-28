package client

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/senyosimpson/tutorials/grokkingrpc/helloworld"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func Client() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewHelloWorldClient(conn)

	// Contact the server and print out its response.
	name := "sujith"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.Greet(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
