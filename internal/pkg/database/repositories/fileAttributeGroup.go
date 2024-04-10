package repositories

import (
	"github.com/mattmoran/fyp/api/pkg/database/models"
	"gorm.io/gorm"
)

type FileAttributeGroupRepository struct {
	db *gorm.DB
}

func NewFileAttributeGroupRepository(database *gorm.DB) *FileAttributeGroupRepository {
	return &FileAttributeGroupRepository{
		db: database,
	}
}

func (d *FileAttributeGroupRepository) CreateFileAttributeGroup(fileAttributeGroup models.FileAttributeGroup) (models.FileAttributeGroup, error) {
	result := d.db.Create(&fileAttributeGroup)
	return fileAttributeGroup, result.Error
}

func (d *FileAttributeGroupRepository) GetFileAttributeGroupByFileID(id string) ([]models.FileAttributeGroup, error) {
	var fileAttributeGroups []models.FileAttributeGroup
	result := d.db.Model(&models.FileAttributeGroup{}).Where("file_id = ? AND attribute_group_name = 'root'", id).Preload("Children", preloadChildren).Find(&fileAttributeGroups)
	return fileAttributeGroups, result.Error
}

func (d *FileAttributeGroupRepository) GetFileAttributeGroupByFileIDOrderBy(id, orderBy string) ([]models.FileAttributeGroup, error) {
	var fileAttributeGroups []models.FileAttributeGroup
	result := d.db.Model(&models.FileAttributeGroup{}).Where("file_id = ? AND attribute_group_name = 'root'", id).Preload("Children", preloadChildren).Order(orderBy).Find(&fileAttributeGroups)
	return fileAttributeGroups, result.Error
}

func preloadChildren(db *gorm.DB) *gorm.DB {
	return db.Preload("Children", preloadChildren).Preload("Attributes")
}
