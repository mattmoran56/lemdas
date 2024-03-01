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

func (r *GroupRepository) SearchForGroupUserIsIn(query, userID string) ([]models.UserGroup, error) {
	var groups []models.UserGroup
	err := r.db.Select("user_groups.*").
		Joins("LEFT JOIN group_members ON group_members.group_id = user_groups.id").
		Where("user_groups.group_name LIKE ? AND (user_groups.owner_id = ? OR group_members.user_id = ?)", "%"+query+"%", userID, userID).Find(&groups).Error

	return groups, err
}
