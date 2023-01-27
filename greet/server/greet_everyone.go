package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/AndreAppolariFilho/grpc-go-course/greet/greetpb"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone was invoked")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading from stream: %v\n", err)
		}
		res := fmt.Sprintf("Hello %s %s!", req.Greeting.FirstName, req.Greeting.LastName)
		err = stream.Send(&pb.GreetEveryoneResponse{Result: res})
		if err != nil {
			log.Fatalf("Error while sending from stream: %v\n", err)
		}
	}
	return nil
}
