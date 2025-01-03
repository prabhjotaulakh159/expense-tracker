package services

import "github.com/prabhjotaulakh159/expense-tracker/repositories"

type UserService struct {
	REPO *repositories.UserRepo
}

var instance *UserService

func GetUserServiceInstance(repo *repositories.UserRepo) *UserService {
	if instance == nil {
		instance = &UserService{REPO: repo}
	}
	return instance
}
