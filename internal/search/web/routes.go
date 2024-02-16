package web

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/web/middleware"
	"github.com/mattmoran/fyp/api/search/web/handlers"
	"go.uber.org/zap"
)

func InitiateServer() {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	authGroup := r.Group("/", middleware.JWTAuthMiddleware())
	{
		authGroup.POST("/simpleSearch", handlers.HandleSimpleSearch)
		authGroup.POST("/simpleSearch/", handlers.HandleSimpleSearch)
	}

	err := r.Run(":8005")
	if err != nil {
		zap.S().Fatal("Couldn't start server")
	}

	zap.S().Info("Search Server started on port 8005")
}
