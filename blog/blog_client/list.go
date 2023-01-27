package main

import (
	"context"
	"io"
	"log"

	pb "github.com/AndreAppolariFilho/grpc-go-course/blog/blogpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func doList(c pb.BlogServiceClient) {
	log.Println("doList was invoked")
	stream, err := c.ListBlog(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("error while calling ListBlog RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		log.Println(res)
	}
}
