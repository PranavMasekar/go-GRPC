package main

import (
	"log"
	"net"

	pb "github.com/PranavMasekar/go-GRPC/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8000"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to start the server", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("Server started at %v", listen.Addr())

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failded to start the server %v", err)
	}
}
