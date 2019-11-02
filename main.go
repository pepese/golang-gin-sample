package main

import "github.com/pepese/golang-gin-sample/app/infrastructure/server"

func main() {
	srv := server.NewGinServer()
	srv.Run()
}
