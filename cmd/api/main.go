package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/wolke-gallery/api/config"
	"github.com/wolke-gallery/api/database"
	"github.com/wolke-gallery/api/handlers"
	"github.com/wolke-gallery/api/handlers/images"
	"github.com/wolke-gallery/api/handlers/users"
	"github.com/wolke-gallery/api/medium"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}

	if err := database.Connect(config.Config.DatabaseUrl); err != nil {
		log.Fatal(err)
	}

	if err := database.Migrate(); err != nil {
		log.Fatal(err)
	}

	if config.Config.Medium == "local" {
		_ = os.Mkdir(config.Config.Directory, os.ModePerm)
	}

	if err := medium.Initialize(); err != nil {
		log.Fatal(err)
	}

	log.Println("🚀 Server starting")

	r := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 10 << 20 // 10 MiB

	r.Use(gin.Recovery())

	r.GET("/", handlers.Index)
	r.GET("/domains", handlers.GetDomains)

	imagesGroup := r.Group("/images")
	imagesGroup.Use(images.AuthMiddleware())
	imagesGroup.POST("/new", images.NewImage)

	r.GET("/images/:id", images.GetImage)

	userGroup := r.Group("/users")
	userGroup.Use(users.AuthMiddleware())
	userGroup.POST("/new", users.NewUser)

	_ = r.Run(config.Config.Port)
}
