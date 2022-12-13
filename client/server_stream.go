package main

import (
	"context"
	"io"
	"log"

	pb "github.com/PranavMasekar/go-GRPC/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Streaming has started")

	stream, err := client.SayHelloServerStreaming(context.Background(), names)

	if err != nil {
		log.Fatalf("Could not send Stream %v", err)
	}

	for {
		message, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while straming %v", err)
		}

		log.Println(message)
	}

	log.Printf("Streaming finished")
}
