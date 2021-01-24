package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/grpc-go-course/unary_grpc/deadline/protocolbuffer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

func (*server) GreetWithDeadline(ctx context.Context, req *protocolbuffer.GreetWithDeadlineRequest) (*protocolbuffer.GreetWithDeadlineResponse, error) {
	log.Printf("Server function GreetWithDeadline() RPC was invoked\nwith request: ==> %v\n", req)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			log.Fatalf("The Client cancelled the request%v\n", ctx.Err())
			return nil, status.Error(codes.Canceled, "The Client cancelled the request")
		}
		time.Sleep(1 * time.Second)
	}
	firstName := req.GetGreeting().GetFirstName()

	result := "Hello " + firstName

	res := &protocolbuffer.GreetWithDeadlineResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	log.Println("Greet Server With Deadline!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	s := grpc.NewServer()

	protocolbuffer.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
