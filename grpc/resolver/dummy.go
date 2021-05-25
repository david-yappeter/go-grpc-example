package resolver

import (
	"context"
	"fmt"
	pb "myapp/rpc/dummy"
)

type Server struct {
}

//Grpc Function
func (s *Server) DummyRpc(ctx context.Context, request *pb.DummyRpcParam) (*pb.DummyRpcResponse, error) {

	// Get the parameter
	temp := request.A * request.B
	fmt.Println(temp)

	return &pb.DummyRpcResponse{
		Result: temp,
	}, nil
}
