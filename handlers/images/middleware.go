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

		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			error, status := handlers.ErrInvalidAuthorization()

			c.AbortWithStatusJSON(status, gin.H{
				"success": false,
				"message": error,
			})
			return
		}

		if err != nil {
			error, status := handlers.ErrUnknownErrorOccurred()

			c.AbortWithStatusJSON(status, gin.H{
				"success": false,
				"message": error,
			})
			return
		}

		c.Next()
	}
}
