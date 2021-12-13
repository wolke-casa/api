package users

import (
	"github.com/gin-gonic/gin"
	"github.com/wolke-gallery/api/cmd/api/config"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.Request.Header.Get("Authorization")

		if config.Config.BotApiToken != authorizationHeader {
			c.AbortWithStatusJSON(401, gin.H{
				"success": false,
				"message": "Invalid authorization token provided",
			})
			return
		}

		c.Next()
	}
}
