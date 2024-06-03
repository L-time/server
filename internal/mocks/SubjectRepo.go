// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/bangumi/server/domain"
	mock "github.com/stretchr/testify/mock"

	model "github.com/bangumi/server/internal/model"

	subject "github.com/bangumi/server/internal/subject"
)

// SubjectRepo is an autogenerated mock type for the Repo type
type SubjectRepo struct {
	mock.Mock
}

type SubjectRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *SubjectRepo) EXPECT() *SubjectRepo_Expecter {
	return &SubjectRepo_Expecter{mock: &_m.Mock}
}

// Browse provides a mock function with given fields: ctx, filter, limit, offset
func (_m *SubjectRepo) Browse(ctx context.Context, filter subject.BrowseFilter, limit int, offset int) ([]model.Subject, error) {
	ret := _m.Called(ctx, filter, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for Browse")
	}

	var r0 []model.Subject
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, subject.BrowseFilter, int, int) ([]model.Subject, error)); ok {
		return rf(ctx, filter, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(context.Context, subject.BrowseFilter, int, int) []model.Subject); ok {
		r0 = rf(ctx, filter, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Subject)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, subject.BrowseFilter, int, int) error); ok {
		r1 = rf(ctx, filter, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubjectRepo_Browse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Browse'
type SubjectRepo_Browse_Call struct {
	*mock.Call
}

// Browse is a helper method to define mock.On call
//   - ctx context.Context
//   - filter subject.BrowseFilter
//   - limit int
//   - offset int
func (_e *SubjectRepo_Expecter) Browse(ctx interface{}, filter interface{}, limit interface{}, offset interface{}) *SubjectRepo_Browse_Call {
	return &SubjectRepo_Browse_Call{Call: _e.mock.On("Browse", ctx, filter, limit, offset)}
}

func (_c *SubjectRepo_Browse_Call) Run(run func(ctx context.Context, filter subject.BrowseFilter, limit int, offset int)) *SubjectRepo_Browse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(subject.BrowseFilter), args[2].(int), args[3].(int))
	})
	return _c
}

func (_c *SubjectRepo_Browse_Call) Return(_a0 []model.Subject, _a1 error) *SubjectRepo_Browse_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SubjectRepo_Browse_Call) RunAndReturn(run func(context.Context, subject.BrowseFilter, int, int) ([]model.Subject, error)) *SubjectRepo_Browse_Call {
	_c.Call.Return(run)
	return _c
}

// Count provides a mock function with given fields: ctx, filter
func (_m *SubjectRepo) Count(ctx context.Context, filter subject.BrowseFilter) (int64, error) {
	ret := _m.Called(ctx, filter)

	if len(ret) == 0 {
		panic("no return value specified for Count")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, subject.BrowseFilter) (int64, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, subject.BrowseFilter) int64); ok {
		r0 = rf(ctx, filter)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, subject.BrowseFilter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubjectRepo_Count_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Count'
type SubjectRepo_Count_Call struct {
	*mock.Call
}

// Count is a helper method to define mock.On call
//   - ctx context.Context
//   - filter subject.BrowseFilter
func (_e *SubjectRepo_Expecter) Count(ctx interface{}, filter interface{}) *SubjectRepo_Count_Call {
	return &SubjectRepo_Count_Call{Call: _e.mock.On("Count", ctx, filter)}
}

func (_c *SubjectRepo_Count_Call) Run(run func(ctx context.Context, filter subject.BrowseFilter)) *SubjectRepo_Count_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(subject.BrowseFilter))
	})
	return _c
}

func (_c *SubjectRepo_Count_Call) Return(_a0 int64, _a1 error) *SubjectRepo_Count_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SubjectRepo_Count_Call) RunAndReturn(run func(context.Context, subject.BrowseFilter) (int64, error)) *SubjectRepo_Count_Call {
	_c.Call.Return(run)
	return _c
}

