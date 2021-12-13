package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	ID   string `gorm:"type:uuid;default:gen_random_uuid();autoIncrement"`
	User string `gorm:"unique"`
	Key  string `gorm:"type:varchar(2048);unique"`
}

type RequestUser struct {
	User string `json:"user"  binding:"required"`
}
