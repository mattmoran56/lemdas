package repositories

import (
	"github.com/mattmoran/fyp/api/pkg/database/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(database *gorm.DB) *UserRepo {
	return &UserRepo{
		db: database,
	}
}

func (u *UserRepo) CreateUser(user models.User) error {
	result := u.db.Create(&user)
	return result.Error
}

func (u *UserRepo) CheckUserExistsByEmail(email string) (bool, error) {
	var user models.User
	result := u.db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, result.Error
	}

	return true, nil
}

func (u *UserRepo) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	result := u.db.Where("email = ?", email).First(&user)

	return user, result.Error
}

func (u *UserRepo) GetUserByID(UserID string) (models.User, error) {
	var user models.User
	result := u.db.Where("ID = ?", UserID).First(&user)

	return user, result.Error
}

func (u *UserRepo) SearchForUser(query string) ([]models.User, error) {
	var users []models.User
	result := u.db.Where("first_name LIKE ? OR last_name LIKE ? OR email LIKE ? ", "%"+query+"%", "%"+query+"%", "%"+query+"%").Find(&users)

	return users, result.Error

}
