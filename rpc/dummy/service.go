package dummy

import (
	"fmt"
	"os"

	"google.golang.org/grpc"
)

//Create Client
func Connect() (TesterGuideClient, *grpc.ClientConn) {
	grpcPort := os.Getenv("GRPC_ADDR")
	if grpcPort == "" {
		grpcPort = ":8080"
	}

	fmt.Println("port: ", grpcPort)

	conn, err := grpc.Dial(grpcPort, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return NewTesterGuideClient(conn), conn
}
