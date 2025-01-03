package test

import (
	"testing"

	"github.com/prabhjotaulakh159/expense-tracker/models"
	"github.com/prabhjotaulakh159/expense-tracker/repositories"
)

func Test_CheckIfUserExistsByUsername_ThrowsError(test *testing.T) {
	repo := repositories.MockUserRepo{
		THROWS_ERROR:          true,
		IS_DUPLICATE_USERNAME: false,
	}
	_, err := repo.CheckIfUserExistsByUsername("test")
	if err == nil {
		test.Error("Expected function to return an error")
	}
}

func Test_CheckIfUserExistsByUsername_ReturnsTrue(test *testing.T) {
	repo := repositories.MockUserRepo{
		THROWS_ERROR:          false,
		IS_DUPLICATE_USERNAME: true,
	}
	exists, err := repo.CheckIfUserExistsByUsername("test")
	if err != nil {
		test.Error("Function returned an error")
	}
	if !exists {
		test.Error("Expected true")
	}
}

func Test_CheckIfUserExistsByUsername_ReturnsFalse(test *testing.T) {
	repo := repositories.MockUserRepo{
		THROWS_ERROR:          false,
		IS_DUPLICATE_USERNAME: false,
	}
	exists, err := repo.CheckIfUserExistsByUsername("test")
	if err != nil {
		test.Error("Function returned an error")
	}
	if exists {
		test.Error("Expected false")
	}
}

func Test_AddNewUser_ThrowsError(test *testing.T) {
	repo := repositories.MockUserRepo{THROWS_ERROR: true}
	user := &models.User{USERNAME: "test", PASSWORD: "john"}
	err := repo.AddNewUser(user)
	if err == nil {
		test.Error("Expected function to return an error")
	}
}

func Test_AddNewUser_Works(test *testing.T) {
	repo := repositories.MockUserRepo{THROWS_ERROR: false}
	user := &models.User{USERNAME: "test", PASSWORD: "john"}
	err := repo.AddNewUser(user)
	if err != nil {
		test.Error("Function returned an error")
	}
}

func Test_DeleteUser_ThrowsError(test *testing.T) {
	repo := repositories.MockUserRepo{THROWS_ERROR: true}
	err := repo.DeleteUser(1)
	if err == nil {
		test.Error("Expected Function to return an error")
	}
}

func Test_DeleteUser_Works(test *testing.T) {
	repo := repositories.MockUserRepo{THROWS_ERROR: false}
	err := repo.DeleteUser(1)
	if err != nil {
		test.Error("functon returned an error")
	}
}
