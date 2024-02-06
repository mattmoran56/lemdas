package repositories

import (
	"github.com/mattmoran/fyp/api/pkg/database/models"
	"gorm.io/gorm"
)

type DatasetRepository struct {
	db *gorm.DB
}

func NewDatasetRepository(database *gorm.DB) *DatasetRepository {
	return &DatasetRepository{
		db: database,
	}
}

func (d *DatasetRepository) CreateDataset(dataset models.Dataset) error {
	result := d.db.Create(&dataset)
	return result.Error
}
