package main

import (
	"context"
	"log"

	pb "github.com/AndreAppolariFilho/grpc-go-course/blog/blogpb"
)

func doDelete(c pb.BlogServiceClient, blog_id string) {
	log.Println("doCreate was invoked")

	_, deleteErr := c.DeleteBlog(context.Background(), &pb.BlogId{Id: blog_id})
	if deleteErr != nil {
		log.Fatalf("Error happened while deleting: %v \n", deleteErr)
	}
	log.Println("Deleted")

}
