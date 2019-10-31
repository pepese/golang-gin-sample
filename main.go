package main

import "github.com/pepese/golang-gin-sample/app/infrastructure/server"

// import "github.com/gin-gonic/gin"

// func main() {
// 	r := gin.Default()
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "pong",
// 		})
// 	})
// 	r.Run() // listen and serve on 0.0.0.0:8080
// }

func main() {
	srv := server.NewGinServer()
	srv.Run()
}
