package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/PranavMasekar/go-GRPC/proto"
)

func callSayHelloBidirectionalStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Bidirectional Streaming started")

	stream, err := client.SayHelloBidirectionaStreaming(context.Background())

	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	channel := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming %v", err)
			}
			log.Println(message)
		}
		close(channel)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		time.Sleep(1 * time.Second)
	}

	stream.CloseSend()
	<-channel
	log.Printf("Bidirectional Streaming finished")

}
