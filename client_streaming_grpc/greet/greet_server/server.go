package main

import (
	"context"
	"io"
	"log"
	"net"
	"strconv"

	"github.com/grpc-go-course/server_streaming_grpc/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {

	log.Printf("Greet function was invoked with request :%v\n", req)

	firstName := req.GetGreeting().GetFirstName()

	result := "Hello " + firstName

	res := greetpb.GreetResponse{
		Result: result,
	}

	return &res, nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {

	log.Println("LongGreet function was invoked with Streaming Client request:")
	result := ""
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			log.Printf("Received the content: %v", result)
			// we have finished reading the client stream
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: "Hello " + result + "!!!",
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		firstName := req.GetGreeting().GetFirstName()

		result += firstName + ", "
	}
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {

	log.Printf("GreetManyTimes function was invoked with request:\n==> %v\n", req)

	firstName := req.GetGreeting().GetFirstName()

	for i := 0; i < 10; i++ {

		result := "Hello " + firstName + ", calling " + strconv.Itoa(i)

		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}

		stream.Send(res)
		// time.Sleep(1000 * time.Millisecond)
	}

	return nil
}

func main() {
	log.Println("Hello World - am a Server!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
