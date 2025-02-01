package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}

func NewUser(username string, password string) *User {
	return &User { Username: username, Password: password }
}