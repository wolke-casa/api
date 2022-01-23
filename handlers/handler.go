package handlers

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wolke-gallery/api/config"
)

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
		"message": "hello world",
	})
}

func GetDomains(c *gin.Context) {
	domains := strings.Join(config.Config.Domains, ",")

	c.JSON(200, gin.H{
		"success": true,
		"message": domains,
	})
}
