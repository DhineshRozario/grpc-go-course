package main

import (
	"context"
	"log"

	"github.com/grpc-go-course/unary_grpc/squareroot/protocolbuffer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	log.Println("Calculator client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Not able to connect with server: %v", err)
	}

	defer conn.Close()

	c := protocolbuffer.NewCalculatorServiceClient(conn)
	doCalculate(c)
}

func doCalculate(c protocolbuffer.CalculatorServiceClient) {

	log.Println("Starting the client for sum of two numbers in Unary RPC")

	values := []int32{-10, 5, 36}

	waitc := make(chan string)

	for _, value := range values {
		go callSquareRoot(waitc, c, value)
	}

	<-waitc

	return
}

func callSquareRoot(waitc chan<- string, c protocolbuffer.CalculatorServiceClient, value int32) error {
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

	close(waitc)
	return nil
}
