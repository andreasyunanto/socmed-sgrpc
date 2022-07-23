package repositories

import (
	"github.com/andreasyunanto/socmed-sgrpc/database"
	"github.com/andreasyunanto/socmed-sgrpc/models"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) GetPostById(id string) (models.Post, error) {

	var post models.Post

	err := database.DB.Where("post_id = ?", id).First(&post).Error
	if err != nil {
		return post, err
	}

	return post, nil

}
