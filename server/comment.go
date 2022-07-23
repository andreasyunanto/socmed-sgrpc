package server

import (
	"context"

	"github.com/andreasyunanto/socmed-sgrpc/pb"
	"github.com/andreasyunanto/socmed-sgrpc/repositories"
	"github.com/andreasyunanto/socmed-sgrpc/services"
)

type CommentServer struct{}

func (s *CommentServer) GetCommentByPost(ctx context.Context, req *pb.GetCommentsRequest) (*pb.GetCommentsResponse, error) {

	postId := req.PostId

	return services.GetCommentByPost(&repositories.CommentRepository{}, postId)

}
