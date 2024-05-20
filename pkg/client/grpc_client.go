package client

import (
	"flag"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr    = flag.String("addr", "localhost:50051", "the address to connect to")
	topic   = flag.String("topic", "test-topic", "topic")
	message = flag.String("message", "test-message", "message")
)

func GrpcClient() grpc.ClientConnInterface {
	serverAddress := "localhost:50051"
	creds := insecure.NewCredentials()
	opts := grpc.WithTransportCredentials(creds)

	conn, err := grpc.Dial(serverAddress, opts)

	if err != nil {
		panic(err)
	}

	fmt.Println("initiated grpc client to localhost:50051")
	return conn
}
