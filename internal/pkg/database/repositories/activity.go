package repositories

import (
	"github.com/mattmoran/fyp/api/pkg/database/models"
	"gorm.io/gorm"
)

type ActivityRepository struct {
	db *gorm.DB
}

func NewActivityRepository(database *gorm.DB) *ActivityRepository {
	return &ActivityRepository{
		db: database,
	}
}

func (d *ActivityRepository) CreateActivity(activity models.Activity) (models.Activity, error) {
	result := d.db.Create(&activity)
	return activity, result.Error
}

func (d *ActivityRepository) GetActivitiesByUserIDAndType(userId string, activityType string) ([]models.Activity, error) {
	var activities []models.Activity
	result := d.db.
		Where("user_id = ? AND type = ?", userId, activityType).
		Order("created_at desc").
		Find(&activities)
	return activities, result.Error
}
