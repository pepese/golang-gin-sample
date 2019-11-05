package gorm

import (
	"testing"

	"github.com/pepese/golang-gin-sample/app/domain/model"
)

var user *model.User
var repo = NewUserRepository()

func TestCreate(t *testing.T) {
	Init()
	user = &model.User{FirstName: "testFirst", LastName: "testLast"}
	result, err := repo.Create(user)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if *result != *user {
		t.Fatal("failed test")
	}
	user = result
}

func TestGet(t *testing.T) {
	a := &model.User{ID: (*user).ID}
	result, err := repo.Get(a)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if *result != *user {
		t.Fatal("failed test")
	}
}
