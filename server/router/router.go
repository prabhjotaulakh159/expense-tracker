package router

import "github.com/gin-gonic/gin"

var instance *gin.Engine

func GetRouterInstance() *gin.Engine {
	if instance == nil {
		instance = gin.Default()
	}
	return instance
}
