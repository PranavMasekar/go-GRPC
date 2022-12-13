package main

import (
	"context"

	pb "github.com/PranavMasekar/go-GRPC/proto"
)

func (s *helloServer) SayHello(ctx context.Context, request *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
