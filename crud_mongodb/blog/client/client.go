package main

import (
	"context"
	"log"

	"github.com/grpc-go-course/crud_mongodb/blog/protocolbuffer"
	"google.golang.org/grpc"
)

func main() {
	//Adding the logger level for the details output
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Starting the Blog Client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Not able to connect with server: %v", err)
	}
	defer conn.Close()

	c := protocolbuffer.NewBlogServiceClient(conn)

	//Blog 1
	log.Println("Creating the blog 1")
	blog := &protocolbuffer.Blog{
		AuthorId: "Dewiz Rozario",
		Title:    "My Fathers Blog using Go-Proto with MongoDB",
		Content:  "Just trying out something to add with the blog",
	}

	blogResponse, err := c.CreateBlog(context.Background(), &protocolbuffer.CreateBlogRequest{
		Blog: blog,
	})

	if err != nil {
		log.Fatalf("Not able to create the blog in server: %v", err)
	}

	log.Printf("The blog has been created: %v", blogResponse)

	// Read the Created Blog
	log.Println("Reading the blog")

	_, err2 := c.ReadBlog(context.Background(), &protocolbuffer.ReadBlogRequest{
		BlogId: "some_random_id",
	})
	if err2 != nil {
		log.Printf("Error happened while reading: %v", err2)
	}

	readBlogResponse, err3 := c.ReadBlog(context.Background(), &protocolbuffer.ReadBlogRequest{
		BlogId: blogResponse.GetBlog().GetId(),
	})
	if err3 != nil {
		log.Printf("Error happened while reading: %v", err3)
	}
	log.Printf("Id found: %v", readBlogResponse.GetBlog())

	// //Blog 2
	// log.Println("Creating the blog 2")
	// blog = &protocolbuffer.Blog{
	// 	AuthorId: "Dewin Rozario",
	// 	Title:    "My Fathers Blog using Go-Proto with MongoDB",
	// 	Content:  "Just trying out something to add with the blog",
	// }

	// blogResponse, err = c.CreateBlog(context.Background(), &protocolbuffer.CreateBlogRequest{
	// 	Blog: blog,
	// })

	// if err != nil {
	// 	log.Fatalf("Not able to create the blog in server: %v", err)
	// }

	// log.Printf("The blog has been created: %v", blogResponse)
}
