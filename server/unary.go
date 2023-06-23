package main

import (
	"context"

	pb "github.com/akhil/grpc-demo/proto"
)

// Обычный вызов
func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
