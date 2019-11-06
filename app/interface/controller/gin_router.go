package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pepese/golang-gin-sample/app/usecase"
)

func NewGinRouter(e *gin.Engine) {
	router := e.Group("/api/v1")

	ping := router.Group("/ping")
	pingUc := usecase.NewPingUsecase()
	ping.GET("", func(ctx *gin.Context) {
		pingUc.Say("pong")
		ctx.JSON(http.StatusOK, *pingUc)
	})

	user := router.Group("/users")
	userUc := usecase.NewUserUsecase()
	user.GET("", func(ctx *gin.Context) {
		results := userUc.GetAll()
		ctx.JSON(http.StatusOK, results)
	})
}
