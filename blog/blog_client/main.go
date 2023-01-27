package main

import (
	"fmt"
	"log"

	pb "github.com/AndreAppolariFilho/grpc-go-course/blog/blogpb"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Blog Client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close() // Maybe this should be in a separate function and the error handled?

	c := pb.NewBlogServiceClient(cc)
	id := doCreate(c)
	doRead(c, id)
	doUpdate(c, id)
	doList(c)
	doDelete(c, id)
}
