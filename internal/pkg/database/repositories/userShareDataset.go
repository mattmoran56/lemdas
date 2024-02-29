package repositories

import (
	"github.com/mattmoran/fyp/api/pkg/database/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserShareDatasetRepository struct {
	db *gorm.DB
}

func NewUserShareDatasetRepository(db *gorm.DB) *UserShareDatasetRepository {
	return &UserShareDatasetRepository{db}
}

func (r *UserShareDatasetRepository) Create(userShareDataset models.UserShareDataset) (*models.UserShareDataset, error) {
	zap.S().Debug(userShareDataset)
	return &userShareDataset, r.db.Create(&userShareDataset).Error
}

func (r *UserShareDatasetRepository) ShareDatasetWithUser(datasetId, userId string, access bool) (*models.UserShareDataset, error) {
	userShareDataset := models.UserShareDataset{
		DatasetID:   datasetId,
		UserID:      userId,
		WriteAccess: access,
	}
	return r.Create(userShareDataset)
}

func (r *UserShareDatasetRepository) DeleteUserShareDataset(datasetId, userId string) error {
	return r.db.Where("dataset_id = ? AND user_id = ?", datasetId, userId).Delete(&models.UserShareDataset{}).Error
}

func (r *UserShareDatasetRepository) GetUserShareDatasetsForUserId(userId string) (*[]models.UserShareDataset, error) {
	var userShareDatasets []models.UserShareDataset
	err := r.db.Where("user_id = ?", userId).Find(&userShareDatasets)
	return &userShareDatasets, err.Error
}

func (r *UserShareDatasetRepository) GetUserShareDatasetsForDatasetId(datasetId string) (*[]models.UserShareDataset, error) {
	var userShareDatasets []models.UserShareDataset
	err := r.db.Where("dataset_id = ?", datasetId).
		Find(&userShareDatasets)
	for i, userShareDataset := range userShareDatasets {
		r.db.Model(&userShareDataset).Association("User").Find(&userShareDatasets[i].User)
	}
	return &userShareDatasets, err.Error
}

func (r *UserShareDatasetRepository) GetUserShareDatasetForDatasetIdAndUserId(datasetId, userId string) (*models.UserShareDataset, error) {
	var userShareDataset models.UserShareDataset
	err := r.db.Where("dataset_id = ? AND user_id = ?", datasetId, userId).First(&userShareDataset)
	return &userShareDataset, err.Error

}
