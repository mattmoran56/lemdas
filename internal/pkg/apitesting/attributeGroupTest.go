package apitesting

import (
	"github.com/mattmoran/fyp/api/pkg/database"
	"github.com/mattmoran/fyp/api/pkg/database/models"
	"os"
	"testing"
)

func AttributeGroupTest(t *testing.T, r Request) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := "fyp_test"
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	database.ConnectToDatabase(dbUsername, dbPassword, dbName, dbHost, dbPort)

	database.DatasetAttributeRepo.CreateDatasetAttribute(models.DatasetAttribute{
		Base: models.Base{
			ID:        "testattribute1",
			CreatedAt: 100,
			UpdatedAt: 100,
		},
		DatasetID:      "testdataset",
		AttributeName:  "test1",
		AttributeValue: "value1",
	})
	database.DatasetAttributeRepo.CreateDatasetAttribute(models.DatasetAttribute{
		Base: models.Base{
			ID:        "testattribute2",
			CreatedAt: 100,
			UpdatedAt: 100,
		},
		DatasetID:      "testdataset",
		AttributeName:  "test2",
		AttributeValue: "value2",
	})

	DatasetGroupTest(t, r)
}
