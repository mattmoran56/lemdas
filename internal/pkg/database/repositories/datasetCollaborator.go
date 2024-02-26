package repositories

import (
	"github.com/mattmoran/fyp/api/pkg/database/models"
	"gorm.io/gorm"
)

type DatasetCollaboratorRepository struct {
	db *gorm.DB
}

func NewDatasetCollaboratorRepository(database *gorm.DB) *DatasetCollaboratorRepository {
	return &DatasetCollaboratorRepository{
		db: database,
	}
}

func (d *DatasetCollaboratorRepository) CreateDatabaseCollaborator(collaborator models.DatasetCollaborator) (models.DatasetCollaborator, error) {
	result := d.db.Create(&collaborator)
	return collaborator, result.Error
}

func (d *DatasetCollaboratorRepository) AddCollaborator(datasetID, userID string) (models.DatasetCollaborator, error) {
	datasetCollaborator := models.DatasetCollaborator{
		UserID:    userID,
		DatasetID: datasetID,
	}

	result := d.db.Create(&datasetCollaborator)
	return datasetCollaborator, result.Error
}

func (d *DatasetCollaboratorRepository) RemoveCollaborator(datasetID, userID string) error {
	result := d.db.Where("dataset_id = ? AND user_id = ?", datasetID, userID).Delete(&models.DatasetCollaborator{})
	return result.Error
}

func (d *DatasetCollaboratorRepository) GetCollaborators(datasetID string) ([]models.DatasetCollaborator, error) {
	var collaborators []models.DatasetCollaborator
	result := d.db.Where("dataset_id = ?", datasetID).Find(&collaborators)
	return collaborators, result.Error
}

func (d *DatasetCollaboratorRepository) GetCollaborator(datasetID, userID string) (models.DatasetCollaborator, error) {
	var collaborator models.DatasetCollaborator
	result := d.db.Where("dataset_id = ? AND user_id = ?", datasetID, userID).First(&collaborator)
	return collaborator, result.Error
}
