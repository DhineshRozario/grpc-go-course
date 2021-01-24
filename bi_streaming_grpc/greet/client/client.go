package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/grpc-go-course/bi_streaming_grpc/greet/protocolbuffer"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Greeting Bi-Directional Streaming Client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Not able to connect with server: %v", err)
	}

	defer conn.Close()

	c := protocolbuffer.NewGreetServiceClient(conn)
	doBiDirectionalStreaming(c)
}

func doBiDirectionalStreaming(c protocolbuffer.GreetServiceClient) error {

	log.Println("Starting to do the Bi-Directional Client RPC...")

	//we create stream by invoking the client
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while calling ComputeAverage RPC: %v", err)
		return err
	}

	waitc := make(chan struct{})

	//we send a bunch of messages to the server (using go routine)
	go func() {
		//for loop to send a bunch of messages
		values := []string{"Dhinesh", "Radhika", "Dewin", "Dewiz"}
		for _, value := range values {
			req := &protocolbuffer.GreetEveryoneRequest{
				Greeting: &protocolbuffer.Greeting{
					FirstName: value,
				},
			}
			log.Printf("Sending the request: %v\n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	//we receive a bunch of messages from the server (using go routine)
	go func() {
		//for loop to receive a bunch of messages
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving GreetEveryone RPC: %v", err)
				break
			}

			log.Printf("Received from Server: %v", res.GetResult())
		}
		close(waitc)
	}()

	//Block until everything is done
	<-waitc

	return nil
}
