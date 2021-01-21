package main

import (
	"io"
	"log"
	"net"
	"strconv"

	"github.com/grpc-go-course/server_streaming_grpc/average/protocolbuffer"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) ComputeAverage(stream protocolbuffer.ComputeAverageService_ComputeAverageServer) error {

	log.Println("ComputeAverage function was invoked with Streaming Client request:")
	average := float32(0)
	sum := float32(0)
	numbers := ""
	count := 0
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			average = float32(sum / float32(count));
			log.Printf("For the given numbers (%v), the sum %v and the Average is: %v", numbers, sum, average)
			// we have finished reading the client stream

			return stream.SendAndClose(&protocolbuffer.ComputeAverageResponse{
				// Average: strconv.FormatFloat(average, 'f', 0, 64),
				Average: average,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		number := req.GetComputeAverage().GetNumber()
		sum += float32(number)
		numbers += strconv.Itoa(int(number)) + ", "
		count++
	}
}

func main() {
	log.Println("Compute Average Server!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	s := grpc.NewServer()

	protocolbuffer.RegisterComputeAverageServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
