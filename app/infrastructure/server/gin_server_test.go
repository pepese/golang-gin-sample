package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pepese/golang-gin-sample/app"
	"github.com/pepese/golang-gin-sample/app/infrastructure/datastore/gorm"
	"github.com/gin-gonic/gin"
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

func TestPostCompanies(t *testing.T) {
	req, _ := http.NewRequest("POST", "/api/v1/users", strings.NewReader(`
	{
		"first_name": "First",
		"last_name": "Last"
	}`))
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()
	testRouter.ServeHTTP(res, req)
	assert.Equal(t, res.Code, 200)
}
