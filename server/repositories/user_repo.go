package repositories

import (
	"errors"

	"github.com/prabhjotaulakh159/expense-tracker/models"
	"gorm.io/gorm"
)

type IUserRepo interface {
	CheckIfUserExistsByUsername() (bool, error)
	AddNewUser(user *models.User) error
}

type UserRepo struct {
	GORM *gorm.DB
}

type MockUserRepo struct {
	IS_DUPLICATE_USERNAME bool
	THROWS_ERROR          bool
}

var instance *UserRepo

func GetUserRepoInstance(gormDb *gorm.DB) *UserRepo {
	if instance == nil {
		instance = &UserRepo{GORM: gormDb}
	}
	return instance
}

func (u *UserRepo) CheckIfUserExistsByUsername(username string) (bool, error) {
	var exists bool
	err := u.GORM.Model(&models.User{}).Select("1").Where("username = ?", username).Scan(&exists).Error
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (u *UserRepo) AddNewUser(user *models.User) error {
	res := u.GORM.Create(user)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (m *MockUserRepo) CheckIfUserExistsByUsername(username string) (bool, error) {
	if m.THROWS_ERROR {
		return false, errors.New("Error checking if username is unique")
	}
	if m.IS_DUPLICATE_USERNAME {
		return true, nil
	}
	return false, nil
}

func (m *MockUserRepo) AddNewUser(user *models.User) error {
	if m.THROWS_ERROR {
		return errors.New("Error creating user")
	}
	return nil
}
