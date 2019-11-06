package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pepese/golang-gin-sample/app/interface/controller"
)

type ginServer struct{}

func (hs *ginServer) Run() {
	e := gin.New()
	e.Use(gin.Logger())
	e.Use(gin.Recovery())

	controller.NewGinRouter(e)
	RunWithGracefulStop(e)
}

func NewGinServer() *ginServer {
	return &ginServer{}
}
