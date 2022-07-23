package models

import (
	"time"
)

type Comment struct {
	CommentId string    `gorm:"primaryKey;type:varchar(36);" json:"commentId"`
	PostId    string    `gorm:"type:varchar(36);index;" json:"postId"`
	Comments  string    `gorm:"type:text;" json:"comments"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Comments []Comment
