package main

import (
	"context"
	"log"

	"github.com/grpc-go-course/unary_grpc/greet/greetpb"
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
	doUnary(c)
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
