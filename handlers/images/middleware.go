package images

import (
	"github.com/gin-gonic/gin"
	"github.com/wolke-gallery/api/database"
	"github.com/wolke-gallery/api/database/models"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.Request.Header.Get("Authorization")

		var user models.User
		result := database.Db.First(&user, "key = ?", authorizationHeader)

		if result.Error != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"success": false,
				"message": "Invalid authorization token provided",
			})
			return
		}

		c.Next()
	}
}
