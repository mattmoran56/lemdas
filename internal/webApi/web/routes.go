package web

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/web/middleware"
	"github.com/mattmoran/fyp/api/webApi/web/handlers/dataset"
	"github.com/mattmoran/fyp/api/webApi/web/handlers/file"
	"github.com/mattmoran/fyp/api/webApi/web/handlers/user"
	m "github.com/mattmoran/fyp/api/webApi/web/middleware"
	"go.uber.org/zap"
)

func InitiateServer() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	authGroup := r.Group("/", middleware.JWTAuthMiddleware())
	{
		authGroup.GET("user/search", user.HandleSearchForUser)
		authGroup.GET("user/search/", user.HandleSearchForUser)

		authGroup.GET("dataset", dataset.HandleGetDatasets)
		authGroup.GET("dataset/", dataset.HandleGetDatasets)

		authGroup.GET("datasets/stared", dataset.HandleGetStaredDatasets)
		authGroup.GET("datasets/stared/", dataset.HandleGetStaredDatasets)

		authGroup.POST("dataset", dataset.HandleCreateDataset)
		authGroup.POST("dataset/", dataset.HandleCreateDataset)

		datasetGroup := authGroup.Group("/dataset/:datasetId", m.CheckDatasetAccess())
		{

			datasetGroup.GET("", dataset.HandleGetDataset)
			datasetGroup.GET("/", dataset.HandleGetDataset)

			datasetGroup.PUT("", dataset.HandleUpdateDataset)
			datasetGroup.PUT("/", dataset.HandleUpdateDataset)

			datasetGroup.DELETE("", dataset.HandleDeleteDataset)
			datasetGroup.DELETE("/", dataset.HandleDeleteDataset)

			datasetGroup.GET("/files", dataset.HandleGetFiles)
			datasetGroup.GET("/files/", dataset.HandleGetFiles)

			datasetGroup.GET("/stared", dataset.HandleGetStared)
			datasetGroup.GET("/stared/", dataset.HandleGetStared)

			datasetGroup.POST("/stared", dataset.HandleStarDataset)
			datasetGroup.POST("/stared/", dataset.HandleStarDataset)

			datasetGroup.GET("/collaborator", dataset.HandleGetCollaborators)
			datasetGroup.GET("/collaborator/", dataset.HandleGetCollaborators)

			datasetGroup.POST("/collaborator", dataset.HandleAddCollaborator)
			datasetGroup.POST("/collaborator/", dataset.HandleAddCollaborator)

			datasetGroup.DELETE("/collaborator/:collaboratorId", dataset.HandleDeleteCollaborator)
			datasetGroup.DELETE("/collaborator/:collaboratorId/", dataset.HandleDeleteCollaborator)

			datasetAttributesGroup := datasetGroup.Group("/attribute")
			{
				datasetAttributesGroup.GET("", dataset.HandleGetAttributes)
				datasetAttributesGroup.GET("/", dataset.HandleGetAttributes)

				datasetAttributesGroup.POST("", dataset.HandleCreateAttribute)
				datasetAttributesGroup.POST("/", dataset.HandleCreateAttribute)

				datasetAttributeGroup := datasetAttributesGroup.Group("/:datasetAttributeId", m.CheckDatasetAttributeAccess())
				{
					datasetAttributeGroup.PUT("", dataset.HandleUpdateAttribute)
					datasetAttributeGroup.PUT("/", dataset.HandleUpdateAttribute)

					datasetAttributeGroup.DELETE("", dataset.HandleDeleteAttribute)
					datasetAttributeGroup.DELETE("/", dataset.HandleDeleteAttribute)
				}
			}
		}

		fileGroup := authGroup.Group("/file/:fileId", m.CheckFileAccess())
		{
			fileGroup.GET("", file.HandleGetFile)
			fileGroup.GET("/", file.HandleGetFile)

			fileGroup.PUT("", file.HandleUpdateFile)
			fileGroup.PUT("/", file.HandleUpdateFile)

			fileGroup.DELETE("", file.HandleDeleteFile)
			fileGroup.DELETE("/", file.HandleDeleteFile)

			fileGroup.GET("/preview", file.HandlePreview)
			fileGroup.GET("/preview/", file.HandlePreview)

			fileAttributesGroup := fileGroup.Group("/attribute")
			{
				fileAttributesGroup.GET("", file.HandleGetAttributes)
				fileAttributesGroup.GET("/", file.HandleGetAttributes)

				fileAttributesGroup.POST("", file.HandleCreateAttribute)
				fileAttributesGroup.POST("/", file.HandleCreateAttribute)

				fileAttributeGroup := fileAttributesGroup.Group("/:fileAttributeId", m.CheckFileAttributeAccess())
				{
					fileAttributeGroup.PUT("", file.HandleUpdateAttribute)
					fileAttributeGroup.PUT("/", file.HandleUpdateAttribute)

					fileAttributeGroup.DELETE("", file.HandleDeleteAttribute)
					fileAttributeGroup.DELETE("/", file.HandleDeleteAttribute)
				}
			}
		}
	}

	if gin.Mode() == "test" {
		return r
	}

	err := r.Run(":8003")
	if err != nil {
		zap.S().Fatal("Couldn't start server")
	}

	zap.S().Info("Web API Server started on port 8003")

	return r
}
