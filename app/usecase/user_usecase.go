package usecase

import (
	"context"

	"github.com/pepese/golang-gin-sample/app"
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
func (UserUsecase) List(c context.Context, m *model.User) model.Users {
	logger := app.GetLoggerFromContext(c)
	result, err := userRepository.List(m)
	if err != nil {
		logger.Error(err)
	}
	return result
}

func (UserUsecase) Get(c context.Context, m *model.User) *model.User {
	logger := app.GetLoggerFromContext(c)
	result, err := userRepository.Get(m)
	if err != nil {
		logger.Error(err)
	}
	return result
}

func (UserUsecase) Create(c context.Context, m *model.User) *model.User {
	logger := app.GetLoggerFromContext(c)
	result, err := userRepository.Create(m)
	if err != nil {
		logger.Error(err)
	}
	return result
}

func (UserUsecase) Update(c context.Context, m *model.User) *model.User {
	logger := app.GetLoggerFromContext(c)
	result, err := userRepository.Update(m)
	if err != nil {
		logger.Error(err)
	}
	return result
}

func (UserUsecase) Delete(c context.Context, m *model.User) *model.User {
	logger := app.GetLoggerFromContext(c)
	result, err := userRepository.Delete(m)
	if err != nil {
		logger.Error(err)
	}
	return result
}

/*
NewUserUsecase returns *userUsecase.
*/
// func NewUserUsecase() *userUsecase {
// 	return &userUsecase{}
// }
