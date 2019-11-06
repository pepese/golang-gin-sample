package main

import (
	"github.com/pepese/golang-gin-sample/app"
	"github.com/pepese/golang-gin-sample/app/infrastructure/datastore/gorm"
	"github.com/pepese/golang-gin-sample/app/infrastructure/server"
)

func main() {
	app.InitConfig()

	db := gorm.Init()
	defer db.Close()

	srv := server.NewGinServer()
	srv.Run()
}
