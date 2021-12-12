package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	ID    int    `gorm:"primaryKey;autoIncrement"`
	User  int    `gorm:"unique"`
	Key string `gorm:"type:varchar(2048)"`
}
