// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nashmaniac/hi-mama/hi-mama-backend/adapters (interfaces: PeristenceStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/nashmaniac/hi-mama/hi-mama-backend/models"
)

// MockPeristenceStore is a mock of PeristenceStore interface.
type MockPeristenceStore struct {
	ctrl     *gomock.Controller
	recorder *MockPeristenceStoreMockRecorder
}

// MockPeristenceStoreMockRecorder is the mock recorder for MockPeristenceStore.
type MockPeristenceStoreMockRecorder struct {
	mock *MockPeristenceStore
}

// NewMockPeristenceStore creates a new mock instance.
func NewMockPeristenceStore(ctrl *gomock.Controller) *MockPeristenceStore {
	mock := &MockPeristenceStore{ctrl: ctrl}
	mock.recorder = &MockPeristenceStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPeristenceStore) EXPECT() *MockPeristenceStoreMockRecorder {
	return m.recorder
}

// CloseDB mocks base method.
func (m *MockPeristenceStore) CloseDB(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CloseDB", arg0)
}

// CloseDB indicates an expected call of CloseDB.
func (mr *MockPeristenceStoreMockRecorder) CloseDB(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseDB", reflect.TypeOf((*MockPeristenceStore)(nil).CloseDB), arg0)
}

// CreateEntry mocks base method.
func (m *MockPeristenceStore) CreateEntry(arg0 context.Context, arg1 *models.Entry) (*models.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEntry", arg0, arg1)
	ret0, _ := ret[0].(*models.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEntry indicates an expected call of CreateEntry.
func (mr *MockPeristenceStoreMockRecorder) CreateEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEntry", reflect.TypeOf((*MockPeristenceStore)(nil).CreateEntry), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockPeristenceStore) CreateUser(arg0 context.Context, arg1 *models.User) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockPeristenceStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockPeristenceStore)(nil).CreateUser), arg0, arg1)
}

// FindEntries mocks base method.
func (m *MockPeristenceStore) FindEntries(arg0 context.Context, arg1 *models.User) ([]models.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEntries", arg0, arg1)
	ret0, _ := ret[0].([]models.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindEntries indicates an expected call of FindEntries.
func (mr *MockPeristenceStoreMockRecorder) FindEntries(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEntries", reflect.TypeOf((*MockPeristenceStore)(nil).FindEntries), arg0, arg1)
}

// FindEntryById mocks base method.
func (m *MockPeristenceStore) FindEntryById(arg0 context.Context, arg1 uint) (*models.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEntryById", arg0, arg1)
	ret0, _ := ret[0].(*models.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindEntryById indicates an expected call of FindEntryById.
func (mr *MockPeristenceStoreMockRecorder) FindEntryById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEntryById", reflect.TypeOf((*MockPeristenceStore)(nil).FindEntryById), arg0, arg1)
}

// FindOngoingTime mocks base method.
func (m *MockPeristenceStore) FindOngoingTime(arg0 context.Context, arg1 models.User) (*models.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOngoingTime", arg0, arg1)
	ret0, _ := ret[0].(*models.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOngoingTime indicates an expected call of FindOngoingTime.
func (mr *MockPeristenceStoreMockRecorder) FindOngoingTime(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOngoingTime", reflect.TypeOf((*MockPeristenceStore)(nil).FindOngoingTime), arg0, arg1)
}

// FindUserByUsername mocks base method.
func (m *MockPeristenceStore) FindUserByUsername(arg0 context.Context, arg1 string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByUsername", arg0, arg1)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByUsername indicates an expected call of FindUserByUsername.
func (mr *MockPeristenceStoreMockRecorder) FindUserByUsername(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByUsername", reflect.TypeOf((*MockPeristenceStore)(nil).FindUserByUsername), arg0, arg1)
}

// SaveEntry mocks base method.
func (m *MockPeristenceStore) SaveEntry(arg0 context.Context, arg1 *models.Entry) (*models.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveEntry", arg0, arg1)
	ret0, _ := ret[0].(*models.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveEntry indicates an expected call of SaveEntry.
func (mr *MockPeristenceStoreMockRecorder) SaveEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveEntry", reflect.TypeOf((*MockPeristenceStore)(nil).SaveEntry), arg0, arg1)
}
