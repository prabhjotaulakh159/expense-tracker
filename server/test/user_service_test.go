package test

import (
	"errors"
	"testing"

	"github.com/prabhjotaulakh159/expense-tracker/myerrors"
	"github.com/prabhjotaulakh159/expense-tracker/repositories"
	"github.com/prabhjotaulakh159/expense-tracker/services"
)

func Test_CreateUser_BlankUsernamePassword_ReturnsError(test *testing.T) {
	repo := repositories.MockUserRepo{THROWS_ERROR: false, IS_DUPLICATE_USERNAME: true}
	service := services.MockUserService{REPO: &repo}
	err := service.CreateUser(" ", "  ")
	var validationErr *myerrors.ValidationError
	if !errors.As(err, &validationErr) {
		test.Error("Expected error to be of type ValidationError")
	}
	if err.Error() != "username and/or password cannot be blank" {
		test.Error("Expected error message to be: username and/or password cannot be blank")
	}
}

func Test_CreateUser_SameUsernamePassword_ReturnsError(test *testing.T) {
	repo := repositories.MockUserRepo{THROWS_ERROR: false, IS_DUPLICATE_USERNAME: true}
	service := services.MockUserService{REPO: &repo}
	err := service.CreateUser("samesame", "samesame")
	var validationErr *myerrors.ValidationError
	if !errors.As(err, &validationErr) {
		test.Error("Expected error to be of type ValidationError")
	}
	if err.Error() != "username and password cannot be equal" {
		test.Error("Expected error message to be: username and password cannot be equal")
	}
}
