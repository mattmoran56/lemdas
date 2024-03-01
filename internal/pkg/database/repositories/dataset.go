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

func (d *DatasetRepository) GetStaredDatasets(userID string) ([]models.Dataset, error) {
	var datasets []models.Dataset
	result := d.db.Select("datasets.*").
		Joins("RIGHT OUTER JOIN stared_datasets ON stared_datasets.dataset_id = datasets.id").
		Where("stared_datasets.user_id = ? AND stared_datasets.dataset_id IS NOT NULL", userID).Find(&datasets)
	return datasets, result.Error
}

func (d *DatasetRepository) GetUsersDatasetsOrderBy(userId string, orderBy string) ([]models.Dataset, error) {
	var datasets []models.Dataset
	result := d.db.Order(orderBy).Where("owner_id = ?", userId).Find(&datasets)
	return datasets, result.Error
}

func (d *DatasetRepository) GetUsersSharedDatasets(userID string) ([]models.Dataset, error) {
	var datasets []models.Dataset
	result := d.db.Select("datasets.*").
		Joins("RIGHT OUTER JOIN user_share_datasets ON user_share_datasets.dataset_id = datasets.id").
		Where("user_share_datasets.user_id = ? AND user_share_datasets.dataset_id IS NOT NULL", userID).Find(&datasets)

	if result.Error != nil {
		return nil, result.Error
	}

	result = d.db.Select("datasets.*").
		Joins("RIGHT OUTER JOIN group_share_datasets ON group_share_datasets.dataset_id = datasets.id").
		Joins("RIGHT OUTER JOIN group_members ON group_members.group_id = group_share_datasets.group_id").
		Where("group_members.user_id = ? AND group_share_datasets.dataset_id IS NOT NULL", userID).Find(&datasets)
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

func (d *DatasetRepository) DeleteDatasetByID(DatasetID string) error {
	result := d.db.Where("id = ?", DatasetID).Delete(&models.Dataset{})
	return result.Error
}

func (d *DatasetRepository) UpdateDataset(dataset models.Dataset) (models.Dataset, error) {
	updates := map[string]interface{}{}
	if dataset.DatasetName != "" {
		updates["dataset_name"] = dataset.DatasetName
	}
	if dataset.IsPublic {
		updates["is_public"] = 1
	} else {
		updates["is_public"] = 0
	}
	if dataset.OwnerID != "" {
		updates["owner_id"] = dataset.OwnerID
	}

	result := d.db.Model(&dataset).Where("id = ?", dataset.ID).Updates(updates)
	return dataset, result.Error
}
