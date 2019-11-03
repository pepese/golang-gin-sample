package domain

import (
	"github.com/pepese/golang-gin-sample/app/domain/model"
)

/*
UserLogic Interface
*/
type UserLogic interface {
	GetFullName() string
}

type userLogic struct {
	model *model.User
}

func (d *userLogic) GetFullName() string {
	return d.model.FirstName + " " + d.model.LastName
}
