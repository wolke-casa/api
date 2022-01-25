package users

import (
	"github.com/gin-gonic/gin"
	"github.com/wolke-gallery/api/config"
	"github.com/wolke-gallery/api/handlers"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.Request.Header.Get("Authorization")

		if config.Config.BotApiKey != authorizationHeader {
			error, status := handlers.ErrInvalidAuthorization()

			c.AbortWithStatusJSON(status, gin.H{
				"success": false,
				"message": error,
			})
			return
		}

		c.Next()
	}
}
