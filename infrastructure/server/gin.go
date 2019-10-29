package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pepese/golang-gin-sample/interface"
)

type Server struct {}

func (*Server) Run() {

	api := gin.New()
	interface.Route(api)
	api.Run()
}