// DeletePost provides a mock function with given fields: ctx, id
func (_m *SubjectRepo) DeletePost(ctx context.Context, id uint32) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeletePost")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubjectRepo_DeletePost_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeletePost'
type SubjectRepo_DeletePost_Call struct {
	*mock.Call
}

// DeletePost is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint32
func (_e *SubjectRepo_Expecter) DeletePost(ctx interface{}, id interface{}) *SubjectRepo_DeletePost_Call {
	return &SubjectRepo_DeletePost_Call{Call: _e.mock.On("DeletePost", ctx, id)}
}

func (_c *SubjectRepo_DeletePost_Call) Run(run func(ctx context.Context, id uint32)) *SubjectRepo_DeletePost_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *SubjectRepo_DeletePost_Call) Return(_a0 error) *SubjectRepo_DeletePost_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SubjectRepo_DeletePost_Call) RunAndReturn(run func(context.Context, uint32) error) *SubjectRepo_DeletePost_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, id, filter
func (_m *SubjectRepo) Get(ctx context.Context, id uint32, filter subject.Filter) (model.Subject, error) {
	ret := _m.Called(ctx, id, filter)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 model.Subject
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, subject.Filter) (model.Subject, error)); ok {
		return rf(ctx, id, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32, subject.Filter) model.Subject); ok {
		r0 = rf(ctx, id, filter)
	} else {
		r0 = ret.Get(0).(model.Subject)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32, subject.Filter) error); ok {
		r1 = rf(ctx, id, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubjectRepo_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type SubjectRepo_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint32
//   - filter subject.Filter
func (_e *SubjectRepo_Expecter) Get(ctx interface{}, id interface{}, filter interface{}) *SubjectRepo_Get_Call {
	return &SubjectRepo_Get_Call{Call: _e.mock.On("Get", ctx, id, filter)}
}

func (_c *SubjectRepo_Get_Call) Run(run func(ctx context.Context, id uint32, filter subject.Filter)) *SubjectRepo_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(subject.Filter))
	})
	return _c
}

func (_c *SubjectRepo_Get_Call) Return(_a0 model.Subject, _a1 error) *SubjectRepo_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SubjectRepo_Get_Call) RunAndReturn(run func(context.Context, uint32, subject.Filter) (model.Subject, error)) *SubjectRepo_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetActors provides a mock function with given fields: ctx, subjectID, characterIDs
func (_m *SubjectRepo) GetActors(ctx context.Context, subjectID uint32, characterIDs []uint32) (map[uint32][]uint32, error) {
	ret := _m.Called(ctx, subjectID, characterIDs)

	if len(ret) == 0 {
		panic("no return value specified for GetActors")
	}

	var r0 map[uint32][]uint32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, []uint32) (map[uint32][]uint32, error)); ok {
		return rf(ctx, subjectID, characterIDs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32, []uint32) map[uint32][]uint32); ok {
		r0 = rf(ctx, subjectID, characterIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[uint32][]uint32)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32, []uint32) error); ok {
		r1 = rf(ctx, subjectID, characterIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubjectRepo_GetActors_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetActors'
type SubjectRepo_GetActors_Call struct {
	*mock.Call
}

// GetActors is a helper method to define mock.On call
//   - ctx context.Context
//   - subjectID uint32
//   - characterIDs []uint32
func (_e *SubjectRepo_Expecter) GetActors(ctx interface{}, subjectID interface{}, characterIDs interface{}) *SubjectRepo_GetActors_Call {
	return &SubjectRepo_GetActors_Call{Call: _e.mock.On("GetActors", ctx, subjectID, characterIDs)}
}

func (_c *SubjectRepo_GetActors_Call) Run(run func(ctx context.Context, subjectID uint32, characterIDs []uint32)) *SubjectRepo_GetActors_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].([]uint32))
	})
	return _c
}

