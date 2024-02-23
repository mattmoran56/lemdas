package repositories

import (
	"github.com/mattmoran/fyp/api/pkg/database/models"
	"gorm.io/gorm"
)

type FileAttributeRepository struct {
	db *gorm.DB
}

func NewFileAttributeRepository(database *gorm.DB) *FileAttributeRepository {
	return &FileAttributeRepository{
		db: database,
	}
}

func (d *FileAttributeRepository) CreateFileAttribute(datasetAttribute models.FileAttribute) (models.FileAttribute, error) {
	result := d.db.Create(&datasetAttribute)
	return datasetAttribute, result.Error
}

func (d *FileAttributeRepository) GetFileAttributeByID(id string) (models.FileAttribute, error) {
	var datasetAttribute models.FileAttribute
	result := d.db.Where("ID = ?", id).First(&datasetAttribute)

	return datasetAttribute, result.Error
}

func (d *FileAttributeRepository) GetFileAttributeByFileID(datasetID string) ([]models.FileAttribute, error) {
	var datasetAttributes []models.FileAttribute
	result := d.db.Where("File_id = ?", datasetID).Find(&datasetAttributes)

	return datasetAttributes, result.Error
}

func (d *FileAttributeRepository) GetFileAttributeByFileIDOrderBy(datasetID string, orderBy string) ([]models.FileAttribute, error) {
	var datasetAttributes []models.FileAttribute
	result := d.db.Where("File_id = ?", datasetID).Order(orderBy).Find(&datasetAttributes)

	return datasetAttributes, result.Error
}

func (d *FileAttributeRepository) UpdateFileAttribute(fileAttribute models.FileAttribute) (models.FileAttribute, error) {
	result := d.db.Model(&models.FileAttribute{}).Where("ID = ?", fileAttribute.ID).Updates(&fileAttribute)
	return fileAttribute, result.Error
}

func (d *FileAttributeRepository) DeleteFileAttribute(id string) error {
	result := d.db.Where("ID = ?", id).Delete(&models.FileAttribute{})
	return result.Error
}
