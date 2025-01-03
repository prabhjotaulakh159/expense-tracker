package router

import (
	"github.com/gin-gonic/gin"
	"github.com/prabhjotaulakh159/expense-tracker/controllers"
	"github.com/prabhjotaulakh159/expense-tracker/repositories"
	"github.com/prabhjotaulakh159/expense-tracker/services"
	"gorm.io/gorm"
)

var instance *gin.Engine

func GetRouterInstance(gormDb *gorm.DB) *gin.Engine {
	if instance == nil {
		userRepo := repositories.GetUserRepoInstance(gormDb)
		userService := services.GetUserServiceInstance(userRepo)
		_ = controllers.GetUserControllerInstance(userService)

		instance = gin.Default()
	}
	return instance
}