func (_c *SubjectRepo_GetActors_Call) Return(_a0 map[uint32][]uint32, _a1 error) *SubjectRepo_GetActors_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SubjectRepo_GetActors_Call) RunAndReturn(run func(context.Context, uint32, []uint32) (map[uint32][]uint32, error)) *SubjectRepo_GetActors_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllPost provides a mock function with given fields: ctx, id, offset, limit
func (_m *SubjectRepo) GetAllPost(ctx context.Context, id uint32, offset int, limit int) ([]model.SubjectPost, error) {
	ret := _m.Called(ctx, id, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for GetAllPost")
	}

	var r0 []model.SubjectPost
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, int, int) ([]model.SubjectPost, error)); ok {
		return rf(ctx, id, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32, int, int) []model.SubjectPost); ok {
		r0 = rf(ctx, id, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.SubjectPost)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32, int, int) error); ok {
		r1 = rf(ctx, id, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubjectRepo_GetAllPost_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllPost'
type SubjectRepo_GetAllPost_Call struct {
	*mock.Call
}

// GetAllPost is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint32
//   - offset int
//   - limit int
func (_e *SubjectRepo_Expecter) GetAllPost(ctx interface{}, id interface{}, offset interface{}, limit interface{}) *SubjectRepo_GetAllPost_Call {
	return &SubjectRepo_GetAllPost_Call{Call: _e.mock.On("GetAllPost", ctx, id, offset, limit)}
}

func (_c *SubjectRepo_GetAllPost_Call) Run(run func(ctx context.Context, id uint32, offset int, limit int)) *SubjectRepo_GetAllPost_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(int), args[3].(int))
	})
	return _c
}

func (_c *SubjectRepo_GetAllPost_Call) Return(_a0 []model.SubjectPost, _a1 error) *SubjectRepo_GetAllPost_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SubjectRepo_GetAllPost_Call) RunAndReturn(run func(context.Context, uint32, int, int) ([]model.SubjectPost, error)) *SubjectRepo_GetAllPost_Call {
	_c.Call.Return(run)
	return _c
}

// GetByIDs provides a mock function with given fields: ctx, ids, filter
func (_m *SubjectRepo) GetByIDs(ctx context.Context, ids []uint32, filter subject.Filter) (map[uint32]model.Subject, error) {
	ret := _m.Called(ctx, ids, filter)

	if len(ret) == 0 {
		panic("no return value specified for GetByIDs")
	}

	var r0 map[uint32]model.Subject
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []uint32, subject.Filter) (map[uint32]model.Subject, error)); ok {
		return rf(ctx, ids, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []uint32, subject.Filter) map[uint32]model.Subject); ok {
		r0 = rf(ctx, ids, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[uint32]model.Subject)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []uint32, subject.Filter) error); ok {
		r1 = rf(ctx, ids, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubjectRepo_GetByIDs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByIDs'
type SubjectRepo_GetByIDs_Call struct {
	*mock.Call
}

// GetByIDs is a helper method to define mock.On call
//   - ctx context.Context
//   - ids []uint32
//   - filter subject.Filter
func (_e *SubjectRepo_Expecter) GetByIDs(ctx interface{}, ids interface{}, filter interface{}) *SubjectRepo_GetByIDs_Call {
	return &SubjectRepo_GetByIDs_Call{Call: _e.mock.On("GetByIDs", ctx, ids, filter)}
}

func (_c *SubjectRepo_GetByIDs_Call) Run(run func(ctx context.Context, ids []uint32, filter subject.Filter)) *SubjectRepo_GetByIDs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]uint32), args[2].(subject.Filter))
	})
	return _c
}

func (_c *SubjectRepo_GetByIDs_Call) Return(_a0 map[uint32]model.Subject, _a1 error) *SubjectRepo_GetByIDs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SubjectRepo_GetByIDs_Call) RunAndReturn(run func(context.Context, []uint32, subject.Filter) (map[uint32]model.Subject, error)) *SubjectRepo_GetByIDs_Call {
	_c.Call.Return(run)
	return _c
}

