package services

import (
	"errors"

	"github.com/andreasyunanto/socmed-sgrpc/pb"
	"github.com/andreasyunanto/socmed-sgrpc/repositories"
)

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
