package main

import (
	"io"
	"log"

	pb "github.com/PranavMasekar/go-GRPC/proto"
)

func (s *helloServer) SayHelloBidirectionaStreaming(stream pb.GreetService_SayHelloBidirectionaStreamingServer) error {
	for {
		request, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		log.Printf("Got request with name : %v", request.Name)

		res := &pb.HelloResponse{
			Message: "Hello " + request.Name,
		}

		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
