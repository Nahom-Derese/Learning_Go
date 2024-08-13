// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/domain"
	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: username
func (_m *UserRepository) Delete(username string) error {
	ret := _m.Called(username)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAll provides a mock function with given fields:
func (_m *UserRepository) DeleteAll() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DeleteAll")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields:
func (_m *UserRepository) FindAll() []domain.User {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 []domain.User
	if rf, ok := ret.Get(0).(func() []domain.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.User)
		}
	}

	return r0
}

// FindByUsername provides a mock function with given fields: username
func (_m *UserRepository) FindByUsername(username string) (domain.User, bool) {
	ret := _m.Called(username)

	if len(ret) == 0 {
		panic("no return value specified for FindByUsername")
	}

	var r0 domain.User
	var r1 bool
	if rf, ok := ret.Get(0).(func(string) (domain.User, bool)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) domain.User); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// FindUser provides a mock function with given fields: id
func (_m *UserRepository) FindUser(id string) (domain.User, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindUser")
	}

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (domain.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) domain.User); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: user
func (_m *UserRepository) Save(user domain.User) (domain.User, error) {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.User) (domain.User, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(domain.User) domain.User); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(domain.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: username, user
func (_m *UserRepository) Update(username string, user domain.User) (domain.User, error) {
	ret := _m.Called(username, user)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string, domain.User) (domain.User, error)); ok {
		return rf(username, user)
	}
	if rf, ok := ret.Get(0).(func(string, domain.User) domain.User); ok {
		r0 = rf(username, user)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(string, domain.User) error); ok {
		r1 = rf(username, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
