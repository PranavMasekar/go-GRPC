package main

import (
	"context"
	"log"
	"time"

	pb "github.com/PranavMasekar/go-GRPC/proto"
)

func callSayHelloClientStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client Streaming started")
	stream, err := client.SayHelloClientStreaming(context.Background())

	if err != nil {
		log.Fatal("Could not send names")
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatal("Error while sending stream")
		}

		log.Printf("Sent the request with name %s", name)
		time.Sleep(2 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	log.Printf("Client Streaming finished")
	if err != nil {
		log.Fatalf("Error while receiving %v", err)
	}

	log.Printf("%v", res.Messages)
}
