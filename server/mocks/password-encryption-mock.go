package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockPasswordEncoder struct {
	mock.Mock
}	

func (passwordEncoder *MockPasswordEncoder) HashPassword(password string) (string, error) {
	args := passwordEncoder.Called(password)
	return args.String(0), args.Error(1)
}