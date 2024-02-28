package repositories

import (
	"github.com/mattmoran/fyp/api/pkg/database/models"
	"gorm.io/gorm"
)

type GroupRepository struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) *GroupRepository {
	return &GroupRepository{db}
}

func (r *GroupRepository) Create(group models.UserGroup) (models.UserGroup, error) {
	result := r.db.Create(&group)
	return group, result.Error
}

func (r *GroupRepository) GetGroupById(groupId string) (models.UserGroup, error) {
	var group models.UserGroup
	err := r.db.Where("id = ?", groupId).First(&group).Error
	return group, err
}

func (r *GroupRepository) DeleteGroup(groupId string) error {
	return r.db.Exec("DELETE FROM user_groups WHERE id = ?", groupId).Error
}
