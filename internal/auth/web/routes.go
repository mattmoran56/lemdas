package web

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/auth/web/handlers"
	"github.com/mattmoran/fyp/api/pkg/web/middleware"
	"go.uber.org/zap"
)

func InitiateServer() {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	r.POST("/token", handlers.HandleToken)
	r.POST("/verify", handlers.HandleVerify)

	err := r.Run(":8001")
	if err != nil {
		zap.S().Fatal("Couldn't start server")
	}

	zap.S().Info("Auth Server started on port 8001")
}
