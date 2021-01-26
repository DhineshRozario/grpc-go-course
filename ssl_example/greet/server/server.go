package main

import (
	"context"
	"log"
	"net"

	"github.com/grpc-go-course/ssl_example/greet/protocolbuffer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *protocolbuffer.GreetRequest) (*protocolbuffer.GreetResponse, error) {
	log.Printf("Server function Greet() RPC was invoked\nwith request: ==> %v\n", req)

	firstName := req.GetGreeting().GetFirstName()

	result := "Hello " + firstName

	res := &protocolbuffer.GreetResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	log.Println("Greet Server With SSL!")
	tls := true

	opts := []grpc.ServerOption{}

	if tls {
		certFile := "./cert/server.crt"
		keyFile := "./cert/server.key"

		//Not working
		// certFile := "./old_cert/server.crt"
		// keyFile := "./old_cert/server.key"

		creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)

		if sslErr != nil {
			log.Fatalf("Failed to loading the certificates: %v\n", sslErr)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)

	lis, err := net.Listen("tcp", "localhost:50051")
	// // error handling omitted
	// s.Serve(lis)

	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	protocolbuffer.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
