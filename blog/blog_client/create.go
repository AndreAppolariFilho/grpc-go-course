package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/AndreAppolariFilho/grpc-go-course/blog/blogpb"
)

func doCreate(c pb.BlogServiceClient) string {
	log.Println("doCreate was invoked")
	blog := &pb.Blog{
		AuthorId: "André",
		Title:    "My First Blog",
		Content:  "Content of the first blog",
	}
	createBlogRes, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("Blog has been created: %v", createBlogRes)
	blogID := createBlogRes.GetId()
	log.Printf("Create Blog %s\n", blogID)
	return blogID
}
