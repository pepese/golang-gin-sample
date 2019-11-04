package presenter

import (
	"github.com/pepese/golang-gin-sample/app/domain/model"
)

type UserRepository interface {
	List(model *model.User) ([]model.User, error)
	Get(model *model.User) (*model.User, error)
	Create(model *model.User) error
	Update(model *model.User) error
	Delete(model *model.User) error
}
