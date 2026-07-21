package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "NovaCommerce Backend API",
			"version": "0.1.0",
		})
	})

	router.Run(":8080")
}
