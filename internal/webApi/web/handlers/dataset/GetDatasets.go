package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database"
)

func HandleGetDatasets(c *gin.Context) {
	orderBy, found := c.GetQuery("orderBy")

	userId := c.MustGet("userID").(string)

	if !found {
		datasets, err := database.DatasetRepo.GetUsersDatasets(userId)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"datasets": datasets})
	} else {
		datasets, err := database.DatasetRepo.GetUsersDatasetsOrderBy(userId, orderBy)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"datasets": datasets})
	}

	return
}
