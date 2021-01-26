package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"

	"github.com/grpc-go-course/unary_grpc/squareroot/protocolbuffer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

func (*server) SquareRoot(ctx context.Context, req *protocolbuffer.SquareRootRequest) (*protocolbuffer.SquareRootResponse, error) {
	log.Printf("Server function SquareRoot() RPC was invoked\nwith request: ==> %v\n", req)

	number := req.GetNumber()

	if number < 0 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Received a negative number: %v", number))
	}

	return &protocolbuffer.SquareRootResponse{
		Result: math.Sqrt(float64(number)),
	}, nil
}

func main() {

	log.Println("Calculator Server - Suqare Root")

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
