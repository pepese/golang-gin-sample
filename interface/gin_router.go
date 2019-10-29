package interface

import {
	"github.com/pepese/golang-gin-sample/usecase"
}

func Route(api *gin.Engine) {
	router := api.Group("")

	ping := usecase.Ping{}{
		PingRoute := router.Group("/ping")
		PingRoute.GET("", ping.GET)
	}
}