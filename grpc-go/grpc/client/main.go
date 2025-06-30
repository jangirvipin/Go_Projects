package main

import (
	pb "github.com/jangirvipin/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var port = ":3000"

func main() {
	lis, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer lis.Close()
	client := pb.NewGreetServiceClient(lis)

	CallUnaryService(client)
}
