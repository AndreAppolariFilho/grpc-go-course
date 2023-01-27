package main

import (
	"io"
	"log"

	pb "github.com/AndreAppolariFilho/grpc-go-course/calculator/calculatorpb"
)

func (s *server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg function was invoked")
	var total float64 = 0
	var quantity float64 = 0
	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Response: total / quantity,
			})
		}
		if err != nil {
			log.Fatalf("An error ocurred while processing the stream: %v\n", err)
		}
		total += float64(msg.Number)
		quantity += 1
	}
}
