// Code generated by MockGen. DO NOT EDIT.
// Source: app/usecase/user.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/pepese/golang-gin-sample/app/domain/model"
)

// MockUserUsecaser is a mock of UserUsecaser interface.
type MockUserUsecaser struct {
	ctrl     *gomock.Controller
	recorder *MockUserUsecaserMockRecorder
}

// MockUserUsecaserMockRecorder is the mock recorder for MockUserUsecaser.
type MockUserUsecaserMockRecorder struct {
	mock *MockUserUsecaser
}

// NewMockUserUsecaser creates a new mock instance.
func NewMockUserUsecaser(ctrl *gomock.Controller) *MockUserUsecaser {
	mock := &MockUserUsecaser{ctrl: ctrl}
	mock.recorder = &MockUserUsecaserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUsecaser) EXPECT() *MockUserUsecaserMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockUserUsecaser) Create(c context.Context, m *model.User) (*model.User, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", c, m)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserUsecaserMockRecorder) Create(c, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserUsecaser)(nil).Create), c, m)
}

// Delete mocks base method.
func (m_2 *MockUserUsecaser) Delete(c context.Context, m *model.User) (*model.User, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Delete", c, m)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockUserUsecaserMockRecorder) Delete(c, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserUsecaser)(nil).Delete), c, m)
}

// Get mocks base method.
func (m_2 *MockUserUsecaser) Get(c context.Context, m *model.User) (*model.User, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Get", c, m)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUserUsecaserMockRecorder) Get(c, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserUsecaser)(nil).Get), c, m)
}

// List mocks base method.
func (m_2 *MockUserUsecaser) List(c context.Context, m *model.User) (model.Users, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "List", c, m)
	ret0, _ := ret[0].(model.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockUserUsecaserMockRecorder) List(c, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUserUsecaser)(nil).List), c, m)
}

// Update mocks base method.
func (m_2 *MockUserUsecaser) Update(c context.Context, m *model.User) (*model.User, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", c, m)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUserUsecaserMockRecorder) Update(c, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserUsecaser)(nil).Update), c, m)
}