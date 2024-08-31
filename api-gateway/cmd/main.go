package main

import (
	"fmt"
	"github.com/tanmaygupta069/api-gateway/config"
	"github.com/tanmaygupta069/api-gateway/internal/router"
	"context"
	"google.golang.org/grpc/connectivity"
	"log"
	"time"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	// postpb "github.com/tanmaygupta069/post-service/generated"

)

func main(){
	cfg,err:= config.GetConfig()
	if(err!=nil){
		fmt.Println("an error occured ",err)
	}
	router := router.GetRouter()
	// Check if the Post Service is up
	// if !waitForService(cfg.PostServiceHost, cfg.PostServicePort, 10*time.Second) {
	// 	log.Fatalf("Post Service at %s:%s is not reachable", cfg.PostServiceHost, cfg.PostServicePort)
	// }
	router.Run(":"+cfg.ServerPort)
}

func waitForService(host, port string, timeout time.Duration) bool {
	address := net.JoinHostPort(host, port)
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Failed to connect to %s: %v", address, err)
		return false
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	for {
		state := conn.GetState()
		if state == connectivity.Ready {
			log.Printf("Service %s is ready", address)
			return true
		}

		if !conn.WaitForStateChange(ctx, state) {
			log.Printf("Timeout waiting for service %s to be ready", address)
			return false
		}
	}
}

