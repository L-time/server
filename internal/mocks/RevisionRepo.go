// Code generated by mockery v2.13.0-beta.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/bangumi/server/internal/model"
)

// RevisionRepo is an autogenerated mock type for the RevisionRepo type
type RevisionRepo struct {
	mock.Mock
}

type RevisionRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *RevisionRepo) EXPECT() *RevisionRepo_Expecter {
	return &RevisionRepo_Expecter{mock: &_m.Mock}
}

// CountCharacterRelated provides a mock function with given fields: ctx, characterID
func (_m *RevisionRepo) CountCharacterRelated(ctx context.Context, characterID uint32) (int64, error) {
	ret := _m.Called(ctx, characterID)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, uint32) int64); ok {
		r0 = rf(ctx, characterID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, characterID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RevisionRepo_CountCharacterRelated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CountCharacterRelated'
type RevisionRepo_CountCharacterRelated_Call struct {
	*mock.Call
}

// CountCharacterRelated is a helper method to define mock.On call
//  - ctx context.Context
//  - characterID uint32
func (_e *RevisionRepo_Expecter) CountCharacterRelated(ctx interface{}, characterID interface{}) *RevisionRepo_CountCharacterRelated_Call {
	return &RevisionRepo_CountCharacterRelated_Call{Call: _e.mock.On("CountCharacterRelated", ctx, characterID)}
}

func (_c *RevisionRepo_CountCharacterRelated_Call) Run(run func(ctx context.Context, characterID uint32)) *RevisionRepo_CountCharacterRelated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *RevisionRepo_CountCharacterRelated_Call) Return(_a0 int64, _a1 error) *RevisionRepo_CountCharacterRelated_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// CountPersonRelated provides a mock function with given fields: ctx, personID
func (_m *RevisionRepo) CountPersonRelated(ctx context.Context, personID uint32) (int64, error) {
	ret := _m.Called(ctx, personID)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, uint32) int64); ok {
		r0 = rf(ctx, personID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, personID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RevisionRepo_CountPersonRelated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CountPersonRelated'
type RevisionRepo_CountPersonRelated_Call struct {
	*mock.Call
}

// CountPersonRelated is a helper method to define mock.On call
//  - ctx context.Context
//  - personID uint32
func (_e *RevisionRepo_Expecter) CountPersonRelated(ctx interface{}, personID interface{}) *RevisionRepo_CountPersonRelated_Call {
	return &RevisionRepo_CountPersonRelated_Call{Call: _e.mock.On("CountPersonRelated", ctx, personID)}
}

func (_c *RevisionRepo_CountPersonRelated_Call) Run(run func(ctx context.Context, personID uint32)) *RevisionRepo_CountPersonRelated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *RevisionRepo_CountPersonRelated_Call) Return(_a0 int64, _a1 error) *RevisionRepo_CountPersonRelated_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// CountSubjectRelated provides a mock function with given fields: ctx, id
func (_m *RevisionRepo) CountSubjectRelated(ctx context.Context, id uint32) (int64, error) {
	ret := _m.Called(ctx, id)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, uint32) int64); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RevisionRepo_CountSubjectRelated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CountSubjectRelated'
type RevisionRepo_CountSubjectRelated_Call struct {
	*mock.Call
}

// CountSubjectRelated is a helper method to define mock.On call
//  - ctx context.Context
//  - id uint32
func (_e *RevisionRepo_Expecter) CountSubjectRelated(ctx interface{}, id interface{}) *RevisionRepo_CountSubjectRelated_Call {
	return &RevisionRepo_CountSubjectRelated_Call{Call: _e.mock.On("CountSubjectRelated", ctx, id)}
}

func (_c *RevisionRepo_CountSubjectRelated_Call) Run(run func(ctx context.Context, id uint32)) *RevisionRepo_CountSubjectRelated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *RevisionRepo_CountSubjectRelated_Call) Return(_a0 int64, _a1 error) *RevisionRepo_CountSubjectRelated_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetCharacterRelated provides a mock function with given fields: ctx, id
func (_m *RevisionRepo) GetCharacterRelated(ctx context.Context, id uint32) (model.CharacterRevision, error) {
	ret := _m.Called(ctx, id)

	var r0 model.CharacterRevision
	if rf, ok := ret.Get(0).(func(context.Context, uint32) model.CharacterRevision); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.CharacterRevision)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RevisionRepo_GetCharacterRelated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCharacterRelated'
