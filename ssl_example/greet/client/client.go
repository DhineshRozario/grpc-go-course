package main

import (
	"context"
	"log"
	"os"
	"runtime/trace"

	"github.com/grpc-go-course/ssl_example/greet/protocolbuffer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	log.Println("Greeting SSL Client")

	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	//Not working
	// certFile := "./old_cert/ca.crt" // Certificate Authority Trust Certificate

	certFile := "./cert/ca.crt" // Certificate Authority Trust Certificate

	creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")

	if sslErr != nil {
		log.Fatalf("Failed to loading the CA Trust certificate: %v\n", sslErr)
		return
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.Dial("localhost:50051", opts...)

	if err != nil {
		log.Fatalf("Not able to connect with server: %v", err)
	}

	defer conn.Close()

	c := protocolbuffer.NewGreetServiceClient(conn)
	doGreet(c) // Proper result
}

func doGreet(c protocolbuffer.GreetServiceClient) error {

	log.Println("Starting to do Greet() - Unary RPC...")

	req := &protocolbuffer.GreetRequest{
		Greeting: &protocolbuffer.Greeting{
			FirstName: "Dhinesh",
			LastName:  "Rozario",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		//Framework Error - Big error
		log.Fatalf("Error while calling Greet() RPC: %v", err)
		return err
	}

	log.Printf("Response from Greet Server: %v", res.Result)

	return nil
}
