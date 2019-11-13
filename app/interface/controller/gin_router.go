package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pepese/golang-gin-sample/app/domain/model"
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
	userUc := usecase.UserUsecase{}
	{
		userUc.Init()
		// GET /api/v1/users
		user.GET("", func(c *gin.Context) {
			m := model.User{}
			c.Bind(&m)
			result := userUc.List(&m)
			c.JSON(http.StatusOK, result)
		})
		// GET /api/v1/users/:user_id
		user.GET("/:user_id", func(c *gin.Context) {
			m := model.User{}
			c.Bind(&m)
			id, _ := strconv.Atoi(c.Params.ByName("user_id"))
			m.ID = id
			result := userUc.Get(&m)
			c.JSON(http.StatusOK, result)
		})
		// POST /api/v1/users
		user.POST("", func(c *gin.Context) {
			m := model.User{}
			c.Bind(&m)
			result := userUc.Create(&m)
			c.JSON(http.StatusOK, result)
		})
		// PUT /api/v1/users/:user_id
		user.PUT("/:user_id", func(c *gin.Context) {
			m := model.User{}
			c.Bind(&m)
			id, _ := strconv.Atoi(c.Params.ByName("user_id"))
			m.ID = id
			result := userUc.Update(&m)
			c.JSON(http.StatusOK, result)
		})
		// DELETE /api/v1/users/:user_id
		user.DELETE("/:user_id", func(c *gin.Context) {
			m := model.User{}
			c.Bind(&m)
			id, _ := strconv.Atoi(c.Params.ByName("user_id"))
			m.ID = id
			result := userUc.Delete(&m)
			c.JSON(http.StatusOK, result)
		})
	}
}
