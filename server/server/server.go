package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var instance *http.Server

func GetServerInstance(host string, port int, router *gin.Engine) *http.Server {
	if instance == nil {
		instance = &http.Server{
			Addr:    fmt.Sprintf("%s:%d", host, port),
			Handler: router.Handler(),
		}
	}
	return instance
}
