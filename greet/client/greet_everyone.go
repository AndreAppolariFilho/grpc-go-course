package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/AndreAppolariFilho/grpc-go-course/greet/greetpb"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Printf("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Printf("Error while parsing the stream: %v\n", err)
	}
	reqs := []*pb.GreetEveryoneRequest{
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
	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending Request: %v", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()
	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Panicf("Error while reading from stream: %v\n")
				break
			}
			log.Printf("Received: %v\n", res.Result)

		}
		close(waitc)
	}()
	<-waitc
}
