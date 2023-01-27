package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/AndreAppolariFilho/grpc-go-course/blog/blogpb"
)

func doUpdate(c pb.BlogServiceClient, id string) {
	log.Println("doUpdate was invoked")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Changed Author",
		Title:    "My First Blog (edited)",
		Content:  "Content of the first blog, with some awesome additions!",
	}
	_, updateErr := c.UpdateBlog(context.Background(), newBlog)
	if updateErr != nil {
		fmt.Printf("Error happened while updating: %v \n", updateErr)
	}

}
