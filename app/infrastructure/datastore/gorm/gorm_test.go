package gorm

import (
	"testing"

	"github.com/pepese/golang-gin-sample/app"
	"github.com/pepese/golang-gin-sample/app/domain/model"
)

func TestUserRepository(t *testing.T) {
	var user *model.User
	var repo = NewUserRepository()
	app.InitConfig()
	rdb := Init()
	defer rdb.Close()

	t.Run("Create user", func(t *testing.T) {
		user = &model.User{FirstName: "testFirst", LastName: "testLast"}
		result, err := repo.Create(user)
		if err != nil {
			t.Fatalf("failed test %#v", err)
		}
		if *result != *user {
			t.Fatal("failed test")
		}
		user = result
	})

	t.Run("Create same user", func(t *testing.T) {
		result, err := repo.Create(user)
		if result != nil && err == nil {
			t.Fatalf("failed test %#v", err)
		}
	})

	t.Run("Get user", func(t *testing.T) {
		a := &model.User{ID: (*user).ID}
		result, err := repo.Get(a)
		if err != nil {
			t.Fatalf("failed test %#v", err)
		}
		if (*result).FirstName != (*user).FirstName || (*result).LastName != (*user).LastName {
			t.Fatal("failed test")
		}
	})

	t.Run("List users", func(t *testing.T) {
		a := &model.User{ID: (*user).ID}
		result, err := repo.List(a)
		if err != nil {
			t.Fatalf("failed test %#v", err)
		}
		element := result[0]
		if element.FirstName != (*user).FirstName || element.LastName != (*user).LastName {
			t.Fatal("failed test")
		}
	})

	t.Run("Update user", func(t *testing.T) {
		a := &model.User{ID: (*user).ID, FirstName: "testFirst2"}
		result, err := repo.Update(a)
		if err != nil {
			t.Fatalf("failed test %#v", err)
		}
		if (*result).FirstName != "testFirst2" {
			t.Fatal("failed test")
		}
		user = result
	})

	t.Run("Delete user", func(t *testing.T) {
		result, err := repo.Delete(user)
		if err != nil {
			t.Fatalf("failed test %#v", err)
		}
		if *result != *user {
			t.Fatal("failed test")
		}
		result, _ = repo.Get(user)
		if result != nil {
			t.Fatal("failed test")
		}
	})
}
