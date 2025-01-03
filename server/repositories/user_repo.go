package repositories

import "gorm.io/gorm"

type UserRepo struct {
	GORM *gorm.DB
}

var instance *UserRepo

func GetUserRepoInstance(gormDb *gorm.DB) *UserRepo {
	if instance == nil {
		instance = &UserRepo{GORM: gormDb}
	}
	return instance
}
