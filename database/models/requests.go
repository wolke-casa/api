package models

type RequestUser struct {
	User string `json:"user" binding:"required"`
}

type RequestImage struct {
	Domain string `form:"domain" binding:"required"`
}

type RequestGetImage struct {
	Id string `uri:"id" binding:"required"`
}
