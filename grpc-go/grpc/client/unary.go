package main

import (
	"context"
	"fmt"
	pb "github.com/jangirvipin/grpc/proto"
	"io"
	"log"
	"sync"
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

	log.Printf("Response from Greet service: %s", res.Result)
	return
}

func CallGreetManyTimeService(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &pb.MessageList{
		Messages: []string{"Alice", "John Doe", "Jane Smith"},
	}
	stream, err := client.GreetManyTime(ctx, req)
	if err != nil {
		log.Fatalf("Error calling GreetManyTime service: %v", err)
		return
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			if err.Error() == "EOF" {
				log.Println("Stream closed by server")
				break
			}
			log.Fatalf("Error receiving response from GreetManyTime service: %v", err)
		}
		log.Printf("Response from GreetManyTime service: %s", res.Result)
	}
}

func CallLongGreet(client pb.GreetServiceClient) {
	stream, err := client.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error calling LongGreet service: %v", err)
		return
	}
	names := []string{"Alice", "John Doe", "Jane Smith"}

	for _, name := range names {
		fmt.Println("Sending name:", name)
		err := stream.Send(&pb.GreetRequest{
			Name: name,
		})
		if err != nil {
			log.Fatalf("Error sending name to LongGreet service: %v", err)
			return
		}
		time.Sleep(500 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response from LongGreet service: %v", err)
		return
	}
	fmt.Printf("Response from LongGreet service: %v\n", res.Names)
}

func CallGreetEveryone(client pb.GreetServiceClient) {
	var wg sync.WaitGroup

	fmt.Println("Calling GreetEveryone")
	stream, err := client.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error calling GreetEveryone service: %v", err)
		return
	}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error receiving response from GreetEveryone service: %v", err)
			}
			fmt.Println("Received GreetEveryone response:", res.Result)
		}
	}(&wg)

	names := []string{"Alice", "John Doe", "Jane Smith"}

	for _, name := range names {
		fmt.Println("Sending name:", name)
		err := stream.Send(&pb.GreetRequest{
			Name: name,
		})
		if err != nil {
			log.Fatalf("Error sending name to GreetEveryone service: %v", err)
			return
		}
		time.Sleep(500 * time.Millisecond)
	}
	stream.CloseSend()
	wg.Wait()
}
