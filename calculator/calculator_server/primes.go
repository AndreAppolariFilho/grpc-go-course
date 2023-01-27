package main

import (
	"log"

	pb "github.com/AndreAppolariFilho/grpc-go-course/calculator/calculatorpb"
)

func (s *server) Primes(in *pb.PrimesRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Println("Function Primes was invoked")
	var k int32 = 2
	for N := in.Number; N > 1; {
		if N%k == 0 {
			N = N / k
			stream.Send(&pb.PrimesResponse{PrimerNumber: k})
		} else {
			k = k + 1
		}
	}
	return nil
}
