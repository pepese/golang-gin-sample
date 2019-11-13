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
type UserUsecase struct{}

var userRepository presenter.UserRepository

func (UserUsecase) Init() {
	userRepository = db.NewUserRepository()
}

/*
List func is UserUsecase implement.
*/
func (UserUsecase) List(m *model.User) model.Users {
	result, _ := userRepository.List(m)
	return result
}

func (UserUsecase) Get(m *model.User) *model.User {
	result, _ := userRepository.Get(m)
	return result
}

func (UserUsecase) Create(m *model.User) *model.User {
	result, _ := userRepository.Create(m)
	return result
}

func (UserUsecase) Update(m *model.User) *model.User {
	result, _ := userRepository.Update(m)
	return result
}

func (UserUsecase) Delete(m *model.User) *model.User {
	result, _ := userRepository.Delete(m)
	return result
}

/*
NewUserUsecase returns *userUsecase.
*/
// func NewUserUsecase() *userUsecase {
// 	return &userUsecase{}
// }
