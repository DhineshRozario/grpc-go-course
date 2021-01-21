package main

import (
	"context"
	"io"
	"log"

	"github.com/grpc-go-course/server_streaming_grpc/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Greeting Client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Not able to connect with server: %v", err)
	}

	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	// doUnary(c)
	doServerStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {

	log.Println("Starting to do the Unary RPC...")

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Dhinesh",
			LastName:  "Rozario",
		},
	}
	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}

	log.Printf("Response from Greet Server: %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {

	log.Println("Starting to do the Server Streaming RPC...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Dhinesh",
			LastName:  "Rozario",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()

		if err == io.EOF {
			log.Println("Reached end of file...")
			break
		} else if err != nil {
			log.Fatalf("Error while receiving the response message: %v", err)
		}
		log.Printf("Response from GreetManyTimes Server: %v", msg.GetResult())
	}
}
