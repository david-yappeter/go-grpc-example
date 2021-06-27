package main

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "myapp/rpc/math"

	"myapp/grpc/resolver"

	"google.golang.org/grpc"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterMathGuideServer(grpcServer, &resolver.Server{})

	log.Printf("Listen on port : %s\n", port)
	grpcServer.Serve(lis)
}
