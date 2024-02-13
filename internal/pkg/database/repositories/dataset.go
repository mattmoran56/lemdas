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

func (d *DatasetRepository) GetDatasets(userId string) ([]models.Dataset, error) {
	var datasets []models.Dataset
	result := d.db.Where("owner_id = ?", userId).Find(&datasets)
	return datasets, result.Error
}

func (d *DatasetRepository) CreateDataset(dataset models.Dataset) (string, error) {
	result := d.db.Create(&dataset)
	return dataset.ID, result.Error
}

func (d *DatasetRepository) GetDatasetsOrderBy(userId string, orderBy string) ([]models.Dataset, error) {
	var datasets []models.Dataset
	result := d.db.Order(orderBy).Where("owner_id = ?", userId).Find(&datasets)
	return datasets, result.Error
}

func (d *DatasetRepository) GetDatasetByID(DatasetID string) (models.Dataset, error) {
	var dataset models.Dataset
	result := d.db.Where("id = ?", DatasetID).First(&dataset)
	return dataset, result.Error
}

func (d *DatasetRepository) GetDatasetByName(DatasetName string) (models.Dataset, error) {
	var dataset models.Dataset
	result := d.db.Where("dataset_name = ?", DatasetName).First(&dataset)
	return dataset, result.Error
}
