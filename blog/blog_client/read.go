package main

import (
	"context"
	"log"

	pb "github.com/AndreAppolariFilho/grpc-go-course/blog/blogpb"
)

func doRead(c pb.BlogServiceClient, blog_id string) *pb.Blog {
	log.Println("doRead was invoked")
	readBlogReq := &pb.BlogId{Id: blog_id}
	readBlogRes, readBlogErr := c.ReadBlog(context.Background(), readBlogReq)
	if readBlogErr != nil {
		log.Fatalf("Error happened while reading: %v \n", readBlogErr)
	}

	log.Printf("Blog was read: %v \n", readBlogRes)
	return readBlogRes
}
