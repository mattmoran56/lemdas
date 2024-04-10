package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/database/models"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func HandleFix(c *gin.Context) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.S().Fatalf("Error connecting to database, %v", err)
		c.JSON(500, gin.H{"error": "Error connecting to database"})
		return
	}

	//db.Create(&models.FileAttributeGroup{
	//	Base: models.Base{
	//		ID:        "root",
	//		CreatedAt: 0,
	//		UpdatedAt: 0,
	//	},
	//	AttributeGroupName: "",
	//	FileID:             "",
	//	ParentGroupID:      "root",
	//	Children:           nil,
	//	Attributes:         nil,
	//})

	var files []models.File
	db.Model(&models.File{}).Find(&files)

	var nullStr *string = nil

	for _, file := range files {
		fileAttributeGroup := models.FileAttributeGroup{
			FileID:             file.ID,
			AttributeGroupName: "root",
			ParentGroupID:      nullStr,
		}
		db.Create(&fileAttributeGroup)

		var attributes []models.FileAttribute
		db.Model(&models.FileAttribute{}).Where("file_id = ?", file.ID).Find(&attributes)
		for _, attribute := range attributes {
			attribute.AttributeGroupID = fileAttributeGroup.ID
			db.Save(&attribute)
		}
	}

}
