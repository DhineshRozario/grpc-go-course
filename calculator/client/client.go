package main

import (
	"context"
	"log"

	"github.com/grpc-go-course/calculator/protocolbuffer"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Hello I'm a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer conn.Close()

	c := protocolbuffer.NewCalculatorServiceClient(conn)
	doCalculate(c)
}

func doCalculate(c protocolbuffer.CalculatorServiceClient) {

	log.Println("Starting the client for Adding two numbers in Unary RPC")

	req := &protocolbuffer.CalculatorRequest{
		Calculator: &protocolbuffer.Calculator{
			FirstNumber:  3,
			SecondNumber: 10,
		},
	}

	res, err := c.SumOfNumbers(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Add Service: %v", err)
	}

	log.Printf("Response from Calculator Add Sercer: %v", res.Result)
}
