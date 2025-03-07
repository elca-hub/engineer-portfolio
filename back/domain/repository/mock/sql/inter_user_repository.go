// Code generated by MockGen. DO NOT EDIT.
// Source: inter_user_repository.go
//
// Generated by this command:
//
//	mockgen -source=inter_user_repository.go -package=mock_sql -destination=../mock/sql/inter_user_repository.go
//

// Package mock_sql is a generated GoMock package.
package mock_sql

import (
	model "devport/domain/model"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
	isgomock struct{}
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserRepository) Create(u *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", u)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserRepositoryMockRecorder) Create(u any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRepository)(nil).Create), u)
}

// Exists mocks base method.
func (m *MockUserRepository) Exists(email *model.Email) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", email)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockUserRepositoryMockRecorder) Exists(email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockUserRepository)(nil).Exists), email)
}

// ExistsByName mocks base method.
func (m *MockUserRepository) ExistsByName(name string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistsByName", name)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExistsByName indicates an expected call of ExistsByName.
func (mr *MockUserRepositoryMockRecorder) ExistsByName(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistsByName", reflect.TypeOf((*MockUserRepository)(nil).ExistsByName), name)
}

// FetchInConfirmationUsers mocks base method.
func (m *MockUserRepository) FetchInConfirmationUsers() ([]*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchInConfirmationUsers")
	ret0, _ := ret[0].([]*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchInConfirmationUsers indicates an expected call of FetchInConfirmationUsers.
func (mr *MockUserRepositoryMockRecorder) FetchInConfirmationUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchInConfirmationUsers", reflect.TypeOf((*MockUserRepository)(nil).FetchInConfirmationUsers))
}

// FindByEmail mocks base method.
func (m *MockUserRepository) FindByEmail(email *model.Email) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", email)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail.
func (mr *MockUserRepositoryMockRecorder) FindByEmail(email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindByEmail), email)
}

// Update mocks base method.
func (m *MockUserRepository) Update(u *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", u)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserRepositoryMockRecorder) Update(u any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserRepository)(nil).Update), u)
}
