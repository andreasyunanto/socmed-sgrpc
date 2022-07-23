package database

import (
	"github.com/andreasyunanto/socmed-sgrpc/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {

	db.AutoMigrate(
		&models.Post{},
		&models.Comment{},
	)
}
