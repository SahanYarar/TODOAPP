// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import (
	entities "todoapi/entities"

	mock "github.com/stretchr/testify/mock"
)

// ToDoRepositoryInterface is an autogenerated mock type for the ToDoRepositoryInterface type
type ToDoRepositoryInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: u
func (_m *ToDoRepositoryInterface) Create(u *entities.ToDo) error {
	ret := _m.Called(u)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.ToDo) error); ok {
		r0 = rf(u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *ToDoRepositoryInterface) Delete(id uint64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: id
func (_m *ToDoRepositoryInterface) Get(id uint64) (*entities.ToDo, error) {
	ret := _m.Called(id)

	var r0 *entities.ToDo
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (*entities.ToDo, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint64) *entities.ToDo); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.ToDo)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: u
func (_m *ToDoRepositoryInterface) GetAll(u []*entities.ToDo) ([]*entities.ToDo, error) {
	ret := _m.Called(u)

	var r0 []*entities.ToDo
	var r1 error
	if rf, ok := ret.Get(0).(func([]*entities.ToDo) ([]*entities.ToDo, error)); ok {
		return rf(u)
	}
	if rf, ok := ret.Get(0).(func([]*entities.ToDo) []*entities.ToDo); ok {
		r0 = rf(u)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.ToDo)
		}
	}

	if rf, ok := ret.Get(1).(func([]*entities.ToDo) error); ok {
		r1 = rf(u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: u
func (_m *ToDoRepositoryInterface) Update(u *entities.ToDo) error {
	ret := _m.Called(u)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.ToDo) error); ok {
		r0 = rf(u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewToDoRepositoryInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewToDoRepositoryInterface creates a new instance of ToDoRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewToDoRepositoryInterface(t mockConstructorTestingTNewToDoRepositoryInterface) *ToDoRepositoryInterface {
	mock := &ToDoRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}