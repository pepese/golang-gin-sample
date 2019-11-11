package usecase

import (
	"github.com/pepese/golang-gin-sample/app/domain/model"
	db "github.com/pepese/golang-gin-sample/app/infrastructure/datastore/gorm"
	"github.com/pepese/golang-gin-sample/app/interface/presenter"
)

/*
UserUsecase interface.
*/
// type UserUsecase interface {
// 	List() model.Users
// }

/*
UserUsecase mode.
*/
type UserUsecase struct {
}

/*
List func is UserUsecase implement.
*/
func (UserUsecase) List(m *model.User) model.Users {
	var r presenter.UserRepository
	r = db.NewUserRepository()
	result, _ := r.List(m)
	return result
}

func (UserUsecase) Get(m *model.User) *model.User {
	var r presenter.UserRepository
	r = db.NewUserRepository()
	result, _ := r.Get(m)
	return result
}

func (UserUsecase) Create(m *model.User) *model.User {
	var r presenter.UserRepository
	r = db.NewUserRepository()
	result, _ := r.Create(m)
	return result
}

func (UserUsecase) Update(m *model.User) *model.User {
	var r presenter.UserRepository
	r = db.NewUserRepository()
	result, _ := r.Update(m)
	return result
}

func (UserUsecase) Delete(m *model.User) *model.User {
	var r presenter.UserRepository
	r = db.NewUserRepository()
	result, _ := r.Delete(m)
	return result
}

/*
NewUserUsecase returns *userUsecase.
*/
// func NewUserUsecase() *userUsecase {
// 	return &userUsecase{}
// }
