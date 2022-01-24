package users

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wolke-gallery/api/config"
	"github.com/wolke-gallery/api/database"
	"github.com/wolke-gallery/api/database/models"
	"github.com/wolke-gallery/api/handlers"
	"github.com/wolke-gallery/api/utils"
)

func NewUser(c *gin.Context) {
	var data models.RequestUser

	if err := c.ShouldBindJSON(&data); err != nil {
		error := handlers.ErrInvalidInRequest

		c.JSON(error.Status, gin.H{
			"success": false,
			"message": strings.Replace(error.Error, "{}", "user", 1),
		})
		return
	}

	key := utils.GenerateSecureToken(config.Config.KeyLength)

	newUser := models.User{User: data.User, Key: key}

	/*
		curl -X POST http://localhost:8080/users/new \
		-H 'Authorization: TOKEN' \
		-H 'Content-Type: application/json' \
		-d '{"user": "300088143422685185"}'
	*/

	err := database.Db.Create(&newUser).Error

	if err != nil && strings.Contains(err.Error(), "SQLSTATE 23505") {
		error := handlers.ErrUserAlreadyHasApiKey

		c.JSON(error.Status, gin.H{
			"success": false,
			"message": error.Error,
		})
		return
	}

	if err != nil {
		error := handlers.ErrUnknownErrorOccurred

		c.JSON(error.Status, gin.H{
			"success": false,
			"message": error.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": newUser.Key,
	})
}
