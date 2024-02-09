package web

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/web/middleware"
	"github.com/mattmoran/fyp/api/upload/web/handlers"
	"go.uber.org/zap"
)

func InitiateServer() {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	authGroup := r.Group("/", middleware.JWTAuthMiddleware())
	{
		authGroup.POST("/upload", handlers.HandleUpload)
	}

	err := r.Run(":8002")
	if err != nil {
		zap.S().Fatal("Couldn't start server")
	}

	zap.S().Info("Upload Server started on port 8002")
}
