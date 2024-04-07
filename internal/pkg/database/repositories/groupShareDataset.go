package repositories

import (
	"github.com/mattmoran/fyp/api/pkg/database/models"
	"gorm.io/gorm"
)

type GroupShareDatasetRepository struct {
	db *gorm.DB
}

func NewGroupShareDatasetRepository(db *gorm.DB) *GroupShareDatasetRepository {
	return &GroupShareDatasetRepository{db}
}

func (r *GroupShareDatasetRepository) Create(groupShareDataset models.GroupShareDataset) (*models.GroupShareDataset, error) {
	return &groupShareDataset, r.db.Create(&groupShareDataset).Error
}

func (r *GroupShareDatasetRepository) ShareDatasetWithGroup(datasetId, groupId string, access bool) (*models.GroupShareDataset, error) {
	groupShareDataset := models.GroupShareDataset{
		DatasetID:   datasetId,
		GroupID:     groupId,
		WriteAccess: access,
	}
	return r.Create(groupShareDataset)
}

func (r *GroupShareDatasetRepository) DeleteGroupShareDataset(datasetId, groupId string) error {
	return r.db.Where("dataset_id = ? AND group_id = ?", datasetId, groupId).Delete(&models.GroupShareDataset{}).Error
}

func (r *GroupShareDatasetRepository) GetGroupShareDatasetsForGroupId(groupId string) (*[]models.GroupShareDataset, error) {
	var groupShareDatasets []models.GroupShareDataset
	err := r.db.Where("group_id = ?", groupId).Find(&groupShareDatasets)
	return &groupShareDatasets, err.Error
}

func (r *GroupShareDatasetRepository) GetGroupShareDatasetsForDatasetId(datasetId string) (*[]models.GroupShareDataset, error) {
	var groupShareDatasets []models.GroupShareDataset
	err := r.db.Where("dataset_id = ?", datasetId).
		Find(&groupShareDatasets)
	for i, groupShareDataset := range groupShareDatasets {
		r.db.Model(&groupShareDataset).Association("UserGroup").Find(&groupShareDatasets[i].UserGroup)
	}
	return &groupShareDatasets, err.Error
}

func (r *GroupShareDatasetRepository) GetGroupShareDatasetForDatasetIdAndGroupId(datasetId, groupId string) (*models.GroupShareDataset, error) {
	var groupShareDataset models.GroupShareDataset
	err := r.db.Where("dataset_id = ? AND group_id = ?", datasetId, groupId).
		Order("write_access DESC").First(&groupShareDataset)
	return &groupShareDataset, err.Error
}
