package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/prabhjotaulakh159/expense-tracker/models"
)

type MockUserRepository struct {
	mock.Mock
}

func (userRepository *MockUserRepository) CreateUser(user *models.User) error {
	args := userRepository.Called(user)
	return args.Error(0)
}

func (userRepository *MockUserRepository) IsUsernameUnique(username string) (bool, error) {
	args := userRepository.Called(username)
	return args.Bool(0), args.Error(1)
}