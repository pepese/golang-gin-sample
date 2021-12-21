package main

import (
	"github.com/pepese/golang-gin-sample/app"
	"github.com/pepese/golang-gin-sample/app/infrastructure/datastore"
	"github.com/pepese/golang-gin-sample/app/infrastructure/server"
)

func main() {
	app.Config()

	gm := datastore.Gorm()
	db, err := gm.DB()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	srv := server.NewGinServer()
	srv.Run()
}
