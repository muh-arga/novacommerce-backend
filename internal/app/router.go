package app

import (
	"github.com/gin-gonic/gin"
	"github.com/muh-arga/novacommerce-backend/internal/handler"
)

func NewRouter(healthHandler *handler.HealthHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/", healthHandler.Index)

	return router
}
