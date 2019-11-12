package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pepese/golang-gin-sample/app"
	"github.com/pepese/golang-gin-sample/app/domain/model"
	"github.com/pepese/golang-gin-sample/app/infrastructure/datastore/gorm"
	"gopkg.in/go-playground/assert.v1"
)

var testRouter *gin.Engine

func TestInit(t *testing.T) {
	app.InitConfig()
	gorm.Init()
	testRouter = NewGinServer()
}

////////////////////////////////////////////////////////////////////////
// TEST /api/v1/companies
////////////////////////////////////////////////////////////////////////

var tmpUser *model.User

func TestPostUsers(t *testing.T) {
	req, _ := http.NewRequest("POST", "/api/v1/users", strings.NewReader(`
	{
		"first_name": "First",
		"last_name": "Last"
	}`))
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()
	testRouter.ServeHTTP(res, req)
	assert.Equal(t, res.Code, 200)
	tmpUser = getStruct(&model.User{}, res).(*model.User)
}

func TestGetUser(t *testing.T) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/users/%d", tmpUser.ID), nil)

	res := httptest.NewRecorder()
	testRouter.ServeHTTP(res, req)
	assert.Equal(t, res.Code, 200)
	resUser := getStruct(&model.User{}, res).(*model.User)
	assert.Equal(t, (*resUser).FirstName, (*tmpUser).FirstName)
	assert.Equal(t, (*resUser).LastName, (*tmpUser).LastName)
}

func TestPutUser(t *testing.T) {
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/api/v1/users/%d", tmpUser.ID), strings.NewReader(`
	{
		"first_name": "FirstX",
		"last_name": "LastX"
	}`))
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()
	testRouter.ServeHTTP(res, req)
	assert.Equal(t, res.Code, 200)
	resUser := getStruct(&model.User{}, res).(*model.User)
	assert.Equal(t, (*resUser).FirstName, "FirstX")
	assert.Equal(t, (*resUser).LastName, "LastX")
	tmpUser = resUser

	req, _ = http.NewRequest("GET", fmt.Sprintf("/api/v1/users/%d", tmpUser.ID), nil)

	res = httptest.NewRecorder()
	testRouter.ServeHTTP(res, req)
	assert.Equal(t, res.Code, 200)
	resUser = getStruct(&model.User{}, res).(*model.User)
	assert.Equal(t, (*resUser).FirstName, (*tmpUser).FirstName)
	assert.Equal(t, (*resUser).LastName, (*tmpUser).LastName)
}

func TestDeleteUser(t *testing.T) {
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/v1/users/%d", tmpUser.ID), nil)

	res := httptest.NewRecorder()
	testRouter.ServeHTTP(res, req)
	assert.Equal(t, res.Code, 200)
	resUser := getStruct(&model.User{}, res).(*model.User)
	assert.Equal(t, (*resUser).FirstName, "")
	assert.Equal(t, (*resUser).LastName, "")

	req, _ = http.NewRequest("GET", fmt.Sprintf("/api/v1/users/%d", tmpUser.ID), nil)

	res = httptest.NewRecorder()
	testRouter.ServeHTTP(res, req)
	assert.Equal(t, res.Code, 200)
	assert.Equal(t, getBodyString(res), "null")
}

////////////////////////////////////////////////////////////////////////
// TEST Utilities
////////////////////////////////////////////////////////////////////////

func getBodyString(res *httptest.ResponseRecorder) string {
	defer res.Result().Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Result().Body)
	return buf.String()
}

func getBodyBytes(res *httptest.ResponseRecorder) []byte {
	defer res.Result().Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Result().Body)
	return buf.Bytes()
}

func getStruct(entity interface{}, res *httptest.ResponseRecorder) interface{} {
	jsonBytes := getBodyBytes(res)
	if err := json.Unmarshal(jsonBytes, entity); err != nil {
		return nil
	}
	return entity
}
