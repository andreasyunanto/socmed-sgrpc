package models

import (
	"time"
)

type Post struct {
	PostId    string    `gorm:"primaryKey;type:varchar(36);" json:"postId"`
	UserId    string    `gorm:"type:varchar(36);index;" json:"userId"`
	Contents  string    `gorm:"type:text;" json:"contents"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Posts []Post
