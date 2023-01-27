package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/AndreAppolariFilho/grpc-go-course/greet/greetpb"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("Function LongGreet was invoked")
	res := ""
	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.LongGreetResponse{Result: res})
		}
		if err != nil {
			log.Printf("Error while reading the stream: %v\n", err)
		}
		res += fmt.Sprintf("Hello %s %s\n", msg.Greeting.FirstName, msg.Greeting.LastName)

	}
}
