package server

import (
	"context"

	"github.com/andreasyunanto/socmed-sgrpc/pb"
	"github.com/andreasyunanto/socmed-sgrpc/repositories"
	"github.com/andreasyunanto/socmed-sgrpc/services"
)

type PostServer struct{}

// Get Post By ID
func (s *PostServer) GetPostById(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	return services.GetPostById(&repositories.PostRepository{}, req.PostId)
}

// Create Post
func (s *PostServer) AddPost(ctx context.Context, req *pb.AddPostRequest) (*pb.GetPostResponse, error) {
	return services.CreatePost(&repositories.PostRepository{}, req)
}
