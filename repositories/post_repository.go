package repositories

import (
	"github.com/andreasyunanto/socmed-sgrpc/database"
	"github.com/andreasyunanto/socmed-sgrpc/models"
	"github.com/andreasyunanto/socmed-sgrpc/pb"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

// Get Post By Id
func (r *PostRepository) GetPostById(id string) (models.Post, error) {

	var post models.Post

	err := database.DB.Where("post_id = ?", id).First(&post).Error
	if err != nil {
		return post, err
	}

	return post, nil

}

// Create Socmed Post
func (r *PostRepository) CreatePost(req *pb.AddPostRequest) (models.Post, error) {

	var post models.Post

	post.PostId = uuid.New().String()
	post.Contents = req.Contents
	err := database.DB.Save(&post).Error
	if err != nil {
		return post, err
	}

	return post, nil
}
