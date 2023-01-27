package main

import (
	"fmt"
	"log"

	pb "github.com/AndreAppolariFilho/grpc-go-course/greet/greetpb"
)

func (s *Server) GreetManyTimes(in *pb.GreetManyTimesRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with: %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d\n", in.Greeting.FirstName, i)
		stream.Send(&pb.GreetManytimesResponse{Result: res})
	}
	return nil
}
