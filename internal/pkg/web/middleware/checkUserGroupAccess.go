package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"github.com/mattmoran/fyp/api/pkg/utils"
	"net/http"
)

func CheckGroupAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		userGroupID := c.Param("groupId")
		userID := c.MustGet("userID").(string)

		group, err := database.GroupRepo.GetGroupById(userGroupID)
		if err != nil {
			if err.Error() == "record not found" {
				c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
				c.Abort()
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Group error"})
			c.Abort()
			return
		}

		if group.OwnerID == c.GetString("userID") {
			c.Next()
			return
		} else {
			membership, err := database.GroupMemberRepo.IsUserInGroup(userID, userGroupID)
			utils.HandleHandlerError(c, err)
			if membership {
				if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" {
					c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
					c.Abort()
					return
				}
				c.Next()
				return
			} else {
				c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
				c.Abort()
				return
			}
		}
	}
}
