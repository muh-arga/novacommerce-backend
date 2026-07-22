package app

import (
	"github.com/gin-gonic/gin"
	"github.com/muh-arga/novacommerce-backend/configs"
	"github.com/muh-arga/novacommerce-backend/internal/handler"
)

type Application struct {
	Config *configs.Config
	Router *gin.Engine
}

func New() (*Application, error) {
	cfg, err := configs.LoadConfig()

	if err != nil {
		return nil, err
	}

	healthHandler := handler.NewHealthHandler()

	router := NewRouter(healthHandler)

	app := &Application{
		Config: cfg,
		Router: router,
	}

	return app, nil
}
