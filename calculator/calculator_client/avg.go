package main

import (
	"context"
	"log"

	pb "github.com/AndreAppolariFilho/grpc-go-course/calculator/calculatorpb"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg was invoked")
	reqs := []*pb.AvgRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}
	stream, err := c.Avg(context.Background())

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error ocurred while receiving response: %v\n", err)
	}
	log.Printf("Avg result: %f", res.Response)
}
