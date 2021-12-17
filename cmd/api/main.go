package main

import (
	"fmt"
	"log"
	"os"

	"github.com/wolke-gallery/api/cmd/api/handlers/images"
	"github.com/wolke-gallery/api/cmd/api/handlers/users"

	"github.com/gin-gonic/gin"
	"github.com/wolke-gallery/api/cmd/api/config"
	"github.com/wolke-gallery/api/cmd/api/database"
	"github.com/wolke-gallery/api/cmd/api/handlers"
)

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

	_ = os.Mkdir(config.Config.Directory, os.ModePerm)

	log.Println("🚀 Server starting")

	r := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 10 << 20 // 10 MiB

	r.Static("/images", config.Config.Directory)

	r.Use(gin.Recovery())

	r.GET("/", handlers.Index)

	imagesGroup := r.Group("/images")
	imagesGroup.Use(images.AuthMiddleware())
	imagesGroup.POST("/new", images.NewImage)

	userGroup := r.Group("/users")
	userGroup.Use(users.AuthMiddleware())
	userGroup.POST("/new", users.NewUser)

	port := fmt.Sprintf(":%s", config.Config.Port)

	r.Run(port)
}