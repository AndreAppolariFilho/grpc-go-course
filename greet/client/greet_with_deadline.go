package main

import (
	"context"
	"log"
	"time"

	pb "github.com/AndreAppolariFilho/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeOut time.Duration) {
	log.Println("doGreetWithDeadline was invoke")
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()
	req := &pb.GreetWithDeadlineRequest{
		Greeting: &pb.Greeting{
			FirstName: "André",
			LastName:  "Appolari",
		},
	}
	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded!")
				return
			} else {
				log.Fatalf("Unexpected gRPC error %v\n", err)
			}
		} else {
			log.Fatalf("Non grpc error %v\n", err)
		}
	}
	log.Printf("GreetWithDeadline: %s\n", res.Result)
}
