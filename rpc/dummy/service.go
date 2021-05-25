package dummy

import (
	"os"

	"google.golang.org/grpc"
)

//Create Client
func Connect() (TesterGuideClient, *grpc.ClientConn) {
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "8080"
	}

	conn, err := grpc.Dial(":"+grpcPort, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return NewTesterGuideClient(conn), conn
}
