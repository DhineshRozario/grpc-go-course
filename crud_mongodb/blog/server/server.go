package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/grpc-go-course/crud_mongodb/blog/protocolbuffer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var collection *mongo.Collection

type server struct {
}

type blogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

func (*server) ReadBlog(ctx context.Context, req *protocolbuffer.ReadBlogRequest) (*protocolbuffer.ReadBlogResponse, error) {

	blogID := req.GetBlogId()

	oid, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		log.Fatalf("Cannot parse the given id from Request: %v and the error: %v", blogID, err)
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Cannot parse the given id from Request: %v", blogID))
	}

	//Create an empty data
	data := &blogItem{}

	filter := bson.M{"_id": oid}
	result := collection.FindOne(ctx, filter)

	if decodeError := result.Decode(&data); decodeError != nil {
		log.Fatalf("Cannot find the blog with the specified ID: %v, and error: %v", oid, decodeError)
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Cannot find the blog with the specified ID: %v", decodeError))
	}

	responseBlog := &protocolbuffer.ReadBlogResponse{
		Blog: dataToBlogProtocolBuffer(data),
	}

	return responseBlog, nil

}

func (*server) CreateBlog(ctx context.Context, req *protocolbuffer.CreateBlogRequest) (*protocolbuffer.CreateBlogResponse, error) {

	blog := req.GetBlog()

	log.Printf("CreateBlog: Received a request: %v", blog)
	//Creating the request with mongoDB updated format
	data := blogItem{
		AuthorID: blog.GetAuthorId(),
		Title:    blog.GetTitle(),
		Content:  blog.GetContent(),
	}
	log.Printf("CreateBlog: Converted data into a blogItem for inserting into MonogDB: %v", data)
	//Inserting into DB with the given data
	res, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Unable to insert the data into DB: %v", err))
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot convert to OID from Response"))
	}

	//Returning the response with updated id
	responseBlog := protocolbuffer.CreateBlogResponse{
		Blog: &protocolbuffer.Blog{
			Id:       oid.Hex(),
			AuthorId: data.AuthorID,
			Title:    data.Title,
			Content:  data.Content,
		},
	}
	log.Printf("CreateBlog: Successfully inserted into MonogDB, sending the response: %v", responseBlog.GetBlog())
	return &responseBlog, nil
}

func dataToBlogProtocolBuffer(data *blogItem) *protocolbuffer.Blog {
	return &protocolbuffer.Blog{
		Id:       data.ID.Hex(),
		AuthorId: data.AuthorID,
		Title:    data.Title,
		Content:  data.Content,
	}
}

func (*server) UpdateBlog(ctx context.Context, req *protocolbuffer.UpdateBlogRequest) (*protocolbuffer.UpdateBlogResponse, error) {
	blog := req.GetBlog()

	log.Printf("UpdateBlog: Received a request: %v", blog)

	oid, err := primitive.ObjectIDFromHex(blog.GetId())
	if err != nil {
		log.Fatalf("Cannot parse the given id from Request: %v and the error: %v", blog.GetId(), err)
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Cannot parse the given id from Request: %v", blog.GetId()))
	}

	//Create an empty data
	data := &blogItem{}

	filter := bson.M{"_id": oid}
	result := collection.FindOne(ctx, filter)

	if decodeError := result.Decode(&data); decodeError != nil {
		log.Fatalf("Cannot find the blog with the specified ID: %v, and error: %v", oid, decodeError)
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Cannot find the blog with the specified ID: %v", decodeError))
	}

	//Updating the 'data' object with the existing values
	data.AuthorID = blog.GetAuthorId()
	data.Content = blog.GetContent()
	data.Title = blog.GetTitle()

	//Updating the MongoDB with the modified 'data' object
	_, updateErr := collection.ReplaceOne(context.Background(), filter, data)
	if updateErr != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Unable to update the blog into DB: %v", updateErr))
	}

	return &protocolbuffer.UpdateBlogResponse{
		Blog: dataToBlogProtocolBuffer(data),
	}, nil
}

func (s *server) DeleteBlog(ctx context.Context, req *protocolbuffer.DeleteBlogRequest) (*protocolbuffer.DeleteBlogResponse, error) {
	blogID := req.BlogId

	log.Printf("UpdateBlog: Received a request: %v", blogID)

	oid, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		log.Fatalf("Cannot parse the given id from Request: %v and the error: %v", blogID, err)
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Cannot parse the given id from Request: %v", blogID))
	}

	filter := bson.M{"_id": oid}

	//Deleting the blog
	deleteResult, deleteErr := collection.DeleteOne(context.Background(), filter)
	if deleteErr != nil {
		log.Fatalf("Unable to delete the blog:%v in MongoDB: %v", blogID, err)
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Unable to delete the blog:%v in MongoDB: %v", blogID, err))
	}

	if deleteResult.DeletedCount == 0 {
		log.Fatalf("Unable to find the blog:%v in MongoDB: %v", blogID, err)
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Unable to find the blog:%v in MongoDB: %v", blogID, err))
	}

	log.Printf("Deleted the Blog: %v", blogID)

	return &protocolbuffer.DeleteBlogResponse{
		BlogId: blogID,
	}, nil

}

func main() {
	//Adding the logger level for the details output
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//Connect to MongoDB
	log.Println("Creating Client for MongoDB!")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("failed to create a Client for MongoDB: %v\n", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	log.Println("Connecting to MongoDB!")
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("failed to connect the Client to MongoDB: %v\n", err)
		return
	}
	collection = client.Database("mydb").Collection("blog")

	log.Println("Blog Server Started!")

	//Server Options
	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	protocolbuffer.RegisterBlogServiceServer(s, &server{})
	//shutdown hook - for greacefully shutdown the server
	go func() {
		log.Println("Starting the Listning the request!")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v\n", err)
			return
		}
		return
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

	//Closing the mongoDB Connection
	log.Println("Closing the mongoDB Connection!")
	client.Disconnect(context.TODO())

	log.Println("End of the Program!")
}
