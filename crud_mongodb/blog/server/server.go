package main

import (
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/grpc-go-course/crud_mongodb/blog/protocolbuffer"
	"google.golang.org/grpc"
)

type server struct {
}

func main() {
	log.Println("Blog Server Started!")

	//Adding the logger level for the details output
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)

	lis, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	protocolbuffer.RegisterBlogServiceServer(s, &server{})

	//shutdown hook - for greacefully shutdown the server
	go func() {
		log.Println("Starting the Server!")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v\n", err)
		}
	}()

	//Wait for Control+C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	//Blocking the server for signal (Control+C)
	<-ch

	log.Println("Stopping the Server!")
	s.Stop()
	log.Println("Closing the Listener!")
	lis.Close()
	log.Println("End of the Program!")
}
