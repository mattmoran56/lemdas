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

func (f *FileRepository) CreateFile(file models.File) (string, error) {
	result := f.db.Create(&file)
	return file.ID, result.Error
}

func (f *FileRepository) GetFileByID(id string) (models.File, error) {
	var file models.File
	result := f.db.Where("id = ?", id).First(&file)
	return file, result.Error
}

func (f *FileRepository) GetFilesForDataset(datasetID string) ([]models.File, error) {
	var files []models.File
	result := f.db.Where("dataset_id = ?", datasetID).Find(&files)
	return files, result.Error
}

func (f *FileRepository) SearchByName(query, userID string) ([]models.File, error) {
	var files []models.File
	result := f.db.Where("name LIKE ? AND status = ? AND (owner_id = ?)",
		"%"+query+"%", "processed", userID).Find(&files)
	return files, result.Error
}

func (f *FileRepository) DeleteFile(id string) error {
	result := f.db.Where("id = ?", id).Delete(&models.File{})
	return result.Error
}

func (f *FileRepository) UpdateFile(file models.File) (models.File, error) {
	updates := map[string]interface{}{}
	if file.Name != "" {
		updates["name"] = file.Name
	}
	if file.OwnerID != "" {
		updates["owner_id"] = file.OwnerID
	}
	if file.Status != "" {
		updates["status"] = file.Status
	}
	if file.DatasetID != "" {
		updates["dataset_id"] = file.DatasetID
	}

	result := f.db.Model(&file).Where("id = ? ", file.ID).Updates(updates)
	return file, result.Error
}
