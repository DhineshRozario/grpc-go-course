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

	c := protocolbuffer.NewPrimeNumberDeCompositionServiceClient(conn)
	// doUnary(c)
	getPrimeDecomposition(c)
}

func getPrimeDecomposition(c protocolbuffer.PrimeNumberDeCompositionServiceClient) {

	log.Println("Starting the getPrimeDecomposition method using the Server Streaming RPC...")

	for i := 1000; i > 1; i -= 10 {
		req := &protocolbuffer.PrimeNumberDeCompositionRequest{
			PrimeNumber: &protocolbuffer.PrimeNumberDeComposition{
				Number: int64(i),
			},
		}
		resStream, err := c.PrimeNumberDeComposition(context.Background(), req)

		if err != nil {
			log.Fatalf("Error while calling PrimeNumberDeComposition RPC: %v", err)
		}

		log.Printf("Response from PrimeNumberDeComposition Server for the number: %v\n", req.GetPrimeNumber().GetNumber())
		for {
			msg, err := resStream.Recv()

			if err == io.EOF {
				log.Println("Reached end of response...")
				break
			} else if err != nil {
				log.Fatalf("Error while receiving the response from Server: %v", err)
			}
			log.Printf("%v ", msg.GetResult())
		}
	}
}
