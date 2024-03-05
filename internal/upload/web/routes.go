package web

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/web/middleware"
	"github.com/mattmoran/fyp/api/upload/web/handlers"
	"go.uber.org/zap"
)

func InitiateServer() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	authGroup := r.Group("/", middleware.JWTAuthMiddleware())
	{
		authGroup.POST("/upload", handlers.HandleUpload)
	}

	if gin.Mode() == "test" {
		return r
	}

	err := r.Run(":8080")
	if err != nil {
		zap.S().Fatal("Couldn't start server")
	}

	zap.S().Info("Upload Server started on port 8002")

	return r
}
