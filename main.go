package main

import (
	"github.com/pepese/golang-gin-sample/app/infrastructure/server"
	"github.com/pepese/golang-gin-sample/app/infrastructure/datastore/gorm"
)

func main() {
	gorm.InitDB()
	srv := server.NewGinServer()
	srv.Run()
}
