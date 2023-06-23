package main

import (
	"context"

	pb "go_grpc/proto"
)

// Обычный вызов
func (s *helloServer) SayHelloMethod(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
