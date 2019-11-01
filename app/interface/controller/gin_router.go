package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pepese/golang-gin-sample/app/usecase"
)

func NewGinRouter() *gin.Engine {
	e := gin.Default()
	router := e.Group("")

	ping := router.Group("/ping")
	pingUc := usecase.NewPingUsecase()
	ping.GET("", pingUc.GET)

	return e
}
