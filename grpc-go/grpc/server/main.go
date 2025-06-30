package main

import (
	pb "github.com/jangirvipin/grpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.GreetServiceServer
}

var port = ":3000"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error starting the server %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &server{})
	log.Printf("Server is running on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error serving gRPC server %v", err)
	}
}
