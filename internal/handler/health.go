package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muh-arga/novacommerce-backend/configs"
	"github.com/muh-arga/novacommerce-backend/internal/shared/response"
)

type HealthHandler struct {
	config *configs.Config
}

func NewHealthHandler(config *configs.Config) *HealthHandler {
	return &HealthHandler{
		config: config,
	}
}

func (h *HealthHandler) Index(c *gin.Context) {
	data := gin.H{
		"status":      "ok",
		"service":     h.config.App.Name,
		"environment": h.config.App.Env,
		"version":     "0.1.0",
	}

	response.Success(c, http.StatusOK, data)
}