type RevisionRepo_GetCharacterRelated_Call struct {
	*mock.Call
}

// GetCharacterRelated is a helper method to define mock.On call
//  - ctx context.Context
//  - id uint32
func (_e *RevisionRepo_Expecter) GetCharacterRelated(ctx interface{}, id interface{}) *RevisionRepo_GetCharacterRelated_Call {
	return &RevisionRepo_GetCharacterRelated_Call{Call: _e.mock.On("GetCharacterRelated", ctx, id)}
}

func (_c *RevisionRepo_GetCharacterRelated_Call) Run(run func(ctx context.Context, id uint32)) *RevisionRepo_GetCharacterRelated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *RevisionRepo_GetCharacterRelated_Call) Return(_a0 model.CharacterRevision, _a1 error) *RevisionRepo_GetCharacterRelated_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetPersonRelated provides a mock function with given fields: ctx, id
func (_m *RevisionRepo) GetPersonRelated(ctx context.Context, id uint32) (model.PersonRevision, error) {
	ret := _m.Called(ctx, id)

	var r0 model.PersonRevision
	if rf, ok := ret.Get(0).(func(context.Context, uint32) model.PersonRevision); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.PersonRevision)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RevisionRepo_GetPersonRelated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPersonRelated'
type RevisionRepo_GetPersonRelated_Call struct {
	*mock.Call
}

// GetPersonRelated is a helper method to define mock.On call
//  - ctx context.Context
//  - id uint32
func (_e *RevisionRepo_Expecter) GetPersonRelated(ctx interface{}, id interface{}) *RevisionRepo_GetPersonRelated_Call {
	return &RevisionRepo_GetPersonRelated_Call{Call: _e.mock.On("GetPersonRelated", ctx, id)}
}

func (_c *RevisionRepo_GetPersonRelated_Call) Run(run func(ctx context.Context, id uint32)) *RevisionRepo_GetPersonRelated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *RevisionRepo_GetPersonRelated_Call) Return(_a0 model.PersonRevision, _a1 error) *RevisionRepo_GetPersonRelated_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetSubjectRelated provides a mock function with given fields: ctx, id
func (_m *RevisionRepo) GetSubjectRelated(ctx context.Context, id uint32) (model.SubjectRevision, error) {
	ret := _m.Called(ctx, id)

	var r0 model.SubjectRevision
	if rf, ok := ret.Get(0).(func(context.Context, uint32) model.SubjectRevision); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.SubjectRevision)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RevisionRepo_GetSubjectRelated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSubjectRelated'
type RevisionRepo_GetSubjectRelated_Call struct {
	*mock.Call
}

// GetSubjectRelated is a helper method to define mock.On call
//  - ctx context.Context
//  - id uint32
func (_e *RevisionRepo_Expecter) GetSubjectRelated(ctx interface{}, id interface{}) *RevisionRepo_GetSubjectRelated_Call {
	return &RevisionRepo_GetSubjectRelated_Call{Call: _e.mock.On("GetSubjectRelated", ctx, id)}
}

func (_c *RevisionRepo_GetSubjectRelated_Call) Run(run func(ctx context.Context, id uint32)) *RevisionRepo_GetSubjectRelated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *RevisionRepo_GetSubjectRelated_Call) Return(_a0 model.SubjectRevision, _a1 error) *RevisionRepo_GetSubjectRelated_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ListCharacterRelated provides a mock function with given fields: ctx, characterID, limit, offset
func (_m *RevisionRepo) ListCharacterRelated(ctx context.Context, characterID uint32, limit int, offset int) ([]model.CharacterRevision, error) {
	ret := _m.Called(ctx, characterID, limit, offset)

	var r0 []model.CharacterRevision
	if rf, ok := ret.Get(0).(func(context.Context, uint32, int, int) []model.CharacterRevision); ok {
		r0 = rf(ctx, characterID, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.CharacterRevision)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32, int, int) error); ok {
		r1 = rf(ctx, characterID, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RevisionRepo_ListCharacterRelated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListCharacterRelated'
type RevisionRepo_ListCharacterRelated_Call struct {
	*mock.Call
}

// ListCharacterRelated is a helper method to define mock.On call
//  - ctx context.Context
//  - characterID uint32
//  - limit int
//  - offset int
func (_e *RevisionRepo_Expecter) ListCharacterRelated(ctx interface{}, characterID interface{}, limit interface{}, offset interface{}) *RevisionRepo_ListCharacterRelated_Call {
	return &RevisionRepo_ListCharacterRelated_Call{Call: _e.mock.On("ListCharacterRelated", ctx, characterID, limit, offset)}
}

func (_c *RevisionRepo_ListCharacterRelated_Call) Run(run func(ctx context.Context, characterID uint32, limit int, offset int)) *RevisionRepo_ListCharacterRelated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(int), args[3].(int))
	})
	return _c
}

