package handlers

import (
	"fmt"

	"github.com/wolke-gallery/api/cmd/api/config"
	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
		"message": "hello world",
	})
}

func NewImage(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "No file uploaded",
		})
		return
	}

	if file.Size > config.Config.MaxFileSize {
		c.JSON(400, gin.H{
			"success": false,
			"message": "File size too large",
		})
		return
	}

	contentType := file.Header["Content-Type"][0]
	var extension string

	switch contentType {
	case "image/jpeg":
		extension = "jpg"
	case "image/png":
		extension = "png"
	case "image/gif":
		extension = "gif"
	default:
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid file type",
		})
		return
	}

	id, err := gonanoid.New()

	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Failed to generate id.. please try again",
		})
		return
	}

	name := fmt.Sprintf("%s.%s", id, extension)

	if err := c.SaveUploadedFile(file, config.Config.Directory+name); err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Failed to save file",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "uwu",
		"name":    name,
	})
}
