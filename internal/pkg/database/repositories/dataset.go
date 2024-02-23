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

func (d *DatasetRepository) CreateDataset(dataset models.Dataset) (models.Dataset, error) {
	result := d.db.Create(&dataset)
	return dataset, result.Error
}

func (d *DatasetRepository) GetUsersDatasets(userId string) ([]models.Dataset, error) {
	var datasets []models.Dataset
	result := d.db.Where("owner_id = ?", userId).Find(&datasets)
	return datasets, result.Error
}

func (d *DatasetRepository) GetUsersDatasetsOrderBy(userId string, orderBy string) ([]models.Dataset, error) {
	var datasets []models.Dataset
	result := d.db.Order(orderBy).Where("owner_id = ?", userId).Find(&datasets)
	return datasets, result.Error
}

func (d *DatasetRepository) CheckUserAccessToDataset(datasetID, userID string) (bool, error) {
	var dataset models.Dataset
	result := d.db.Where("id = ?", datasetID).First(&dataset)
	accessPermitted := false

	if dataset.OwnerID == userID {
		accessPermitted = true
	} else if dataset.IsPublic {
		accessPermitted = true
	}

	return accessPermitted, result.Error
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

func (d *DatasetRepository) SearchByName(query, userID string) ([]models.Dataset, error) {
	var datasets []models.Dataset
	result := d.db.Where("dataset_name LIKE ? AND (is_public = 1 OR owner_id = ?)", "%"+query+"%", userID).
		Find(&datasets)
	return datasets, result.Error
}
