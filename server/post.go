package server

import (
	"context"

	"github.com/andreasyunanto/socmed-sgrpc/pb"
	"github.com/andreasyunanto/socmed-sgrpc/repositories"
	"github.com/andreasyunanto/socmed-sgrpc/services"
)

type PostServer struct{}

func (s *PostServer) GetPostById(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostResponse, error) {

	postId := req.PostId

	return services.GetPostById(&repositories.PostRepository{}, postId)

}
