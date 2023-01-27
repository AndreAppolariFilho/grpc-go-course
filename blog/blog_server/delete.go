package main

import (
	"context"
	"log"

	"github.com/AndreAppolariFilho/grpc-go-course/blog/blogpb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gopkg.in/mgo.v2/bson"
)

func (s *server) DeleteBlog(ctx context.Context, in *blogpb.BlogId) (*emptypb.Empty, error) {
	log.Printf("UpdateBlog was invoked: %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse the id",
		)
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Could not delete",
		)
	}
	if res.DeletedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "Blog not found")
	}
	return &emptypb.Empty{}, nil
}
