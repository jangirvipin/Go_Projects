package main

import (
	"context"
	pb "github.com/jangirvipin/grpc/proto"
	"log"
	"time"
)

func CallUnaryService(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &pb.GreetRequest{
		Name: "Vipin Jangir",
	}

	res, err := client.Greet(ctx, req)
	if err != nil {
		log.Fatal("Error calling Greet service: ", err)
		return
	}

	log.Printf("Response from Greet service: %s", res.Result())
	return
}
