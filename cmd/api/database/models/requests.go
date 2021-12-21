package models

type RequestUser struct {
	User string `json:"user" binding:"required"`
}

type RequestImage struct {
	Domain string `form:"domain" binding:"required"`
}
