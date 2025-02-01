package encryption

import (
	"golang.org/x/crypto/bcrypt"
)

type IPasswordEncoder interface {
	HashPassword(password string) (string, error)
}

type PasswordEncoder struct {}

func NewPasswordEncoder() *PasswordEncoder {
	return &PasswordEncoder {}
}

func (passwordEncoder *PasswordEncoder) HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}