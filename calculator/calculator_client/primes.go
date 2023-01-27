package main

import (
	"context"
	"io"
	"log"

	pb "github.com/AndreAppolariFilho/grpc-go-course/calculator/calculatorpb"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("doPrimes was invoked")
	req := &pb.PrimesRequest{Number: 120}
	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("An error ocurred in doPrimes: %v\n", err)
	}
	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading the numbers in doPrimes: %v\n", err)
		}
		log.Printf("Primes: %d", msg.PrimerNumber)
	}

}
