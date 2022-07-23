package services

import (
	"errors"
	"fmt"

	"github.com/andreasyunanto/socmed-sgrpc/pb"
	"github.com/andreasyunanto/socmed-sgrpc/repositories"
)

// Get Post By Id
func GetPostById(repo *repositories.PostRepository, id string) (*pb.GetPostResponse, error) {
	operationResult, err := repo.GetPostById(id)

	if err != nil {
		return &pb.GetPostResponse{Status: false, Data: nil, Message: err.Error()}, errors.New(err.Error())
	}

	var data *pb.Post = &pb.Post{
		UserId:   operationResult.UserId,
		PostId:   operationResult.PostId,
		Contents: operationResult.Contents,
	}

	return &pb.GetPostResponse{
		Message: "OK",
		Data:    data,
		Status:  true,
	}, nil
}

// Create Post
func CreatePost(repo *repositories.PostRepository, req *pb.AddPostRequest) (*pb.GetPostResponse, error) {

	fmt.Println("Call CreatePost on Socmed Server via gRPC with param:", req)

	operationResult, err := repo.CreatePost(req)
	if err != nil {
	}

	var data *pb.Post = &pb.Post{
		PostId:   operationResult.PostId,
		UserId:   operationResult.UserId,
		Contents: operationResult.Contents,
	}

	return &pb.GetPostResponse{
		Message: "OK",
		Data:    data,
		Status:  true,
	}, nil
}
