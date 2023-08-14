// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/MuhAndriJP/dealls.git/entity"
	mock "github.com/stretchr/testify/mock"
)

// IUser is an autogenerated mock type for the IUser type
type IUser struct {
	mock.Mock
}

// GetUserByEmail provides a mock function with given fields: _a0, _a1
func (_m *IUser) GetUserByEmail(_a0 context.Context, _a1 string) (entity.Users, error) {
	ret := _m.Called(_a0, _a1)

	var r0 entity.Users
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (entity.Users, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) entity.Users); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(entity.Users)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: _a0, _a1
func (_m *IUser) Insert(_a0 context.Context, _a1 *entity.Users) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Users) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Upsert provides a mock function with given fields: _a0, _a1
func (_m *IUser) Upsert(_a0 context.Context, _a1 *entity.Users) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Users) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIUser creates a new instance of IUser. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIUser(t interface {
	mock.TestingT
	Cleanup(func())
}) *IUser {
	mock := &IUser{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
