package images

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/wolke-gallery/api/database"
	"github.com/wolke-gallery/api/database/models"
	"github.com/wolke-gallery/api/handlers"
	"gorm.io/gorm"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.Request.Header.Get("Authorization")

		var user models.User
		err := database.Db.First(&user, "key = ?", authorizationHeader).Error

		// TODO: Test that this actually works as record not found
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			error := handlers.ErrInvalidAuthorization

			c.AbortWithStatusJSON(error.Status, gin.H{
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

		c.Next()
	}
}
