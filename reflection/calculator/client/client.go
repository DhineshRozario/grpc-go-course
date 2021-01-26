package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/grpc-go-course/reflection/calculator/protocolbuffer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	log.Println("Calculator Bi Directional Streaming Client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Not able to connect with server: %v", err)
	}

	defer conn.Close()

	c := protocolbuffer.NewCalculatorServiceClient(conn)

	//With error handling
	doUnaryCalculate(c)

	doClientStreamingComputeAverage(c)

	doServerStreamingPrimeDecomposition(c)

	doBiDirectionalStreamingFindMaximum(c)
}

func doUnaryCalculate(c protocolbuffer.CalculatorServiceClient) {

	log.Println("Starting the client for sum of two numbers in Unary RPC")

	values := []int32{-10, 5, 36}

	waitc := make(chan bool, 1)

	for _, value := range values {
		go callSquareRoot(&waitc, c, value)
	}
	<-waitc

	return
}

// Go subroutine
// 	- called to calculate the square root with error handling

func callSquareRoot(waitc *chan bool, c protocolbuffer.CalculatorServiceClient, value int32) error {
	req := &protocolbuffer.SquareRootRequest{
		Number: value,
	}

	log.Printf("Calculating the square root for of %v\n", req.GetNumber())

	res, err := c.SquareRoot(context.Background(), req)

	if err != nil {
		respError, ok := status.FromError(err)

		if ok {
			//Actual error from GRPC (user Error)
			log.Printf("User error: %v and code: %v", respError.Message(), respError.Code())

			if respError.Code() == codes.InvalidArgument {
				log.Printf("Check the requested number - might be invalid(i.e. negative)")
			}
		} else {
			//Framework Error - Big error
			log.Fatalf("Error while calling 'SquareRoot' Service: %v", err)
		}
		return err
	}

	log.Printf("==> The response from Calculator Server -> 'SquareRoot' of %v is: %v\n", req.GetNumber(), res.GetResult())

	*waitc <- true
	return nil
}

func doServerStreamingPrimeDecomposition(c protocolbuffer.CalculatorServiceClient) {

	log.Println("Starting the getPrimeDecomposition method using the Server Streaming RPC...")

	for i := 1000; i > 1; i -= 10 {
		req := &protocolbuffer.PrimeNumberDeCompositionRequest{
			PrimeNumber: &protocolbuffer.PrimeNumberDeComposition{
				Number: int64(i),
			},
		}
		resStream, err := c.PrimeNumberDeComposition(context.Background(), req)

		if err != nil {
			log.Fatalf("Error while calling PrimeNumberDeComposition RPC: %v", err)
		}

		log.Printf("Response from PrimeNumberDeComposition Server for the number: %v\n", req.GetPrimeNumber().GetNumber())
		for {
			msg, err := resStream.Recv()

			if err == io.EOF {
				log.Println("Reached end of response...")
				break
			} else if err != nil {
				log.Fatalf("Error while receiving the response from Server: %v", err)
			}
			log.Printf("%v ", msg.GetResult())
		}
	}
}

func doClientStreamingComputeAverage(c protocolbuffer.CalculatorServiceClient) {

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

func doBiDirectionalStreamingFindMaximum(c protocolbuffer.CalculatorServiceClient) error {

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
