package web

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/web/middleware"
	"github.com/mattmoran/fyp/api/webApi/web/handlers/dataset"
	"github.com/mattmoran/fyp/api/webApi/web/handlers/dataset/share"
	"github.com/mattmoran/fyp/api/webApi/web/handlers/feed"
	"github.com/mattmoran/fyp/api/webApi/web/handlers/file"
	"github.com/mattmoran/fyp/api/webApi/web/handlers/group"
	"github.com/mattmoran/fyp/api/webApi/web/handlers/user"
	"go.uber.org/zap"
	"os"
)

func InitiateServer() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	authGroup := r.Group("/", middleware.JWTAuthMiddleware())
	{
		authGroup.GET("group/search", group.HandleSearchForGroup)
		authGroup.GET("group/search/", group.HandleSearchForGroup)

		authGroup.GET("dataset", dataset.HandleGetDatasets)
		authGroup.GET("dataset/", dataset.HandleGetDatasets)

		authGroup.GET("datasets/stared", dataset.HandleGetStaredDatasets)
		authGroup.GET("datasets/stared/", dataset.HandleGetStaredDatasets)

		authGroup.GET("datasets/shared", dataset.HandleGetSharedDatasets)
		authGroup.GET("datasets/shared/", dataset.HandleGetSharedDatasets)

		authGroup.POST("dataset", dataset.HandleCreateDataset)
		authGroup.POST("dataset/", dataset.HandleCreateDataset)

		authGroup.POST("/dataset/:datasetId/stared", dataset.HandleStarDataset)
		authGroup.POST("/dataset/:datasetId/stared/", dataset.HandleStarDataset)

		datasetGroup := authGroup.Group("/dataset/:datasetId", middleware.CheckDatasetAccess())
		{

			datasetGroup.GET("", dataset.HandleGetDataset)
			datasetGroup.GET("/", dataset.HandleGetDataset)

			datasetGroup.PUT("", dataset.HandleUpdateDataset)
			datasetGroup.PUT("/", dataset.HandleUpdateDataset)

			datasetGroup.DELETE("", dataset.HandleDeleteDataset)
			datasetGroup.DELETE("/", dataset.HandleDeleteDataset)

			datasetGroup.GET("/access", dataset.HandleGetAccessLevel)
			datasetGroup.GET("/access/", dataset.HandleGetAccessLevel)

			datasetGroup.GET("/files", dataset.HandleGetFiles)
			datasetGroup.GET("/files/", dataset.HandleGetFiles)

			datasetGroup.GET("/stared", dataset.HandleGetStared)
			datasetGroup.GET("/stared/", dataset.HandleGetStared)

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

				datasetAttributeGroup := datasetAttributesGroup.Group("/:datasetAttributeId", middleware.CheckDatasetAttributeAccess())
				{
					datasetAttributeGroup.PUT("", dataset.HandleUpdateAttribute)
					datasetAttributeGroup.PUT("/", dataset.HandleUpdateAttribute)

					datasetAttributeGroup.DELETE("", dataset.HandleDeleteAttribute)
					datasetAttributeGroup.DELETE("/", dataset.HandleDeleteAttribute)
				}
			}

			sharingGroup := datasetGroup.Group("/share")
			{
				// TODO: Test all these routes
				sharingGroup.POST("/user", share.HandleShareWithUser)
				sharingGroup.POST("/user/", share.HandleShareWithUser)

				sharingGroup.DELETE("/user/:userId", share.HandleDeleteUserShare)
				sharingGroup.DELETE("/user/:userId/", share.HandleDeleteUserShare)

				sharingGroup.POST("/group", share.HandleShareWithGroup)
				sharingGroup.POST("/group/", share.HandleShareWithGroup)

				sharingGroup.DELETE("/group/:groupId", share.HandleDeleteGroupShare)
				sharingGroup.DELETE("/group/:groupId/", share.HandleDeleteGroupShare)

				sharingGroup.GET("/user", share.HandleGetUsersSharedWith)
				sharingGroup.GET("/user/", share.HandleGetUsersSharedWith)

				sharingGroup.GET("/group", share.HandleGetGroupsSharedWith)
				sharingGroup.GET("/group/", share.HandleGetGroupsSharedWith)
			}
		}

		fileGroup := authGroup.Group("/file/:fileId", middleware.CheckFileAccess())
		{
			fileGroup.GET("", file.HandleGetFile)
			fileGroup.GET("/", file.HandleGetFile)

			fileGroup.PUT("", file.HandleUpdateFile)
			fileGroup.PUT("/", file.HandleUpdateFile)

			fileGroup.DELETE("", file.HandleDeleteFile)
			fileGroup.DELETE("/", file.HandleDeleteFile)

			fileGroup.GET("/access", file.HandleGetAccessLevel)
			fileGroup.GET("/access/", file.HandleGetAccessLevel)

			fileGroup.GET("/preview", file.HandlePreview)
			fileGroup.GET("/preview/", file.HandlePreview)

			fileAttributesGroup := fileGroup.Group("/attribute")
			{
				fileAttributesGroup.GET("", file.HandleGetAttributes)
				fileAttributesGroup.GET("/", file.HandleGetAttributes)

				fileAttributesGroup.POST("", file.HandleCreateAttribute)
				fileAttributesGroup.POST("/", file.HandleCreateAttribute)

				fileAttributeGroup := fileAttributesGroup.Group("/:fileAttributeId", middleware.CheckFileAttributeAccess())
				{
					fileAttributeGroup.PUT("", file.HandleUpdateAttribute)
					fileAttributeGroup.PUT("/", file.HandleUpdateAttribute)

					fileAttributeGroup.DELETE("", file.HandleDeleteAttribute)
					fileAttributeGroup.DELETE("/", file.HandleDeleteAttribute)
				}
			}
		}

		authGroup.GET("group", group.HandleGetGroups)
		authGroup.GET("group/", group.HandleGetGroups)

		authGroup.POST("group", group.HandleCreateGroup)
		authGroup.POST("group/", group.HandleCreateGroup)

		groupsGroup := authGroup.Group("/group/:groupId", middleware.CheckGroupAccess())
		{
			groupsGroup.GET("", group.HandleGetGroup)
			groupsGroup.GET("/", group.HandleGetGroup)

			groupsGroup.DELETE("", group.HandleDeleteGroup)
			groupsGroup.DELETE("/", group.HandleDeleteGroup)

			// TODO: This needs testing
			groupsGroup.GET("/member", group.HandleGetMembers)
			groupsGroup.GET("/member/", group.HandleGetMembers)

			groupsGroup.POST("/member", group.HandleAddMember)
			groupsGroup.POST("/member/", group.HandleAddMember)

			groupsGroup.DELETE("/member/:userId", group.HandleDeleteMember)
			groupsGroup.DELETE("/member/:userId/", group.HandleDeleteMember)
		}

		feedGroup := authGroup.Group("/feed")
		{
			feedGroup.GET("/:userId", feed.HandleGetProfileFeed)
			feedGroup.GET("/:userId/", feed.HandleGetProfileFeed)
		}

		userGroup := authGroup.Group("/user")
		{
			userGroup.GET("/search", user.HandleSearchForUser)
			userGroup.GET("/search/", user.HandleSearchForUser)

			userGroup.GET("/profile/:userId", user.HandleGetProfile)
			userGroup.GET("/profile/:userId/", user.HandleGetProfile)
		}

	}

	if gin.Mode() == "test" {
		return r
	}

	port, set := os.LookupEnv("PORT")
	if !set {
		port = "8080"
	}

	err := r.Run(":" + port)
	if err != nil {
		zap.S().Fatal("Couldn't start server")
	}

	zap.S().Info("Web API Server started on port 8003")

	return r
}
