package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muh-arga/novacommerce-backend/configs"
)

func main() {
	cfg, err := configs.LoadConfig()

	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "NovaCommerce Backend API",
			"version": "0.1.0",
		})
	})

	router.Run(fmt.Sprintf(":%d", cfg.App.Port))
}
