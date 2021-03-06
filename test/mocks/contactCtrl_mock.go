// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/controller/contact.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entities "github.com/STreeChin/contactapi/pkg/entities"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockContactService is a mock of ContactService interface
type MockContactService struct {
	ctrl     *gomock.Controller
	recorder *MockContactServiceMockRecorder
}

// MockContactServiceMockRecorder is the mock recorder for MockContactService
type MockContactServiceMockRecorder struct {
	mock *MockContactService
}

// NewMockContactService creates a new mock instance
func NewMockContactService(ctrl *gomock.Controller) *MockContactService {
	mock := &MockContactService{ctrl: ctrl}
	mock.recorder = &MockContactServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContactService) EXPECT() *MockContactServiceMockRecorder {
	return m.recorder
}

// AddOrUpdateContact mocks base method
func (m *MockContactService) AddOrUpdateContact(contact *entities.Contact) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOrUpdateContact", contact)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddOrUpdateContact indicates an expected call of AddOrUpdateContact
func (mr *MockContactServiceMockRecorder) AddOrUpdateContact(contact interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrUpdateContact", reflect.TypeOf((*MockContactService)(nil).AddOrUpdateContact), contact)
}

// GetOneContact mocks base method
func (m *MockContactService) GetOneContact(key, value string) (*entities.Contact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOneContact", key, value)
	ret0, _ := ret[0].(*entities.Contact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOneContact indicates an expected call of GetOneContact
func (mr *MockContactServiceMockRecorder) GetOneContact(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOneContact", reflect.TypeOf((*MockContactService)(nil).GetOneContact), key, value)
}