// GetCharacterRelated provides a mock function with given fields: ctx, characterID
func (_m *SubjectRepo) GetCharacterRelated(ctx context.Context, characterID uint32) ([]domain.SubjectCharacterRelation, error) {
	ret := _m.Called(ctx, characterID)

	if len(ret) == 0 {
		panic("no return value specified for GetCharacterRelated")
	}

	var r0 []domain.SubjectCharacterRelation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32) ([]domain.SubjectCharacterRelation, error)); ok {
		return rf(ctx, characterID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32) []domain.SubjectCharacterRelation); ok {
		r0 = rf(ctx, characterID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.SubjectCharacterRelation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, characterID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubjectRepo_GetCharacterRelated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCharacterRelated'
type SubjectRepo_GetCharacterRelated_Call struct {
	*mock.Call
}

// GetCharacterRelated is a helper method to define mock.On call
//   - ctx context.Context
//   - characterID uint32
func (_e *SubjectRepo_Expecter) GetCharacterRelated(ctx interface{}, characterID interface{}) *SubjectRepo_GetCharacterRelated_Call {
	return &SubjectRepo_GetCharacterRelated_Call{Call: _e.mock.On("GetCharacterRelated", ctx, characterID)}
}

func (_c *SubjectRepo_GetCharacterRelated_Call) Run(run func(ctx context.Context, characterID uint32)) *SubjectRepo_GetCharacterRelated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *SubjectRepo_GetCharacterRelated_Call) Return(_a0 []domain.SubjectCharacterRelation, _a1 error) *SubjectRepo_GetCharacterRelated_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SubjectRepo_GetCharacterRelated_Call) RunAndReturn(run func(context.Context, uint32) ([]domain.SubjectCharacterRelation, error)) *SubjectRepo_GetCharacterRelated_Call {
	_c.Call.Return(run)
	return _c
}

// GetPersonRelated provides a mock function with given fields: ctx, personID
func (_m *SubjectRepo) GetPersonRelated(ctx context.Context, personID uint32) ([]domain.SubjectPersonRelation, error) {
	ret := _m.Called(ctx, personID)

	if len(ret) == 0 {
		panic("no return value specified for GetPersonRelated")
	}

	var r0 []domain.SubjectPersonRelation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32) ([]domain.SubjectPersonRelation, error)); ok {
		return rf(ctx, personID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32) []domain.SubjectPersonRelation); ok {
		r0 = rf(ctx, personID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.SubjectPersonRelation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, personID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubjectRepo_GetPersonRelated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPersonRelated'
type SubjectRepo_GetPersonRelated_Call struct {
	*mock.Call
}

// GetPersonRelated is a helper method to define mock.On call
//   - ctx context.Context
//   - personID uint32
func (_e *SubjectRepo_Expecter) GetPersonRelated(ctx interface{}, personID interface{}) *SubjectRepo_GetPersonRelated_Call {
	return &SubjectRepo_GetPersonRelated_Call{Call: _e.mock.On("GetPersonRelated", ctx, personID)}
}

func (_c *SubjectRepo_GetPersonRelated_Call) Run(run func(ctx context.Context, personID uint32)) *SubjectRepo_GetPersonRelated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *SubjectRepo_GetPersonRelated_Call) Return(_a0 []domain.SubjectPersonRelation, _a1 error) *SubjectRepo_GetPersonRelated_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SubjectRepo_GetPersonRelated_Call) RunAndReturn(run func(context.Context, uint32) ([]domain.SubjectPersonRelation, error)) *SubjectRepo_GetPersonRelated_Call {
	_c.Call.Return(run)
	return _c
}

// GetPost provides a mock function with given fields: ctx, id
func (_m *SubjectRepo) GetPost(ctx context.Context, id uint32) (model.SubjectPost, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetPost")
	}

	var r0 model.SubjectPost
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32) (model.SubjectPost, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32) model.SubjectPost); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.SubjectPost)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubjectRepo_GetPost_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPost'
type SubjectRepo_GetPost_Call struct {
	*mock.Call
}

