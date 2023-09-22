package main

import (
	"context"
	pb "shyam_kadiyam/grpc/proto"
)

// Implementing SayHello
func (hello *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "hello",
	}, nil
}
