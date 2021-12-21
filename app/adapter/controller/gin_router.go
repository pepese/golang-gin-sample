package controller

import (
	"context"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/pepese/golang-gin-sample/app/adapter/gateway"
	"github.com/pepese/golang-gin-sample/app/domain/model"
	"github.com/pepese/golang-gin-sample/app/usecase"
)

type ginRouter struct {
	e      *gin.Engine
	pingUc usecase.PingUsecaser
	userUc usecase.UserUsecaser
}

func (g *ginRouter) GinRouter() *gin.Engine {
	router := g.e.Group("/api/v1")

	//////////////////
	// Ping UseCase
	//////////////////
	ping := router.Group("/ping")
	if (g.pingUc == nil) || reflect.ValueOf(g.pingUc).IsNil() {
		g.pingUc = &usecase.PingUsecase{}
	}
	// GET /api/v1/ping
	ping.GET("", func(c *gin.Context) {
		result, err := g.pingUc.Say(c.Value("ctx").(context.Context), "pong")
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
		} else {
			c.JSON(http.StatusOK, result)
		}
	})

	//////////////////
	// User UseCase
	//////////////////
	user := router.Group("/users")
	if (g.userUc == nil) || reflect.ValueOf(g.userUc).IsNil() {
		g.userUc = &usecase.UserUsecase{
			UserRepo: db.NewUserRepository(),
		}
	}
	// GET /api/v1/users
	user.GET("", func(c *gin.Context) {
		m := model.User{}
		c.Bind(&m)
		result, err := g.userUc.List(c.Value("ctx").(context.Context), &m)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
		} else {
			c.JSON(http.StatusOK, result)
		}
	})
	// GET /api/v1/users/:user_id
	user.GET("/:user_id", func(c *gin.Context) {
		m := model.User{}
		c.Bind(&m)
		id, _ := strconv.Atoi(c.Params.ByName("user_id"))
		m.ID = id
		result, err := g.userUc.Get(c.Value("ctx").(context.Context), &m)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
		} else {
			c.JSON(http.StatusOK, result)
		}
	})
	// POST /api/v1/users
	user.POST("", func(c *gin.Context) {
		m := model.User{}
		c.Bind(&m)
		result, err := g.userUc.Create(c.Value("ctx").(context.Context), &m)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
		} else {
			c.JSON(http.StatusOK, result)
		}
	})
	// PUT /api/v1/users/:user_id
	user.PUT("/:user_id", func(c *gin.Context) {
		m := model.User{}
		c.Bind(&m)
		id, _ := strconv.Atoi(c.Params.ByName("user_id"))
		m.ID = id
		result, err := g.userUc.Update(c.Value("ctx").(context.Context), &m)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
		} else {
			c.JSON(http.StatusOK, result)
		}
	})
	// DELETE /api/v1/users/:user_id
	user.DELETE("/:user_id", func(c *gin.Context) {
		m := model.User{}
		c.Bind(&m)
		id, _ := strconv.Atoi(c.Params.ByName("user_id"))
		m.ID = id
		result, err := g.userUc.Delete(c.Value("ctx").(context.Context), &m)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
		} else {
			c.JSON(http.StatusOK, result)
		}
	})

	return g.e
}

func NewGinRouter(e *gin.Engine, pingUc usecase.PingUsecaser, userUc usecase.UserUsecaser) *ginRouter {
	return &ginRouter{
		e:      e,
		pingUc: pingUc,
		userUc: userUc,
	}
}
