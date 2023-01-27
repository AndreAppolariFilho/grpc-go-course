package main

import (
	"context"
	"fmt"
	"log"

	"github.com/AndreAppolariFilho/grpc-go-course/blog/blogpb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) ListBlog(e *emptypb.Empty, stream blogpb.BlogService_ListBlogServer) error {
	log.Println("ListBlog was invoked")
	res, err := collection.Find(context.Background(), primitive.D{{}})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v\n", err))
	}
	defer res.Close(context.Background())
	for res.Next(context.Background()) {
		elem := &BlogItem{}
		err := res.Decode(elem)
		if err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v\n", err))
		}
		stream.Send(documentToBlog(elem))
	}
	if err = res.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v\n", err))
	}
	return nil
}
