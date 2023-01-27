package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/AndreAppolariFilho/grpc-go-course/calculator/calculatorpb"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax was invoked")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Printf("Error while parsing the stream: %v\n", err)
	}
	reqs := []*pb.MaxRequest{
		{Number: 5},
		{Number: 1},
		{Number: 2},
		{Number: 6},
		{Number: 9},
		{Number: 7},
	}
	waitc := make(chan struct{})
	go func() {
		for _, req := range reqs {
			log.Printf("Sending Request: %v", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()
	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Panicf("Error while reading from stream: %v\n")
				break
			}
			log.Printf("Received: %v\n", res.Response)

		}
		close(waitc)
	}()
	<-waitc

}
