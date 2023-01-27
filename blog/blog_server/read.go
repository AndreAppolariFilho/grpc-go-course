package main

import (
	"context"
	"log"

	pb "github.com/AndreAppolariFilho/grpc-go-course/blog/blogpb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"
)

func (s *server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {
	log.Printf("ReadBlog was invoked: %v\n", in)
	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Cannot parse Id")
	}
	data := &BlogItem{}
	filter := bson.M{"_id": oid}
	res := collection.FindOne(ctx, filter)
	err = res.Decode(data)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Cannot find blog with the id provided")
	}
	return documentToBlog(data), nil
}
