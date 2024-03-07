package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"github.com/mattmoran/fyp/api/pkg/utils"
	"net/http"
)

func CheckDatasetAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		datasetID := c.Param("datasetId")
		userID := c.GetString("userID")

		access := "none"

		dataset, err := database.DatasetRepo.GetDatasetByID(datasetID)
		if err != nil && err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Dataset not found"})
			c.Abort()
			return
		}
		utils.HandleHandlerError(c, err)

		if dataset.OwnerID != userID {
			userShare, err := database.UserShareDatasetRepo.GetUserShareDatasetForDatasetIdAndUserId(datasetID, userID)
			if err != nil && err.Error() != "record not found" {
				utils.HandleHandlerError(c, err)
			}
			if err != nil && err.Error() == "record not found" {
				groups, err := database.GroupMemberRepo.GetGroupsForUser(userID)
				if err != nil && err.Error() != "record not found" {
					utils.HandleHandlerError(c, err)
				}
				for _, group := range groups {
					groupShare, err := database.GroupShareDatasetRepo.GetGroupShareDatasetForDatasetIdAndGroupId(datasetID, group.ID)
					if err != nil && err.Error() != "record not found" {
						utils.HandleHandlerError(c, err)
					}
					if err == nil {
						if groupShare.WriteAccess {
							access = "write"
						} else {
							access = "read"
						}
						break
					}
				}
			} else {
				if userShare.WriteAccess {
					access = "write"
				} else {
					access = "read"
				}
			}
		} else {
			access = "write"
		}
		if access == "none" {
			if dataset.IsPublic {
				access = "read"
			}
		}
		c.Set("access", access)

		if access == "none" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Dataset not found"})
			c.Abort()
			return
		}
		if access == "write" {
			c.Next()
			return
		}
		if access == "read" {
			if c.Request.Method == "GET" {
				c.Next()
				return
			}
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}
	}
}
