package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/tanmaygupta069/post-service/generated"
	"github.com/tanmaygupta069/post-service/internal"
	"github.com/tanmaygupta069/post-service/config"
	"github.com/tanmaygupta069/post-service/internal/pkg/dynamo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


func main() {
	cfg,_:=config.GetConfig()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s",cfg.Port))
	if err != nil {
		log.Fatalf("Failed to listen on port %s: ",cfg.Port, err.Error())
	}

	grpcServer := grpc.NewServer()

	
    dynamo.InitializeDynamoClient() 
	
	dynamo:=dynamo.NewDynamoClient()
    postRepo:=internal.NewDynamoDBPostRepository(*dynamo,"posts")

	if postRepo==nil{
		fmt.Println("post repo nil")
	}


	postService := internal.NewPostServiceServer(postRepo)

	if postService==nil{
		fmt.Println("postservice nil")
	}


	postController := internal.NewControllerServer(postService)

	if postController==nil{
		fmt.Println("postcontroller nil")
	}

	pb.RegisterPostServiceServer(grpcServer,postController)

	reflection.Register(grpcServer)

	log.Printf("Server is listening on port %s",cfg.Port)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server on port %s: %v",cfg.Port, err.Error())
	}
}
