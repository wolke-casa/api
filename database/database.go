package database

import (
	"github.com/wolke-gallery/api/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect(databaseUrl string) error {
	var err error

	Db, err = gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})

	return err
}

func Migrate() error {
	err := Db.AutoMigrate(&models.User{})

	return err
}
