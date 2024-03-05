package web

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/web/middleware"
	"github.com/mattmoran/fyp/api/search/web/handlers"
	"go.uber.org/zap"
)

func InitiateServer() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	authGroup := r.Group("/", middleware.JWTAuthMiddleware())
	{
		authGroup.POST("/simpleSearch", handlers.HandleSimpleSearch)
		authGroup.POST("/simpleSearch/", handlers.HandleSimpleSearch)

		authGroup.POST("/search", handlers.HandleSearch)
		authGroup.POST("/search/", handlers.HandleSearch)

		// TODo: condider getting oonly user attributes
		authGroup.GET("fileAttributes", handlers.HandleGetFileAttributes)
		authGroup.GET("fileAttributes/", handlers.HandleGetFileAttributes)

		authGroup.GET("datasetAttributes", handlers.HandleGetDatasetAttributes)
		authGroup.GET("datasetAttributes/", handlers.HandleGetDatasetAttributes)
	}

	if gin.Mode() == "test" {
		return r
	}

	err := r.Run(":8080")
	if err != nil {
		zap.S().Fatal("Couldn't start server")
	}

	zap.S().Info("Search Server started on port 8005")

	return r
}
