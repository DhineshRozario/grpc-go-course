package main

import (
	"io"
	"log"
	"net"

	"github.com/grpc-go-course/bi_streaming_grpc/calculator/protocolbuffer"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) FindMaximum(stream protocolbuffer.CalculatorService_FindMaximumServer) error {
	log.Println("FindMaximum() was invoked to calculate the maximum from given numbers")

	maximum := int32(0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
			return err
		}

		number := req.GetNumber()
		log.Printf("Received the request :%v\n", number)

		if number > maximum {
			maximum = number

			res := &protocolbuffer.FindMaximumResponse{
				Maximum: maximum,
			}

			sendErr := stream.Send(res)
			if sendErr == io.EOF {
				return nil
			}
			if err != nil {
				log.Fatalf("Error while reading client stream: %v\n", sendErr)
				return sendErr
			}
		}

	}
}

func main() {

	log.Println("Bi-Directional - Calculator Server")

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
