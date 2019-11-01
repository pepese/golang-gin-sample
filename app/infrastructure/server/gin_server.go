package server

import "github.com/pepese/golang-gin-sample/app/interface/controller"

type ginServer struct{}

func (hs *ginServer) Run() {
	router := controller.NewGinRouter()
	RunWithGracefulStop(router)
}

func NewGinServer() *ginServer {
	return &ginServer{}
}
