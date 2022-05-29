// Code generated by mockery v2.13.0-beta.1. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/bangumi/server/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// AuthRepo is an autogenerated mock type for the AuthRepo type
type AuthRepo struct {
	mock.Mock
}

type AuthRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *AuthRepo) EXPECT() *AuthRepo_Expecter {
	return &AuthRepo_Expecter{mock: &_m.Mock}
}

// GetByEmail provides a mock function with given fields: ctx, email
func (_m *AuthRepo) GetByEmail(ctx context.Context, email string) (domain.Auth, []byte, error) {
	ret := _m.Called(ctx, email)

	var r0 domain.Auth
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Auth); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(domain.Auth)
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func(context.Context, string) []byte); ok {
		r1 = rf(ctx, email)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, email)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// AuthRepo_GetByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByEmail'
type AuthRepo_GetByEmail_Call struct {
	*mock.Call
}

// GetByEmail is a helper method to define mock.On call
//  - ctx context.Context
//  - email string
func (_e *AuthRepo_Expecter) GetByEmail(ctx interface{}, email interface{}) *AuthRepo_GetByEmail_Call {
	return &AuthRepo_GetByEmail_Call{Call: _e.mock.On("GetByEmail", ctx, email)}
}

func (_c *AuthRepo_GetByEmail_Call) Run(run func(ctx context.Context, email string)) *AuthRepo_GetByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AuthRepo_GetByEmail_Call) Return(_a0 domain.Auth, _a1 []byte, _a2 error) *AuthRepo_GetByEmail_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

// GetByToken provides a mock function with given fields: ctx, token
func (_m *AuthRepo) GetByToken(ctx context.Context, token string) (domain.Auth, error) {
	ret := _m.Called(ctx, token)

	var r0 domain.Auth
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Auth); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Get(0).(domain.Auth)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthRepo_GetByToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByToken'
type AuthRepo_GetByToken_Call struct {
	*mock.Call
}

// GetByToken is a helper method to define mock.On call
//  - ctx context.Context
//  - token string
func (_e *AuthRepo_Expecter) GetByToken(ctx interface{}, token interface{}) *AuthRepo_GetByToken_Call {
	return &AuthRepo_GetByToken_Call{Call: _e.mock.On("GetByToken", ctx, token)}
}

func (_c *AuthRepo_GetByToken_Call) Run(run func(ctx context.Context, token string)) *AuthRepo_GetByToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AuthRepo_GetByToken_Call) Return(_a0 domain.Auth, _a1 error) *AuthRepo_GetByToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetPermission provides a mock function with given fields: ctx, groupID
func (_m *AuthRepo) GetPermission(ctx context.Context, groupID uint8) (domain.Permission, error) {
	ret := _m.Called(ctx, groupID)

	var r0 domain.Permission
	if rf, ok := ret.Get(0).(func(context.Context, uint8) domain.Permission); ok {
		r0 = rf(ctx, groupID)
	} else {
		r0 = ret.Get(0).(domain.Permission)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint8) error); ok {
		r1 = rf(ctx, groupID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthRepo_GetPermission_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPermission'
type AuthRepo_GetPermission_Call struct {
	*mock.Call
}

// GetPermission is a helper method to define mock.On call
//  - ctx context.Context
//  - groupID uint8
func (_e *AuthRepo_Expecter) GetPermission(ctx interface{}, groupID interface{}) *AuthRepo_GetPermission_Call {
	return &AuthRepo_GetPermission_Call{Call: _e.mock.On("GetPermission", ctx, groupID)}
}

func (_c *AuthRepo_GetPermission_Call) Run(run func(ctx context.Context, groupID uint8)) *AuthRepo_GetPermission_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint8))
	})
	return _c
}

func (_c *AuthRepo_GetPermission_Call) Return(_a0 domain.Permission, _a1 error) *AuthRepo_GetPermission_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type NewAuthRepoT interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthRepo creates a new instance of AuthRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthRepo(t NewAuthRepoT) *AuthRepo {
	mock := &AuthRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
