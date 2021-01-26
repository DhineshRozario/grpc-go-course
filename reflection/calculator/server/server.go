package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/grpc-go-course/reflection/calculator/protocolbuffer"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) ComputeAverage(stream protocolbuffer.CalculatorService_ComputeAverageServer) error {

	log.Println("ComputeAverage function was invoked with Streaming Client request:")
	average := float32(0)
	sum := float32(0)
	numbers := ""
	count := 0
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			average = float32(sum / float32(count))
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

func (*server) PrimeNumberDeComposition(req *protocolbuffer.PrimeNumberDeCompositionRequest, stream protocolbuffer.CalculatorService_PrimeNumberDeCompositionServer) error {

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

func (*server) SquareRoot(ctx context.Context, req *protocolbuffer.SquareRootRequest) (*protocolbuffer.SquareRootResponse, error) {
	log.Printf("Server function SquareRoot() RPC was invoked\nwith request: ==> %v\n", req)

	number := req.GetNumber()

	if number < 0 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Received a negative number: %v", number))
	}

	return &protocolbuffer.SquareRootResponse{
		Result: math.Sqrt(float64(number)),
	}, nil
}

func (*server) SumOfNumbers(ctx context.Context, req *protocolbuffer.CalculatorRequest) (*protocolbuffer.CalculatorResponse, error) {
	log.Printf("Server function SumOfNumbers() was invoked\nwith request: ==> %v\n", req)

	result := req.GetCalculator().GetFirstNumber() + req.GetCalculator().GetSecondNumber()

	res := protocolbuffer.CalculatorResponse{
		Result: result,
	}

	return &res, nil
}

func main() {

	log.Println("Calculator Server - Unary, Client, Server and Bi-Directional Streaming examples")

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
