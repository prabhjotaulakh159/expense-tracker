package test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/prabhjotaulakh159/expense-tracker/mocks"
	"github.com/prabhjotaulakh159/expense-tracker/exceptions"
	"github.com/prabhjotaulakh159/expense-tracker/services"
	"github.com/prabhjotaulakh159/expense-tracker/models"
	"errors"
)

func Test_Register_EmptyUsername_ReturnsError(t *testing.T) {
	mockUserRepo := &mocks.MockUserRepository {}
	mockPasswordEncoder := &mocks.MockPasswordEncoder {}
	userService := &services.UserService{
		UserRepository: mockUserRepo,
		PasswordEncoder: mockPasswordEncoder,
	}
	err := userService.Register("", "validPassword123")
	assert.NotNil(t, err)
	assert.IsType(t, &exceptions.ValidationError{}, err)
	assert.Equal(t, "username cannot be empty", err.Error())
}

func Test_Register_ShortPassword_ReturnsError(t *testing.T) {
	mockUserRepo := &mocks.MockUserRepository{}
	mockPasswordEncoder := &mocks.MockPasswordEncoder{}
	userService := &services.UserService{
		UserRepository:  mockUserRepo,
		PasswordEncoder: mockPasswordEncoder,
	}
	err := userService.Register("valid", "va")
	assert.NotNil(t, err)
	assert.IsType(t, &exceptions.ValidationError{}, err)
	assert.Equal(t, "password cannot be smaller than 8 characters", err.Error())
}

func Test_Register_UsernameAndPasswordSame_ReturnsError(t *testing.T) {
	mockUserRepo := &mocks.MockUserRepository{}
	mockPasswordEncoder := &mocks.MockPasswordEncoder{}
	userService := &services.UserService{
		UserRepository:  mockUserRepo,
		PasswordEncoder: mockPasswordEncoder,
	}
	err := userService.Register("sameValue", "sameValue")
	assert.NotNil(t, err)
	assert.IsType(t, &exceptions.ValidationError{}, err)
	assert.Equal(t, "username and password cannot be the same", err.Error())
}

func Test_Register_IsUsernameUniqueReturnsError_ReturnsServerError(t *testing.T) {
	mockUserRepo := &mocks.MockUserRepository{}
	mockPasswordEncoder := &mocks.MockPasswordEncoder{}
	userService := &services.UserService{
		UserRepository:  mockUserRepo,
		PasswordEncoder: mockPasswordEncoder,
	}
	mockUserRepo.On("IsUsernameUnique", "testUser").Return(false, errors.New("database connection error"))
	err := userService.Register("testUser", "validPassword123")
	assert.NotNil(t, err)
	assert.IsType(t, &exceptions.ServerError{}, err)
	assert.Equal(t, "database connection error", err.Error())
	mockUserRepo.AssertCalled(t, "IsUsernameUnique", "testUser")
}

func Test_Register_UsernameNotUnique_ReturnsValidationError(t *testing.T) {
	mockUserRepo := &mocks.MockUserRepository{}
	mockPasswordEncoder := &mocks.MockPasswordEncoder{}
	userService := &services.UserService{
		UserRepository:  mockUserRepo,
		PasswordEncoder: mockPasswordEncoder,
	}
	mockUserRepo.On("IsUsernameUnique", "existingUser").Return(false, nil)
	err := userService.Register("existingUser", "validPassword123")
	assert.NotNil(t, err)
	assert.IsType(t, &exceptions.ValidationError{}, err)
	assert.Equal(t, "username is already taken", err.Error())
	mockUserRepo.AssertCalled(t, "IsUsernameUnique", "existingUser")
}

func Test_Register_HashPasswordError_ReturnsServerError(t *testing.T) {
	mockUserRepo := &mocks.MockUserRepository{}
	mockPasswordEncoder := &mocks.MockPasswordEncoder{}
	userService := &services.UserService{
		UserRepository:  mockUserRepo,
		PasswordEncoder: mockPasswordEncoder,
	}
	username := "newUser"
	password := "validPassword123"
	mockUserRepo.On("IsUsernameUnique", username).Return(true, nil)
	mockPasswordEncoder.On("HashPassword", password).Return("", errors.New("hashing failed"))
	err := userService.Register(username, password)
	assert.NotNil(t, err)
	assert.IsType(t, &exceptions.ServerError{}, err)
	assert.Equal(t, "hashing failed", err.Error())
	mockUserRepo.AssertCalled(t, "IsUsernameUnique", username)
	mockPasswordEncoder.AssertCalled(t, "HashPassword", password)
}

func Test_Register_CreateUserError_ReturnsServerError(t *testing.T) {
	mockUserRepo := &mocks.MockUserRepository{}
	mockPasswordEncoder := &mocks.MockPasswordEncoder{}
	userService := &services.UserService{
		UserRepository:  mockUserRepo,
		PasswordEncoder: mockPasswordEncoder,
	}
	username := "newUser"
	password := "validPassword123"
	user := &models.User{
		Username: username,
		Password: password,
	}
	mockUserRepo.On("IsUsernameUnique", username).Return(true, nil)
	mockPasswordEncoder.On("HashPassword", password).Return("validPassword123", nil)
	mockUserRepo.On("CreateUser", user).Return(errors.New("database connection error"))
	err := userService.Register(username, password)
	assert.NotNil(t, err)
	assert.IsType(t, &exceptions.ServerError{}, err)
	assert.Equal(t, "database connection error", err.Error())
	mockUserRepo.AssertCalled(t, "IsUsernameUnique", username)
	mockPasswordEncoder.AssertCalled(t, "HashPassword", password)
	mockUserRepo.AssertCalled(t, "CreateUser", user)
}

func Test_Register_Success(t *testing.T) {
    mockUserRepo := &mocks.MockUserRepository{}
    mockPasswordEncoder := &mocks.MockPasswordEncoder{}
    userService := &services.UserService{
        UserRepository:  mockUserRepo,
        PasswordEncoder: mockPasswordEncoder,
    }
    username := "newUser"
    password := "validPassword123"
    user := &models.User{
        Username: username,
        Password: password,
    }
    mockUserRepo.On("IsUsernameUnique", username).Return(true, nil)
    mockPasswordEncoder.On("HashPassword", password).Return("validPassword123", nil)
    mockUserRepo.On("CreateUser", user).Return(nil)
    err := userService.Register(username, password)
    assert.Nil(t, err)
    mockUserRepo.AssertCalled(t, "IsUsernameUnique", username)
    mockPasswordEncoder.AssertCalled(t, "HashPassword", password)
    mockUserRepo.AssertCalled(t, "CreateUser", user)
}
