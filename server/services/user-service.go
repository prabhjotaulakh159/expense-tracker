package services

import (
	"github.com/prabhjotaulakh159/expense-tracker/repositories"
	"github.com/prabhjotaulakh159/expense-tracker/exceptions"
	"github.com/prabhjotaulakh159/expense-tracker/models"
	"github.com/prabhjotaulakh159/expense-tracker/encryption"
	"strings"
	"fmt"
)

type UserService struct {
	UserRepository repositories.IUserRepository
	PasswordEncoder encryption.IPasswordEncoder
}

func NewUserService(userRepository repositories.IUserRepository, passwordEncoder encryption.IPasswordEncoder ) *UserService {
	return &UserService { UserRepository: userRepository, PasswordEncoder: passwordEncoder }
}

func (userService *UserService) Register(username string, password string) error {
	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)
	
	minPasswordLength := 8
	
	if len(username) == 0  {
		return exceptions.NewValidationError("username cannot be empty")
	}
	if len(password) < minPasswordLength {
		return exceptions.NewValidationError(fmt.Sprintf("password cannot be smaller than %d characters", minPasswordLength))
	}
	if username == password {
		return exceptions.NewValidationError("username and password cannot be the same")
	}
	
	usernameIsUnique, err := userService.UserRepository.IsUsernameUnique(username)
	if err != nil {
		return exceptions.NewServerError(err.Error())
	}
	if !usernameIsUnique {
		return exceptions.NewValidationError("username is already taken")
	}

	hash, err := userService.PasswordEncoder.HashPassword(password)
	if err != nil {
		return exceptions.NewServerError(err.Error())
	}

	user := models.NewUser(username, hash)
	
	if err := userService.UserRepository.CreateUser(user); err != nil {
		return exceptions.NewServerError(err.Error())
	}

	return nil
}