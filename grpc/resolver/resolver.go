package resolver

import (
	"context"
	"myapp/rpc/math"
)

type Server struct {
}

func (s *Server) Plus(ctx context.Context, request *math.PlusParam) (*math.PlusResponse, error) {
	return &math.PlusResponse{
		Result: request.A + request.B,
	}, nil
}
