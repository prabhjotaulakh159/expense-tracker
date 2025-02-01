package services

import (
	"github.com/prabhjotaulakh159/expense-tracker/repositories"
)

type UserService struct {
	UserRepository repositories.IUserRepository
}

func NewUserService(userRepository repositories.IUserRepository) *UserService {
	return &UserService { UserRepository: userRepository }
}

func (userService *UserService) Register(username string, password string) error {
	return nil
}