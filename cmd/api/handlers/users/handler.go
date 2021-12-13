package users

import (
	"github.com/gin-gonic/gin"
	"github.com/wolke-gallery/api/cmd/api/config"
	"github.com/wolke-gallery/api/cmd/api/database"
	"github.com/wolke-gallery/api/cmd/api/database/models"
	"github.com/wolke-gallery/api/cmd/api/utils"
)

func NewUser(c *gin.Context) {

	var user models.RequestUser

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid user",
		})
		return
	}

	key := utils.GenerateSecureToken(config.Config.KeyLength)

	newUser := models.User{User: user.User, Key: key}

	// This catches *all* errors when creating
	if err := database.Db.Create(&newUser).Error; err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "That user already has an API key.",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": newUser.Key,
	})
}
