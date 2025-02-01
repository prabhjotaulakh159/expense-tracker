package controllers

import (
	"net/http"
)

type UserController struct {}

func NewUserController() *UserController {
	return &UserController {}
}

func (userController *UserController) Register(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}