func (_c *RevisionRepo_ListCharacterRelated_Call) Return(_a0 []model.CharacterRevision, _a1 error) *RevisionRepo_ListCharacterRelated_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ListPersonRelated provides a mock function with given fields: ctx, personID, limit, offset
func (_m *RevisionRepo) ListPersonRelated(ctx context.Context, personID uint32, limit int, offset int) ([]model.PersonRevision, error) {
	ret := _m.Called(ctx, personID, limit, offset)

	var r0 []model.PersonRevision
	if rf, ok := ret.Get(0).(func(context.Context, uint32, int, int) []model.PersonRevision); ok {
		r0 = rf(ctx, personID, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.PersonRevision)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32, int, int) error); ok {
		r1 = rf(ctx, personID, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RevisionRepo_ListPersonRelated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListPersonRelated'
type RevisionRepo_ListPersonRelated_Call struct {
	*mock.Call
}

// ListPersonRelated is a helper method to define mock.On call
//  - ctx context.Context
//  - personID uint32
//  - limit int
//  - offset int
func (_e *RevisionRepo_Expecter) ListPersonRelated(ctx interface{}, personID interface{}, limit interface{}, offset interface{}) *RevisionRepo_ListPersonRelated_Call {
	return &RevisionRepo_ListPersonRelated_Call{Call: _e.mock.On("ListPersonRelated", ctx, personID, limit, offset)}
}

func (_c *RevisionRepo_ListPersonRelated_Call) Run(run func(ctx context.Context, personID uint32, limit int, offset int)) *RevisionRepo_ListPersonRelated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(int), args[3].(int))
	})
	return _c
}

func (_c *RevisionRepo_ListPersonRelated_Call) Return(_a0 []model.PersonRevision, _a1 error) *RevisionRepo_ListPersonRelated_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ListSubjectRelated provides a mock function with given fields: ctx, id, limit, offset
func (_m *RevisionRepo) ListSubjectRelated(ctx context.Context, id uint32, limit int, offset int) ([]model.SubjectRevision, error) {
	ret := _m.Called(ctx, id, limit, offset)

	var r0 []model.SubjectRevision
	if rf, ok := ret.Get(0).(func(context.Context, uint32, int, int) []model.SubjectRevision); ok {
		r0 = rf(ctx, id, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.SubjectRevision)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32, int, int) error); ok {
		r1 = rf(ctx, id, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RevisionRepo_ListSubjectRelated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSubjectRelated'
type RevisionRepo_ListSubjectRelated_Call struct {
	*mock.Call
}

// ListSubjectRelated is a helper method to define mock.On call
//  - ctx context.Context
//  - id uint32
//  - limit int
//  - offset int
func (_e *RevisionRepo_Expecter) ListSubjectRelated(ctx interface{}, id interface{}, limit interface{}, offset interface{}) *RevisionRepo_ListSubjectRelated_Call {
	return &RevisionRepo_ListSubjectRelated_Call{Call: _e.mock.On("ListSubjectRelated", ctx, id, limit, offset)}
}

func (_c *RevisionRepo_ListSubjectRelated_Call) Run(run func(ctx context.Context, id uint32, limit int, offset int)) *RevisionRepo_ListSubjectRelated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(int), args[3].(int))
	})
	return _c
}

func (_c *RevisionRepo_ListSubjectRelated_Call) Return(_a0 []model.SubjectRevision, _a1 error) *RevisionRepo_ListSubjectRelated_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type NewRevisionRepoT interface {
	mock.TestingT
	Cleanup(func())
}

// NewRevisionRepo creates a new instance of RevisionRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRevisionRepo(t NewRevisionRepoT) *RevisionRepo {
	mock := &RevisionRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
