package database

import (
	"github.com/wolke-gallery/api/cmd/api/config"
	"github.com/wolke-gallery/api/cmd/api/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() error {
	var err error

	Db, err = gorm.Open(postgres.Open(config.Config.Database), &gorm.Config{})

	return err
}

func Migrate() error {
	err := Db.AutoMigrate(&models.User{})

	return err
}
