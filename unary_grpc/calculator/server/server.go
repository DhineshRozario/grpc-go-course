package main

import (
	"context"
	"log"
	"net"

	"github.com/grpc-go-course/calculator/protocolbuffer"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) SumOfNumbers(ctx context.Context, req *protocolbuffer.CalculatorRequest) (*protocolbuffer.CalculatorResponse, error) {
	log.Printf("Server function SumOfNumbers() was invoked\nwith request: ==> %v\n", req)

	result := req.GetCalculator().GetFirstNumber() + req.GetCalculator().GetSecondNumber()

	res := protocolbuffer.CalculatorResponse{
		Result: result,
	}

	return &res, nil
}

func main() {

	log.Println("Calculator Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	s := grpc.NewServer()

	protocolbuffer.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
