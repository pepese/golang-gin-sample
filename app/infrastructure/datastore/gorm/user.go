package gorm

import (
	"github.com/jinzhu/gorm"
	"github.com/pepese/golang-gin-sample/app/domain/model"
)

type User struct {
	gorm.Model
	model.User
}
