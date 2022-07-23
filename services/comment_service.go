package services

import (
	"errors"
	"fmt"

	"github.com/andreasyunanto/socmed-sgrpc/pb"
	"github.com/andreasyunanto/socmed-sgrpc/repositories"
)

func GetCommentByPost(repo *repositories.CommentRepository, id string) (*pb.GetCommentsResponse, error) {

	fmt.Println("Call GetCommentByPost on Socmed Server via gRPC with param:", id)

	operationResult, err := repo.GetCommentByPost(id)

	if err != nil {
		return &pb.GetCommentsResponse{Status: false, Data: nil, Message: err.Error()}, errors.New(err.Error())
	}

	comments := make([]*pb.Comment, 0)
	for _, val := range operationResult {
		comments = append(comments, &pb.Comment{
			CommentId: val.CommentId,
			PostId:    val.PostId,
			Comments:  val.Comments,
		})
	}

	c := pb.CommentData{
		Items: comments,
	}

	return &pb.GetCommentsResponse{
		Message: "OK",
		Data:    &c,
		Status:  true,
	}, nil
}
