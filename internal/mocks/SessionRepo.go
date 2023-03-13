// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	context "context"

	session "github.com/bangumi/server/web/session"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// SessionRepo is an autogenerated mock type for the Repo type
type SessionRepo struct {
	mock.Mock
}

type SessionRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *SessionRepo) EXPECT() *SessionRepo_Expecter {
	return &SessionRepo_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, userID, regTime, keyGen
func (_m *SessionRepo) Create(ctx context.Context, userID uint32, regTime time.Time, keyGen func() string) (string, session.Session, error) {
	ret := _m.Called(ctx, userID, regTime, keyGen)

	var r0 string
	var r1 session.Session
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, time.Time, func() string) (string, session.Session, error)); ok {
		return rf(ctx, userID, regTime, keyGen)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32, time.Time, func() string) string); ok {
		r0 = rf(ctx, userID, regTime, keyGen)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32, time.Time, func() string) session.Session); ok {
		r1 = rf(ctx, userID, regTime, keyGen)
	} else {
		r1 = ret.Get(1).(session.Session)
	}

	if rf, ok := ret.Get(2).(func(context.Context, uint32, time.Time, func() string) error); ok {
		r2 = rf(ctx, userID, regTime, keyGen)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SessionRepo_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type SessionRepo_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uint32
//   - regTime time.Time
//   - keyGen func() string
func (_e *SessionRepo_Expecter) Create(ctx interface{}, userID interface{}, regTime interface{}, keyGen interface{}) *SessionRepo_Create_Call {
	return &SessionRepo_Create_Call{Call: _e.mock.On("Create", ctx, userID, regTime, keyGen)}
}

func (_c *SessionRepo_Create_Call) Run(run func(ctx context.Context, userID uint32, regTime time.Time, keyGen func() string)) *SessionRepo_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(time.Time), args[3].(func() string))
	})
	return _c
}

func (_c *SessionRepo_Create_Call) Return(key string, s session.Session, err error) *SessionRepo_Create_Call {
	_c.Call.Return(key, s, err)
	return _c
}

func (_c *SessionRepo_Create_Call) RunAndReturn(run func(context.Context, uint32, time.Time, func() string) (string, session.Session, error)) *SessionRepo_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, key
func (_m *SessionRepo) Get(ctx context.Context, key string) (session.Session, error) {
	ret := _m.Called(ctx, key)

	var r0 session.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (session.Session, error)); ok {
		return rf(ctx, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) session.Session); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(session.Session)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SessionRepo_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type SessionRepo_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
func (_e *SessionRepo_Expecter) Get(ctx interface{}, key interface{}) *SessionRepo_Get_Call {
	return &SessionRepo_Get_Call{Call: _e.mock.On("Get", ctx, key)}
}

func (_c *SessionRepo_Get_Call) Run(run func(ctx context.Context, key string)) *SessionRepo_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *SessionRepo_Get_Call) Return(_a0 session.Session, _a1 error) *SessionRepo_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SessionRepo_Get_Call) RunAndReturn(run func(context.Context, string) (session.Session, error)) *SessionRepo_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Revoke provides a mock function with given fields: ctx, key
func (_m *SessionRepo) Revoke(ctx context.Context, key string) error {
	ret := _m.Called(ctx, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SessionRepo_Revoke_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Revoke'
type SessionRepo_Revoke_Call struct {
	*mock.Call
}

// Revoke is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
func (_e *SessionRepo_Expecter) Revoke(ctx interface{}, key interface{}) *SessionRepo_Revoke_Call {
	return &SessionRepo_Revoke_Call{Call: _e.mock.On("Revoke", ctx, key)}
}

func (_c *SessionRepo_Revoke_Call) Run(run func(ctx context.Context, key string)) *SessionRepo_Revoke_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *SessionRepo_Revoke_Call) Return(_a0 error) *SessionRepo_Revoke_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SessionRepo_Revoke_Call) RunAndReturn(run func(context.Context, string) error) *SessionRepo_Revoke_Call {
	_c.Call.Return(run)
	return _c
}

// RevokeUser provides a mock function with given fields: ctx, userID
func (_m *SessionRepo) RevokeUser(ctx context.Context, userID uint32) ([]string, error) {
	ret := _m.Called(ctx, userID)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32) ([]string, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32) []string); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SessionRepo_RevokeUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RevokeUser'
type SessionRepo_RevokeUser_Call struct {
	*mock.Call
}

// RevokeUser is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uint32
func (_e *SessionRepo_Expecter) RevokeUser(ctx interface{}, userID interface{}) *SessionRepo_RevokeUser_Call {
	return &SessionRepo_RevokeUser_Call{Call: _e.mock.On("RevokeUser", ctx, userID)}
}

func (_c *SessionRepo_RevokeUser_Call) Run(run func(ctx context.Context, userID uint32)) *SessionRepo_RevokeUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *SessionRepo_RevokeUser_Call) Return(keys []string, err error) *SessionRepo_RevokeUser_Call {
	_c.Call.Return(keys, err)
	return _c
}

func (_c *SessionRepo_RevokeUser_Call) RunAndReturn(run func(context.Context, uint32) ([]string, error)) *SessionRepo_RevokeUser_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewSessionRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewSessionRepo creates a new instance of SessionRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSessionRepo(t mockConstructorTestingTNewSessionRepo) *SessionRepo {
	mock := &SessionRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
