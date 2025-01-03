package services

import (
	"strings"

	"github.com/prabhjotaulakh159/expense-tracker/myerrors"
	"github.com/prabhjotaulakh159/expense-tracker/repositories"
)

type IUserService interface {
	CreateUser(username string, password string) error
}

type UserService struct {
	REPO *repositories.UserRepo
}

type MockUserService struct {
	REPO *repositories.MockUserRepo
}

var instance *UserService

func GetUserServiceInstance(repo *repositories.UserRepo) *UserService {
	if instance == nil {
		instance = &UserService{REPO: repo}
	}
	return instance
}

func (u *UserService) CreateUser(username string, password string) error {
	return nil
}

func (m *MockUserService) CreateUser(username string, password string) error {
	_username := strings.TrimSpace(username)
	_password := strings.TrimSpace(password)
	if _username == "" || _password == "" {
		return &myerrors.ValidationError{MESSAGE: "username and/or password cannot be blank"}
	}
	if _username == _password {
		return &myerrors.ValidationError{MESSAGE: "username and password cannot be equal"}
	}
	return nil
}
