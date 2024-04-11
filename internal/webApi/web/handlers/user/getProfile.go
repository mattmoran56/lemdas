package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
	"github.com/mattmoran/fyp/api/pkg/database/models"
)

func HandleGetProfile(c *gin.Context) {
	type profileResponse struct {
		Name     string            `json:"name"`
		Avatar   string            `json:"avatar"`
		Activity []models.Activity `json:"activity"`
		Datasets []models.Dataset  `json:"datasets"`
		Bio      string            `json:"bio"`
	}

	userId := c.Param("userId")

	profile := profileResponse{}

	user, err := database.UserRepo.GetUserByID(userId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	profile.Name = user.FirstName + " " + user.LastName
	profile.Avatar = user.Avatar

	// TODO: expand to more activity types
	activity, err := database.ActivityRepo.GetActivitiesByUserIDAndType(userId, "make_public")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	profile.Activity = activity

	datasets, err := database.DatasetRepo.GetUsersPublicDatasetsIncCollaboration(userId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	profile.Datasets = datasets
	profile.Bio = user.Bio

	c.JSON(200, gin.H{"profile": profile})
	return
}
