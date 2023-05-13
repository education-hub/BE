// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	user "github.com/education-hub/BE/app/entities/user"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// ForgetPass provides a mock function with given fields: ctx, email
func (_m *UserService) ForgetPass(ctx context.Context, email string) error {
	ret := _m.Called(ctx, email)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Login provides a mock function with given fields: ctx, req
func (_m *UserService) Login(ctx context.Context, req user.LoginReq) (int, string, error) {
	ret := _m.Called(ctx, req)

	var r0 int
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, user.LoginReq) (int, string, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, user.LoginReq) int); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, user.LoginReq) string); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, user.LoginReq) error); ok {
		r2 = rf(ctx, req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Register provides a mock function with given fields: ctx, req
func (_m *UserService) Register(ctx context.Context, req user.RegisterReq) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, user.RegisterReq) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResetPass provides a mock function with given fields: ctx, token, newpass
func (_m *UserService) ResetPass(ctx context.Context, token string, newpass string) error {
	ret := _m.Called(ctx, token, newpass)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, token, newpass)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifyEmail provides a mock function with given fields: ctx, verificationcode
func (_m *UserService) VerifyEmail(ctx context.Context, verificationcode string) error {
	ret := _m.Called(ctx, verificationcode)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, verificationcode)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUserService interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserService(t mockConstructorTestingTNewUserService) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}