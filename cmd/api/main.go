package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wolke-gallery/api/cmd/api/config"
	"github.com/wolke-gallery/api/cmd/api/database"
	"github.com/wolke-gallery/api/cmd/api/database/models"
	"github.com/wolke-gallery/api/cmd/api/handlers"
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

// TODO: Convert all responses to JSON
func main() {
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	if err := database.Migrate(); err != nil {
		log.Fatal(err)
	}

	log.Println("🚀 Server starting")

	r := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 10 << 20 // 10 MiB

	r.Use(gin.Recovery())

	r.GET("/", handlers.Index)

	imagesGroup := r.Group("/images")

	imagesGroup.Use(AuthMiddleware())

	imagesGroup.POST("/new", handlers.NewImage)

	port := fmt.Sprintf(":%s", config.Config.Port)

	r.Run(port)
}
