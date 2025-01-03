package controllers

import "github.com/prabhjotaulakh159/expense-tracker/services"

type UserController struct {
	SERVICE *services.UserService
}

var instance *UserController

func GetUserControllerInstance(service *services.UserService) *UserController {
	if instance == nil {
		instance = &UserController{SERVICE: service}
	}
	return instance
}
