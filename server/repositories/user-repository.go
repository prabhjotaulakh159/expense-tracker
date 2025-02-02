package repositories

import (
	"github.com/prabhjotaulakh159/expense-tracker/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user *models.User) error
	IsUsernameUnique(username string) (bool, error)
}

type UserRepository struct {
	Connection *gorm.DB
}

func NewUserRepository(connection *gorm.DB) *UserRepository {
	return &UserRepository { Connection: connection }
}

func (userRepository *UserRepository) CreateUser(user *models.User) error {
	return userRepository.Connection.Create(user).Error
}

func (userRepository *UserRepository) IsUsernameUnique(username string) (bool, error) {
	var count int64
	err := userRepository.Connection.Model(&models.User{}).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == 0, nil
}