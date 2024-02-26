package repositories

import (
	"github.com/mattmoran/fyp/api/pkg/database/models"
	"gorm.io/gorm"
)

type StaredDatasetRepository struct {
	db *gorm.DB
}

func NewStaredDatasetRepository(database *gorm.DB) *StaredDatasetRepository {
	return &StaredDatasetRepository{
		db: database,
	}
}

func (s *StaredDatasetRepository) GetStaredDataset(userID, datasetID string) (bool, error) {
	var staredDataset []models.StaredDataset
	result := s.db.Where("user_id = ? AND dataset_id = ?", userID, datasetID).Find(&staredDataset)
	return len(staredDataset) > 0, result.Error
}

func (s *StaredDatasetRepository) StarDataset(userID, datasetID string) (models.StaredDataset, error) {
	staredDataset := models.StaredDataset{
		UserID:    userID,
		DatasetID: datasetID,
	}
	result := s.db.Create(&staredDataset)
	return staredDataset, result.Error
}

func (s *StaredDatasetRepository) UnstarDataset(userID, datasetID string) error {
	result := s.db.Where("user_id = ? AND dataset_id = ?", userID, datasetID).Delete(&models.StaredDataset{})
	return result.Error
}
