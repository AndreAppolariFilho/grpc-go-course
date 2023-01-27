package main

import (
	pb "github.com/AndreAppolariFilho/grpc-go-course/blog/blogpb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogItem struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorId string             `bson: "author_id"`
	Title    string             `bson: "title"`
	Content  string             `bson: "content"`
}

func documentToBlog(data *BlogItem) *pb.Blog {
	return &pb.Blog{
		Id:       data.Id.Hex(),
		AuthorId: data.AuthorId,
		Title:    data.Title,
		Content:  data.Content,
	}
}
