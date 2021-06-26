package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"myapp/rpc/dummy"

	"myapp/grpc/resolver"

	"google.golang.org/grpc"
)

//GRPC Server
func main() {
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "8080"
	}

	//Net Listen
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		panic(err)
	}

	//Grpc Server
	grpcServer := grpc.NewServer()
	dummy.RegisterTesterGuideServer(grpcServer, &resolver.Server{})

	//Listen and Serve tcp
	log.Println("Listen and Serve at port :" + port)
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println(err)
	}
}
