package pkg

import (
	// "fmt"
	"github.com/tanmaygupta069/api-gateway/config"
	// "github.com/tanmaygupta069/api-gateway/internal/router"
	// "context"
	// "google.golang.org/grpc/connectivity"
	"log"
	// "time"
	// "net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	postpb "github.com/tanmaygupta069/post-service/generated"
	"sync"
)

var once sync.Once

type PostService struct{
	postClient postpb.PostServiceClient
	connection *grpc.ClientConn
}

var postService *PostService

func GetPostServiceClient()(*postpb.PostServiceClient,*grpc.ClientConn){
	if postService.postClient==nil{
		InitializePostServiceClient()
	}
	return &postService.postClient,postService.connection
}

func InitializePostServiceClient(){
	once.Do(func ()  {
		postClient, conn := connectToPostService()
		postService=&PostService{
			postClient:postClient,
			connection:conn,
		}
	})
}

func connectToPostService()(postpb.PostServiceClient, *grpc.ClientConn){
	// Define the Post service host and port
	cfg,_:=config.GetConfig();
	// Dial gRPC server
	conn, err := grpc.NewClient(cfg.PostServiceHost+":"+cfg.PostServicePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Post service: %v", err)
	}

	// Initialize Post service client
	client := postpb.NewPostServiceClient(conn)
	return client, conn
}
