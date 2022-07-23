package repositories

import (
	"github.com/andreasyunanto/socmed-sgrpc/database"
	"github.com/andreasyunanto/socmed-sgrpc/models"
	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) GetCommentByPost(id string) (models.Comments, error) {

	var comments models.Comments

	err := database.DB.Where("post_id = ?", id).Find(&comments).Error
	if err != nil {
		return comments, err
	}

	return comments, nil

}
