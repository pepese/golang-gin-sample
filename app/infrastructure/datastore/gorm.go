package datastore

import (
	"fmt"
	"sync"

	"github.com/pepese/golang-gin-sample/app"
	"github.com/pepese/golang-gin-sample/app/domain/model"
	"gorm.io/driver/postgres"
	gm "gorm.io/gorm"
)

var (
	gorm  *gm.DB
	gOnce sync.Once
)

func Gorm() *gm.DB {
	initGorm()
	return gorm
}

func initGorm() {
	gOnce.Do(func() {
		conf := app.Config()
		// MySQL
		// connectTemplate := "%s:%s@%s(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
		// con := fmt.Sprintf(connectTemplate, conf.RdbUser, conf.RdbPassword, conf.RdbProtocol, conf.RdbHost, conf.RdbPort, conf.RdbName)
		// db, err := gm.Open(mysql.Open(con), &gm.Config{})
		// PostgreSQL
		connectTemplate := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s"
		dsn := fmt.Sprintf(connectTemplate, conf.RdbHost, conf.RdbUser, conf.RdbPassword, conf.RdbName, conf.RdbPort, conf.TimeZone)
		db, err := gm.Open(postgres.Open(dsn), &gm.Config{})
		if err != nil {
			panic(err.Error())
		}
		db.AutoMigrate(&model.User{})
		gorm = db
	})
}
