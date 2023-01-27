package main

import (
	"context"
	"log"
	"time"

	pb "github.com/AndreAppolariFilho/grpc-go-course/greet/greetpb"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")
	reqs := []*pb.LongGreetRequest{
		{
			Greeting: &pb.Greeting{FirstName: "André", LastName: "Appolari"},
		},
		{
			Greeting: &pb.Greeting{FirstName: "Clement", LastName: "Jean"},
		},
		{
			Greeting: &pb.Greeting{FirstName: "Marie", LastName: "Curie"},
		},
	}
	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("An error ocurred while sending request: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n ", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}
	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error ocurred while receiving response: %v\n", err)
	}
	log.Printf("Long Greeting: %s", res.Result)
}
