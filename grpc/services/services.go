package services

import (
	context "context"

	"github.com/thinkgos/distributed/grpc/pb"
)

// user's defined
type Arith struct {
	pb.UnimplementedArithServer
}

func (a *Arith) Mul(ctx context.Context, req *pb.ArithRequest) (*pb.ArithResponse, error) {
	return &pb.ArithResponse{
		Result: req.A * req.B,
	}, nil
}
