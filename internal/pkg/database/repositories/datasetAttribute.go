package repositories

import (
	"github.com/mattmoran/fyp/api/pkg/database/models"
	"gorm.io/gorm"
)

type DatasetAttributeRepository struct {
	db *gorm.DB
}

func NewDatasetAttributeRepository(database *gorm.DB) *DatasetAttributeRepository {
	return &DatasetAttributeRepository{
		db: database,
	}
}

func (d *DatasetAttributeRepository) CreateDatasetAttribute(datasetAttribute models.DatasetAttribute) (string, error) {
	result := d.db.Create(&datasetAttribute)
	return datasetAttribute.ID, result.Error
}

func (d *DatasetAttributeRepository) GetDatasetAttributeByID(id string) (models.DatasetAttribute, error) {
	var datasetAttribute models.DatasetAttribute
	result := d.db.Where("ID = ?", id).First(&datasetAttribute)

	return datasetAttribute, result.Error
}

func (d *DatasetAttributeRepository) GetDatasetAttributeByDatasetID(datasetID string) ([]models.DatasetAttribute, error) {
	var datasetAttributes []models.DatasetAttribute
	result := d.db.Where("Dataset_id = ?", datasetID).Find(&datasetAttributes)

	return datasetAttributes, result.Error
}

func (d *DatasetAttributeRepository) GetDatasetAttributeByDatasetIDOrderBy(datasetID string, orderBy string) ([]models.DatasetAttribute, error) {
	var datasetAttributes []models.DatasetAttribute
	result := d.db.Where("Dataset_id = ?", datasetID).Order(orderBy).Find(&datasetAttributes)

	return datasetAttributes, result.Error
}
