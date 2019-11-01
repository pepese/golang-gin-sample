package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pepese/golang-gin-sample/app/interface/controller"
)

type ginServer struct{}

func (hs *ginServer) Run() {
	e := gin.Default()
	controller.NewGinRouter(e)
	RunWithGracefulStop(e)
}

func NewGinServer() *ginServer {
	return &ginServer{}
}
