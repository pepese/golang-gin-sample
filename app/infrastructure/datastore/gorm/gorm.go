package gorm

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pepese/golang-gin-sample/app"
	"github.com/pepese/golang-gin-sample/app/domain/model"
)

var rdb *gorm.DB

func Init() *gorm.DB {
	config := app.GetConfig()
	connectTemplate := "%s:%s@%s(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	connect := fmt.Sprintf(connectTemplate, config.RdbUser, config.RdbPassword, config.RdbProtocol, config.RdbHost, config.RdbName)
	db, err := gorm.Open(config.RdbType, connect)
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&model.User{})
	rdb = db

	return db
}
