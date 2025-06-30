package main

import (
	"context"
	"fmt"
	pb "github.com/jangirvipin/grpc/proto"
	"io"
	"time"
)

func (s *server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	fmt.Println("Received Greet request")
	response := "Hello " + req.Name + "!"
	return &pb.GreetResponse{
		Result: response,
	}, nil
}

func (s *server) GreetManyTime(req *pb.MessageList, stream pb.GreetService_GreetManyTimeServer) error {
	fmt.Println("Received GreetManyTime request")
	for _, name := range req.Messages {
		response := &pb.GreetResponse{
			Result: "Hello " + name,
		}
		if err := stream.Send(response); err != nil {
			return fmt.Errorf("error sending response: %v", err)
		}
		time.Sleep(500 * time.Millisecond)
	}
	return nil
}

func (s *server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	var names []string
	fmt.Println("Received LongGreet request")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.NameList{
				Names: names,
			})
		}
		if err != nil {
			return fmt.Errorf("error receiving request: %v", err)
		}
		names = append(names, req.Name)
		fmt.Printf("Received name: %s\n", req.Name)
	}
}

func (s *server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	fmt.Println("Received GreetEveryone request")
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("Received GreetEveryone response")
			return nil
		}

		if err != nil {
			return fmt.Errorf("error receiving request: %v", err)
		}
		name := res.Name
		fmt.Println("Received GreetEveryone response: " + name)

		response := &pb.GreetResponse{
			Result: "Hello " + name,
		}

		err = stream.Send(response)
		if err != nil {
			return fmt.Errorf("error sending response: %v", err)
		}
	}
}
