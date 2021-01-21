package main

import (
	"log"
	"net"
	"strconv"

	"github.com/grpc-go-course/server_streaming_grpc/primenumber/protocolbuffer"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) PrimeNumberDeComposition(req *protocolbuffer.PrimeNumberDeCompositionRequest, stream protocolbuffer.PrimeNumberDeCompositionService_PrimeNumberDeCompositionServer) error {

	log.Printf("PrimeNumberDeComposition function was invoked with request:\n==> %v\n", req)

	givenNumber := req.GetPrimeNumber().GetNumber()
	log.Printf("PrimeNumberDeComposition: given number is: %v", givenNumber)

	dividor := int64(2)

	for givenNumber > 1 {

		if givenNumber%dividor == 0 {
			result := "dividor: " + strconv.FormatInt(dividor, 10) + " and number is: " + strconv.FormatInt(givenNumber, 10)
			log.Printf("result: %v", result)
			res := &protocolbuffer.PrimeNumberDeCompositionManyTimesResponse{
				Result: result,
			}
			stream.Send(res)
			givenNumber = givenNumber / dividor
		} else {
			dividor = dividor + 1
		}
	}

	return nil
}

func main() {
	log.Println("Prime Number Decomposition Server!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	s := grpc.NewServer()

	protocolbuffer.RegisterPrimeNumberDeCompositionServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
