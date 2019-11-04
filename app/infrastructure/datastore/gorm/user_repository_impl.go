package gorm

import (
	"github.com/pepese/golang-gin-sample/app/domain/model"
)

type userRepository struct{}

func (repo *userRepository) List(model *model.User) ([]model.User, error) {
	return nil, nil
}

func (repo *userRepository) Get(model *model.User) (*model.User, error) {
	return nil, nil
}

func (repo *userRepository) Create(model *model.User) error {
	return nil
}

func (repo *userRepository) Update(model *model.User) error {
	return nil
}

func (repo *userRepository) Delete(model *model.User) error {
	return nil
}

func NewUserRepository() *userRepository {
	return &userRepository{}
}