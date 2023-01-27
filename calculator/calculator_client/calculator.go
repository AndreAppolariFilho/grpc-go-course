package main

import (
	"context"
	"log"

	pb "github.com/AndreAppolariFilho/grpc-go-course/calculator/calculatorpb"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("Starting to do a Sum Unary RPC...")
	req := &pb.SumRequest{
		FirstNumber:  3,
		SecondNumber: 10,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Sum RPC: %v", err)
	}
	log.Printf("Response from Sum: %v", res.SumResult)
}
