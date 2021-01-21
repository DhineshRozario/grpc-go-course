package main

import (
	"context"
	"log"

	"github.com/grpc-go-course/server_streaming_grpc/calculator/protocolbuffer"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Calculator client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Not able to connect with server: %v", err)
	}

	defer conn.Close()

	c := protocolbuffer.NewCalculatorServiceClient(conn)
	doCalculate(c)
}

func doCalculate(c protocolbuffer.CalculatorServiceClient) {

	log.Println("Starting the client for sum of two numbers in Unary RPC")

	req := &protocolbuffer.CalculatorRequest{
		Calculator: &protocolbuffer.Calculator{
			FirstNumber:  5,
			SecondNumber: 6,
		},
	}

	log.Printf("Calculating the sum of %v and %v:\n", req.GetCalculator().GetFirstNumber(), req.GetCalculator().GetSecondNumber())

	res, err := c.SumOfNumbers(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling 'Sum Of Numbers' Service: %v", err)
	}

	log.Printf("==> The response from Calculator Server -> 'Sum Of Numbers' is: %v\n", res.Result)
}
