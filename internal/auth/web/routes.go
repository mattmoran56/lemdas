package web

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/auth/web/handlers"
	"github.com/mattmoran/fyp/api/pkg/web/middleware"
	"go.uber.org/zap"
	"os"
)

func InitiateServer() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	r.POST("/token", handlers.HandleToken)
	r.POST("/verify", handlers.HandleVerify)

	if gin.Mode() == "test" {
		return r
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := r.Run(":" + port)
	if err != nil {
		zap.S().Fatal("Couldn't start server")
	}

	zap.S().Info("Auth Server started on port 8001")

	return r
}
