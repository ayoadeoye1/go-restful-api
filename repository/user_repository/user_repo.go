package userrepository

import (
	"errors"

	"github.com/ayoadeoye1/insta-shop-screening/models"
	"gorm.io/gorm"
)

type UserRepoImpl struct {
	Db *gorm.DB
}

func NewUserRepoImpl(Db *gorm.DB) UserRepo {
	return &UserRepoImpl{Db: Db}
}

func (u *UserRepoImpl) Add(users models.Users) error {
	result := u.Db.Create(&users)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserRepoImpl) Edit(users models.Users) error {
	result := u.Db.Save(&users)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserRepoImpl) FindAll() ([]models.Users, error) {
	var users []models.Users
	result := u.Db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (u *UserRepoImpl) FindById(userId int) (*models.Users, error) {
	var user models.Users
	result := u.Db.First(&user, userId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}
	return &user, nil
}

func (u *UserRepoImpl) FindByEmail(userEmail string) (*models.Users, error) {
	var user models.Users
	result := u.Db.Where("email = ?", userEmail).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return &models.Users{}, nil
		}
		return &models.Users{}, result.Error
	}
	return &user, nil
}

func (u *UserRepoImpl) Remove(userId int) error {
	result := u.Db.Delete(&models.Users{}, userId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
