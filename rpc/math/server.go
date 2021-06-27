package math

import (
	"os"
	"time"

	"google.golang.org/grpc"
)

const defaultGrpcPort = ":8080"

//ConnectionOption ConnectionOption
type ConnectionOption struct {
	Host    *string
	TimeOut *time.Duration
}

func Connect(connectionOption ...ConnectionOption) (MathGuideClient, *grpc.ClientConn) {
	grpcAddr := os.Getenv("GRPC_ADDRESS")
	if grpcAddr == "" {
		if len(connectionOption) > 0 && connectionOption[0].Host != nil {
			grpcAddr = *connectionOption[0].Host
		}
	}
	if grpcAddr == "" {
		grpcAddr = defaultGrpcPort
	}

	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return NewMathGuideClient(conn), conn
}
