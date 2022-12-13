package main

import (
	"io"
	"log"

	pb "github.com/PranavMasekar/go-GRPC/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string

	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}
		if err != nil {
			return err
		}
		messages = append(messages, "Hello "+request.Name)
		log.Printf("Got request with name : %v", request.Name)
	}
}
