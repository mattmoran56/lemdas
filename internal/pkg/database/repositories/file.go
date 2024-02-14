package repositories

import (
	"github.com/mattmoran/fyp/api/pkg/database/models"
	"gorm.io/gorm"
)

type FileRepository struct {
	db *gorm.DB
}

func NewFileRepository(database *gorm.DB) *FileRepository {
	return &FileRepository{
		db: database,
	}
}

func (f *FileRepository) CreateFile(file models.File) error {
	result := f.db.Create(&file)
	return result.Error
}

func (f *FileRepository) GetFileByID(id string) (models.File, error) {
	var file models.File
	result := f.db.Where("id = ?", id).First(&file)
	return file, result.Error
}

func (f *FileRepository) GetFilesForDataset(datasetId string) ([]models.File, error) {
	var files []models.File
	result := f.db.Where("dataset_id = ?", datasetId).Find(&files)
	return files, result.Error
}
