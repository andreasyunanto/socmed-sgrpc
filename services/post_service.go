package services

import (
	"errors"

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
		PostId:   operationResult.PostId,
		Contents: operationResult.Contents,
	}

	return &pb.GetPostResponse{
		Message: "OK",
		Data:    data,
		Status:  true,
	}, nil
}

func CreatePost(repo *repositories.PostRepository, req *pb.AddPostRequest) (*pb.GetPostResponse, error) {
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
