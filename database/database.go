package database

import (
	"log"

	"github.com/andreasyunanto/socmed-sgrpc/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	dsn := configs.Config("DSN")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error database %v", err)
	}
}
