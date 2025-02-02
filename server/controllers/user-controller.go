package controllers

import (
	"net/http"
	"github.com/prabhjotaulakh159/expense-tracker/services"
	"github.com/prabhjotaulakh159/expense-tracker/exceptions"
	"github.com/prabhjotaulakh159/expense-tracker/types"
	"encoding/json"
	"errors"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController { UserService: userService }
}

func (userController *UserController) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var authRequest *types.AuthRequest
	
	if err := json.NewDecoder(r.Body).Decode(&authRequest); err != nil {
		json.NewEncoder(w).Encode(&types.ErrorResponse {Message: err.Error() })
		return
	}


	var validationError *exceptions.ValidationError
	var serverError *exceptions.ServerError
	if err := userController.UserService.Register(authRequest.Username, authRequest.Password); err != nil {
		status := -1
		if errors.As(err, &validationError) {
			status = http.StatusBadRequest
		} else if errors.As(err, &serverError) {
			status = http.StatusInternalServerError
		} 
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(&types.ErrorResponse { Message: err.Error() })
		return
	}

	w.WriteHeader(http.StatusOK)
}