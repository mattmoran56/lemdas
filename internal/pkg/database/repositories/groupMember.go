package repositories

import (
	"github.com/mattmoran/fyp/api/pkg/database/models"
	"gorm.io/gorm"
)

type GroupMemberRepository struct {
	db *gorm.DB
}

func NewGroupMemberRepository(db *gorm.DB) *GroupMemberRepository {
	return &GroupMemberRepository{db}
}

func (r *GroupMemberRepository) CreateGroupMember(groupMember models.GroupMember) error {
	return r.db.Create(&groupMember).Error
}

func (r *GroupMemberRepository) AddUserToGroup(userId string, groupId string) error {
	result := r.db.Create(&models.GroupMember{UserID: userId, GroupID: groupId})
	return result.Error
}

func (r *GroupMemberRepository) RemoveUserFromGroup(userId string, groupId string) error {
	return r.db.Exec("DELETE FROM group_members WHERE user_id = ? AND group_id = ?", userId, groupId).Error
}

func (r *GroupMemberRepository) GetGroupMembers(groupId string) ([]models.User, error) {
	var users []models.User
	err := r.db.Select("users.*").
		Joins("JOIN group_members ON users.id = group_members.user_id").
		Where("group_members.group_id = ?", groupId).
		Find(&users)
	return users, err.Error
}

func (r *GroupMemberRepository) GetGroupsForUser(userId string) ([]models.UserGroup, error) {
	var groups []models.UserGroup
	err := r.db.Select("user_groups.*").
		Joins("LEFT JOIN group_members ON user_groups.id = group_members.group_id").
		Where("group_members.user_id = ? OR user_groups.owner_id = ?", userId, userId).
		Find(&groups)
	return groups, err.Error
}

func (r *GroupMemberRepository) IsUserInGroup(userId string, groupId string) (bool, error) {
	var count int64
	err := r.db.Model(&models.GroupMember{}).Where("user_id = ? AND group_id = ?", userId, groupId).Count(&count)
	return count > 0, err.Error
}

func (r *GroupMemberRepository) DeleteGroupMembers(groupId string) error {
	return r.db.Exec("DELETE FROM group_members WHERE group_id = ?", groupId).Error
}
