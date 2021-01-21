package main

import (
	"context"
	"log"
	"time"

	"github.com/grpc-go-course/server_streaming_grpc/average/protocolbuffer"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Compute Average Streaming Client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Not able to connect with server: %v", err)
	}

	defer conn.Close()

	c := protocolbuffer.NewComputeAverageServiceClient(conn)

	getAverage(c)
}

func getAverage(c protocolbuffer.ComputeAverageServiceClient) {

	log.Println("Starting the getAverage method using the Client Streaming RPC...")

	stream, err := c.ComputeAverage(context.Background())

	if err != nil {
		log.Fatalf("Error while calling ComputeAverage RPC: %v", err)
	}

	// values := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	values := []int32{45, 46, 75, 60}

	for _, value := range values {
		req := &protocolbuffer.ComputeAverageStreamRequest{
			ComputeAverage: &protocolbuffer.ComputeAverage{
				Number: value,
			},
		}

		// log.Printf("Sending the request with number: %v", req.GetComputeAverage().GetNumber())
		stream.Send(req)
		time.Sleep(1 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while calling closing the stream: %v", err)
	}

	log.Printf("Compute Average for the values %v ->: %v", values, res.GetAverage())
}
