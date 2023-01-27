package main

import (
	"context"
	"log"
	"time"

	pb "github.com/AndreAppolariFilho/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetWithDeadlineRequest) (*pb.GreetWithDeadlineResponse, error) {
	log.Printf("GreetWithDeadline function was invoked with %v\n", in)
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client canceled the request")
			return nil, status.Error(codes.Canceled, "The client canceled the request")
		}
		time.Sleep(1 * time.Second)
	}
	return &pb.GreetWithDeadlineResponse{
		Result: "Hello " + in.Greeting.FirstName,
	}, nil
}
