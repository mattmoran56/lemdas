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

func (d *DatasetRepository) GetDatasets() ([]models.Dataset, error) {
	var datasets []models.Dataset
	result := d.db.Find(&datasets)
	return datasets, result.Error
}

func (d *DatasetRepository) GetDatasetByName(DatasetName string) (models.Dataset, error) {
	var dataset models.Dataset
	result := d.db.Where("dataset_name = ?", DatasetName).First(&dataset)
	return dataset, result.Error
}
