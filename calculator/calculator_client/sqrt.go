package main

import (
	"context"
	"log"

	pb "github.com/AndreAppolariFilho/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Printf("doSqrt was invoked")
	req := &pb.SqrtRequest{Number: n}
	res, err := c.Sqrt(context.Background(), req)
	if err != nil {

		e, ok := status.FromError(err)
		if ok {
			log.Printf("Error message from server: %s\n", e.Message())
			log.Printf("Error code from server: %s\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Println("We probably sent a negative number")
				return
			}
		} else {
			log.Fatalf("Non grpc error while calling Sqrt RPC: %v", err)
		}
	}
	log.Printf("Sqrt of %d is equal to %f", n, res.Result)
}
