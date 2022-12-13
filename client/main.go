package main

import (
	"log"

	pb "github.com/PranavMasekar/go-GRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8000"
)

func main() {
	connection, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("Did not connect ", err)
	}

	defer connection.Close()

	client := pb.NewGreetServiceClient(connection)

	// names := &pb.namesList{
	// 	Names: []string{"Pranav", "Rohit", "Mortal"},
	// }

	callSayHello(client)
}
