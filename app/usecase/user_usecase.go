package usecase

import (
	"github.com/pepese/golang-gin-sample/app/domain/model"
)

/*
UserUsecase interface.
*/
type UserUsecase interface {
	GetAll() model.Users
}

/*
UserUsecase mode.
*/
type userUsecase struct {
}

/*
GetAll func is UserUsecase implement.
*/
func (uc *userUsecase) GetAll() model.Users {
	return nil
}

/*
NewUserUsecase returns *userUsecase.
*/
func NewUserUsecase() *userUsecase {
	return &userUsecase{}
}