// GetPost is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint32
func (_e *SubjectRepo_Expecter) GetPost(ctx interface{}, id interface{}) *SubjectRepo_GetPost_Call {
	return &SubjectRepo_GetPost_Call{Call: _e.mock.On("GetPost", ctx, id)}
}

func (_c *SubjectRepo_GetPost_Call) Run(run func(ctx context.Context, id uint32)) *SubjectRepo_GetPost_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *SubjectRepo_GetPost_Call) Return(_a0 model.SubjectPost, _a1 error) *SubjectRepo_GetPost_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SubjectRepo_GetPost_Call) RunAndReturn(run func(context.Context, uint32) (model.SubjectPost, error)) *SubjectRepo_GetPost_Call {
	_c.Call.Return(run)
	return _c
}

// GetSubjectRelated provides a mock function with given fields: ctx, subjectID
func (_m *SubjectRepo) GetSubjectRelated(ctx context.Context, subjectID uint32) ([]domain.SubjectInternalRelation, error) {
	ret := _m.Called(ctx, subjectID)

	if len(ret) == 0 {
		panic("no return value specified for GetSubjectRelated")
	}

	var r0 []domain.SubjectInternalRelation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32) ([]domain.SubjectInternalRelation, error)); ok {
		return rf(ctx, subjectID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32) []domain.SubjectInternalRelation); ok {
		r0 = rf(ctx, subjectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.SubjectInternalRelation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, subjectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubjectRepo_GetSubjectRelated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSubjectRelated'
type SubjectRepo_GetSubjectRelated_Call struct {
	*mock.Call
}

// GetSubjectRelated is a helper method to define mock.On call
//   - ctx context.Context
//   - subjectID uint32
func (_e *SubjectRepo_Expecter) GetSubjectRelated(ctx interface{}, subjectID interface{}) *SubjectRepo_GetSubjectRelated_Call {
	return &SubjectRepo_GetSubjectRelated_Call{Call: _e.mock.On("GetSubjectRelated", ctx, subjectID)}
}

func (_c *SubjectRepo_GetSubjectRelated_Call) Run(run func(ctx context.Context, subjectID uint32)) *SubjectRepo_GetSubjectRelated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *SubjectRepo_GetSubjectRelated_Call) Return(_a0 []domain.SubjectInternalRelation, _a1 error) *SubjectRepo_GetSubjectRelated_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SubjectRepo_GetSubjectRelated_Call) RunAndReturn(run func(context.Context, uint32) ([]domain.SubjectInternalRelation, error)) *SubjectRepo_GetSubjectRelated_Call {
	_c.Call.Return(run)
	return _c
}

// NewPost provides a mock function with given fields: ctx, post
func (_m *SubjectRepo) NewPost(ctx context.Context, post model.SubjectPost) error {
	ret := _m.Called(ctx, post)

	if len(ret) == 0 {
		panic("no return value specified for NewPost")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.SubjectPost) error); ok {
		r0 = rf(ctx, post)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubjectRepo_NewPost_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewPost'
type SubjectRepo_NewPost_Call struct {
	*mock.Call
}

// NewPost is a helper method to define mock.On call
//   - ctx context.Context
//   - post model.SubjectPost
func (_e *SubjectRepo_Expecter) NewPost(ctx interface{}, post interface{}) *SubjectRepo_NewPost_Call {
	return &SubjectRepo_NewPost_Call{Call: _e.mock.On("NewPost", ctx, post)}
}

func (_c *SubjectRepo_NewPost_Call) Run(run func(ctx context.Context, post model.SubjectPost)) *SubjectRepo_NewPost_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.SubjectPost))
	})
	return _c
}

func (_c *SubjectRepo_NewPost_Call) Return(_a0 error) *SubjectRepo_NewPost_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SubjectRepo_NewPost_Call) RunAndReturn(run func(context.Context, model.SubjectPost) error) *SubjectRepo_NewPost_Call {
	_c.Call.Return(run)
	return _c
}

// NewSubjectRepo creates a new instance of SubjectRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSubjectRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *SubjectRepo {
	mock := &SubjectRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
