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

func (d *DatasetAttributeRepository) CreateDatasetAttribute(datasetAttribute models.DatasetAttribute) (models.DatasetAttribute, error) {
	result := d.db.Create(&datasetAttribute)
	return datasetAttribute, result.Error
}

func (d *DatasetAttributeRepository) GetDatasetAttributeByID(id string) (models.DatasetAttribute, error) {
	var datasetAttribute models.DatasetAttribute
	result := d.db.Where("ID = ?", id).First(&datasetAttribute)

	return datasetAttribute, result.Error
}

func (d *DatasetAttributeRepository) GetDatasetAttributeByDatasetID(datasetID string) ([]models.DatasetAttribute, error) {
	var datasetAttributes []models.DatasetAttribute
	result := d.db.
		Where("dataset_id = ?", datasetID).
		Find(&datasetAttributes)

	return datasetAttributes, result.Error
}

func (d *DatasetAttributeRepository) GetAllAttributeNames() ([]string, error) {
	var attributeNames []string
	result := d.db.
		Table("dataset_attributes").
		Select("DISTINCT attribute_name").Find(&attributeNames)

	return attributeNames, result.Error
}

func (d *DatasetAttributeRepository) GetDatasetAttributeByDatasetIDOrderBy(datasetID, orderBy string) ([]models.DatasetAttribute, error) {
	var datasetAttributes []models.DatasetAttribute
	result := d.db.
		Where("dataset_id = ?", datasetID).
		Order(orderBy).Find(&datasetAttributes)

	return datasetAttributes, result.Error
}

func (d *DatasetAttributeRepository) UpdateDatasetAttribute(datasetAttribute models.DatasetAttribute) (models.DatasetAttribute, error) {
	result := d.db.Model(&models.DatasetAttribute{}).Where("ID = ?", datasetAttribute.ID).Updates(&datasetAttribute)
	return datasetAttribute, result.Error
}

func (d *DatasetAttributeRepository) DeleteDatasetAttribute(id string) error {
	result := d.db.Where("ID = ?", id).Delete(&models.DatasetAttribute{})
	return result.Error
}

func (d *DatasetAttributeRepository) DeleteDatasetAttributeByDatasetID(datasetID string) error {
	result := d.db.Where("dataset_id = ?", datasetID).Delete(&models.DatasetAttribute{})
	return result.Error
}
