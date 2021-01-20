package main

import (
	"context"
	"log"
	"net"

	"github.com/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {

	log.Printf("Greet function was invoked with request :%v\n", req)

	firstName := req.GetGreeting().GetFirstName()

	result := "Hello " + firstName

	res := greetpb.GreetResponse{
		Result: result,
	}

	return &res, nil
}

func main() {
	log.Println("Hello World - am a Server!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
