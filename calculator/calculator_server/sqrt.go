package main

import (
	"context"
	"fmt"
	"log"
	"math"

	pb "github.com/AndreAppolariFilho/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Sqrt was invoked")
	if in.Number < 0 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Received a negative number: %d", in.Number))
	}
	res := math.Sqrt(float64(in.Number))
	return &pb.SqrtResponse{
		Result: res,
	}, nil
}
