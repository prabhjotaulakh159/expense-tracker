package controllers

import (
	"net/http"
	"github.com/prabhjotaulakh159/expense-tracker/services"
)

type UserController struct {}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController {}
}

func (userController *UserController) Register(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}