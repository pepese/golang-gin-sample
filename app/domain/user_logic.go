package domain

import (
	"github.com/pepese/golang-gin-sample/app/domain/model"
)

/*
UserLogic Interface
*/
type UserLogic interface {
	FullName() string
}

type userLogic struct {
	model *model.User
}

func (d *userLogic) FullName() string {
	return d.model.FirstName + " " + d.model.LastName
}
