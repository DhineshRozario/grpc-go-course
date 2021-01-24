package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/grpc-go-course/bi_streaming_grpc/calculator/protocolbuffer"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Calculator Bi Directional Streaming Client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Not able to connect with server: %v", err)
	}

	defer conn.Close()

	c := protocolbuffer.NewCalculatorServiceClient(conn)
	doBiDirectionalStreaming(c)
}

func doBiDirectionalStreaming(c protocolbuffer.CalculatorServiceClient) error {

	log.Println("Starting the client for finding the maximum")

	stream, err := c.FindMaximum(context.Background())

	if err != nil {
		log.Fatalf("Error while calling ComputeAverage RPC: %v", err)
		return err
	}

	waitc := make(chan struct{})

	//we send numbers one by one to server using go-routine
	go func() {
		nunmbers := []int32{5, 20, 9, 4, 3, 21, 6, 90}
		//for loop to send a the numbers
		for _, number := range nunmbers {
			req := &protocolbuffer.FindMaximumRequest{
				Number: number,
			}
			log.Printf("Sending the request: %v\n", number)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}

		stream.CloseSend()
	}()

	//we receive the maximum numbers one by one from the server (go-routine)
	go func() {
		//Loop through each number until receive 'nil' from server
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving GreetEveryone RPC: %v", err)
				break
			}
			log.Printf("Received: %v", res.GetMaximum())
		}
		close(waitc)
	}()

	<-waitc

	return nil
}
