package main

import (
	"io"
	"log"
	"net"

	"github.com/grpc-go-course/bi_streaming_grpc/greet/protocolbuffer"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) GreetEveryone(stream protocolbuffer.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone function was invoked with Bi-Directional Streaming")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
			return err
		}
		
		log.Printf("Received the request :%v\n", req)
		firstName := req.GetGreeting().GetFirstName()
		result := "Hello " + firstName + "!"

		res := &protocolbuffer.GreetEveryoneResponse{
			Result: result,
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

func main() {
	log.Println("Hello World - am a Server!")

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
