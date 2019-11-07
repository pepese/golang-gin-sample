package gorm

import (
	"github.com/pepese/golang-gin-sample/app/domain/model"
)

type userRepository struct{}

func (repo *userRepository) List(m *model.User) ([]model.User, error) {
	searched := []model.User{}
	if result := rdb.Find(&searched, m); result.Error != nil {
		return nil, result.Error
	}
	return searched, nil
}

/*
検索でヒットしない場合は nil が返る
*/
func (repo *userRepository) Get(m *model.User) (*model.User, error) {
	searched := &model.User{}
	if result := rdb.First(searched, m); result.Error != nil {
		return nil, result.Error
	}
	return searched, nil
}

func (repo *userRepository) Create(m *model.User) (*model.User, error) {
	if result := rdb.Create(m); result.Error != nil {
		return nil, result.Error
	}
	return m, nil
}

func (repo *userRepository) Update(m *model.User) (*model.User, error) {
	if result := rdb.Save(m); result.Error != nil {
		return nil, result.Error
	}
	return m, nil
}

func (repo *userRepository) Delete(m *model.User) (*model.User, error) {
	if result := rdb.Delete(m); result.Error != nil {
		return nil, result.Error
	}
	return m, nil
}

func NewUserRepository() *userRepository {
	return &userRepository{}
}
