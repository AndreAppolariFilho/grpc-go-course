package main

import (
	"io"
	"log"
	"math"

	pb "github.com/AndreAppolariFilho/grpc-go-course/calculator/calculatorpb"
)

func (s *server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max was invoked")
	var current_max int32 = math.MinInt32
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Panicf("An error ocurred during the processing of stream: %v\n", err)
			return err
		}
		if msg.Number >= current_max {
			current_max = msg.Number
			err = stream.Send(&pb.MaxResponse{Response: current_max})
			if err != nil {
				log.Panicf("An error ocurred during the sending of stream: %v\n", err)
				break
			}
		}
	}
	return nil
}
