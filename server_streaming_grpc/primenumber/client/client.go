package main

import (
	"context"
	"io"
	"log"

	"github.com/grpc-go-course/server_streaming_grpc/primenumber/protocolbuffer"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Prime Number Decomposition Client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Not able to connect with server: %v", err)
	}

	defer conn.Close()

	c := protocolbuffer.NewPrimeNumberServiceClient(conn)
	// doUnary(c)
	getPrimeDecomposition(c)
}

func getPrimeDecomposition(c protocolbuffer.PrimeNumberServiceClient) {

	log.Println("Starting the getPrimeDecomposition method using the Server Streaming RPC...")

	req := &protocolbuffer.PrimeNumberRequest{
		PrimeNumber: &protocolbuffer.PrimeNumber{
			Number: 220,
		},
	}
	resStream, err := c.PrimeNumberDeComposition(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling PrimeNumberDeComposition RPC: %v", err)
	}

	log.Println("Response from PrimeNumberDeComposition Server")
	for {
		msg, err := resStream.Recv()

		if err == io.EOF {
			log.Println("Reached end of response...")
			break
		} else if err != nil {
			log.Fatalf("Error while receiving the response from Server: %v", err)
		}
		log.Printf("%v", msg.GetResult())
	}
}
