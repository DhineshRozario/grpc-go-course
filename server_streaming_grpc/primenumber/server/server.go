package main

import (
	"log"
	"net"

	"github.com/grpc-go-course/server_streaming_grpc/primenumber/protocolbuffer"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) PrimeNumberDeComposition(req *protocolbuffer.PrimeNumberRequest, stream protocolbuffer.PrimeNumberService_PrimeNumberDeCompositionServer) error {

	log.Printf("PrimeNumberDeComposition function was invoked with request:\n==> %v\n", req)

	givenNumber := int(req.GetPrimeNumber().GetNumber())

	log.Printf("PrimeNumberDeComposition: given number is: %v", givenNumber)

	dividor := 2
	for givenNumber > 1 {
		modValue, err := findDividor(stream, dividor, givenNumber)

		if err != nil {
			log.Fatalf("failed to find the dividor: %v\n", err)
			break
		}

		if modValue != 0 {
			dividor = dividor + 1
		}

		givenNumber = givenNumber / dividor
	}

	return nil
}

func findDividor(stream protocolbuffer.PrimeNumberService_PrimeNumberDeCompositionServer, dividor int, givenNumber int) (int, error) {

	modValue := givenNumber % dividor

	if modValue == 0 {
		res := &protocolbuffer.PrimeNumberManyTimesResponse{
			Result: int32(dividor),
		}
		stream.Send(res)
	}

	return modValue, nil
}

func main() {
	log.Println("Prime Number Decomposition Server!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	s := grpc.NewServer()

	protocolbuffer.RegisterPrimeNumberServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
