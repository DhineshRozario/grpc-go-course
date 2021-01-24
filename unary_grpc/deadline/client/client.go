package main

import (
	"context"
	"log"
	"time"

	"github.com/grpc-go-course/unary_grpc/deadline/protocolbuffer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	log.Println("Greeting Bi-Directional Streaming Client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Not able to connect with server: %v", err)
	}

	defer conn.Close()

	c := protocolbuffer.NewGreetServiceClient(conn)
	doGreetWithDeadline(c, 5*time.Second) // Proper result
	doGreetWithDeadline(c, 1*time.Second) // Timeout
}

func doGreetWithDeadline(c protocolbuffer.GreetServiceClient, timeout time.Duration) error {

	log.Println("Starting to do Greet With Deadline() - Unary RPC...")

	req := &protocolbuffer.GreetWithDeadlineRequest{
		Greeting: &protocolbuffer.Greeting{
			FirstName: "Dhinesh",
			LastName:  "Rozario",
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel() // cancelling the cancel call.

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {

		respError, ok := status.FromError(err)

		if ok {
			//Actual error from GRPC (user Error)
			log.Printf("User error: %v and code: %v", respError.Message(), respError.Code())

			if respError.Code() == codes.DeadlineExceeded {
				log.Printf("Timeout was hit, deadline exceeded)")
			} else {
				//Framework Error - Big error
				log.Printf("Unexpected error GreetWithDeadline RPC: %v", respError)
			}
		} else {
			//Framework Error - Big error
			log.Fatalf("Error while calling GreetWithDeadline RPC: %v", err)
		}
		return err
	}

	log.Printf("Response from GreetWithDeadline Server: %v", res.Result)

	return nil
}
