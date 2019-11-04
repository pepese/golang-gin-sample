package main

import (
	"github.com/pepese/golang-gin-sample/app/infrastructure/datastore/gorm"
	"github.com/pepese/golang-gin-sample/app/infrastructure/server"
)

func main() {
	gorm.Init()
	srv := server.NewGinServer()
	srv.Run()
}
