package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pepese/golang-gin-sample/app/interface/controller"
)

var ginServer *gin.Engine

func GinServerRun() {
	RunWithGracefulStop(ginServer)
}

func NewGinServer() *gin.Engine {
	e := gin.New()
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	controller.NewGinRouter(e)
	ginServer = e

	return e
}
