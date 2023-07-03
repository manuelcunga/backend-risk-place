// Code generated by MockGen. DO NOT EDIT.
// Source: domain/repository/erfce_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entities "github.com/risk-place-angola/backend-risk-place/domain/entities"
)

// MockErfceRepository is a mock of ErfceRepository interface.
type MockErfceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockErfceRepositoryMockRecorder
}

// MockErfceRepositoryMockRecorder is the mock recorder for MockErfceRepository.
type MockErfceRepositoryMockRecorder struct {
	mock *MockErfceRepository
}

// NewMockErfceRepository creates a new mock instance.
func NewMockErfceRepository(ctrl *gomock.Controller) *MockErfceRepository {
	mock := &MockErfceRepository{ctrl: ctrl}
	mock.recorder = &MockErfceRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockErfceRepository) EXPECT() *MockErfceRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockErfceRepository) Delete(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockErfceRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockErfceRepository)(nil).Delete), id)
}

// FindAll mocks base method.
func (m *MockErfceRepository) FindAll() ([]*entities.Erfce, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]*entities.Erfce)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockErfceRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockErfceRepository)(nil).FindAll))
}

// FindAllErffcesWarnings mocks base method.
func (m *MockErfceRepository) FindAllErffcesWarnings() ([]*entities.Erfce, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllErffcesWarnings")
	ret0, _ := ret[0].([]*entities.Erfce)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllErffcesWarnings indicates an expected call of FindAllErffcesWarnings.
func (mr *MockErfceRepositoryMockRecorder) FindAllErffcesWarnings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllErffcesWarnings", reflect.TypeOf((*MockErfceRepository)(nil).FindAllErffcesWarnings))
}

// FindByEmail mocks base method.
func (m *MockErfceRepository) FindByEmail(email string) (*entities.Erfce, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", email)
	ret0, _ := ret[0].(*entities.Erfce)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail.
func (mr *MockErfceRepositoryMockRecorder) FindByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockErfceRepository)(nil).FindByEmail), email)
}

// FindByID mocks base method.
func (m *MockErfceRepository) FindByID(id string) (*entities.Erfce, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", id)
	ret0, _ := ret[0].(*entities.Erfce)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockErfceRepositoryMockRecorder) FindByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockErfceRepository)(nil).FindByID), id)
}

// FindWarningByErfceID mocks base method.
func (m *MockErfceRepository) FindWarningByErfceID(id string) ([]*entities.Erfce, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindWarningByErfceID", id)
	ret0, _ := ret[0].([]*entities.Erfce)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindWarningByErfceID indicates an expected call of FindWarningByErfceID.
func (mr *MockErfceRepositoryMockRecorder) FindWarningByErfceID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindWarningByErfceID", reflect.TypeOf((*MockErfceRepository)(nil).FindWarningByErfceID), id)
}

// Save mocks base method.
func (m *MockErfceRepository) Save(entity *entities.Erfce) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", entity)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockErfceRepositoryMockRecorder) Save(entity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockErfceRepository)(nil).Save), entity)
}

// Update mocks base method.
func (m *MockErfceRepository) Update(entity *entities.Erfce) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", entity)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockErfceRepositoryMockRecorder) Update(entity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockErfceRepository)(nil).Update), entity)
}
