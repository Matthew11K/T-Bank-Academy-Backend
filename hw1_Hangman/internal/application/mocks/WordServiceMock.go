// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	domain "github.com/es-debug/backend-academy-2024-go-template/internal/domain"
)

// WordServiceInterface is an autogenerated mock type for the WordServiceInterface type
type WordServiceInterface struct {
	mock.Mock
}

type WordServiceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *WordServiceInterface) EXPECT() *WordServiceInterface_Expecter {
	return &WordServiceInterface_Expecter{mock: &_m.Mock}
}

// GetHint provides a mock function with given fields: word
func (_m *WordServiceInterface) GetHint(word *domain.Word) string {
	ret := _m.Called(word)

	if len(ret) == 0 {
		panic("no return value specified for GetHint")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(*domain.Word) string); ok {
		r0 = rf(word)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// WordServiceInterface_GetHint_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetHint'
type WordServiceInterface_GetHint_Call struct {
	*mock.Call
}

// GetHint is a helper method to define mock.On call
//   - word *domain.Word
func (_e *WordServiceInterface_Expecter) GetHint(word interface{}) *WordServiceInterface_GetHint_Call {
	return &WordServiceInterface_GetHint_Call{Call: _e.mock.On("GetHint", word)}
}

func (_c *WordServiceInterface_GetHint_Call) Run(run func(word *domain.Word)) *WordServiceInterface_GetHint_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*domain.Word))
	})
	return _c
}

func (_c *WordServiceInterface_GetHint_Call) Return(_a0 string) *WordServiceInterface_GetHint_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *WordServiceInterface_GetHint_Call) RunAndReturn(run func(*domain.Word) string) *WordServiceInterface_GetHint_Call {
	_c.Call.Return(run)
	return _c
}

// GetRandomWord provides a mock function with given fields: category, difficulty
func (_m *WordServiceInterface) GetRandomWord(category string, difficulty string) (*domain.Word, error) {
	ret := _m.Called(category, difficulty)

	if len(ret) == 0 {
		panic("no return value specified for GetRandomWord")
	}

	var r0 *domain.Word
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*domain.Word, error)); ok {
		return rf(category, difficulty)
	}
	if rf, ok := ret.Get(0).(func(string, string) *domain.Word); ok {
		r0 = rf(category, difficulty)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Word)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(category, difficulty)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WordServiceInterface_GetRandomWord_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRandomWord'
type WordServiceInterface_GetRandomWord_Call struct {
	*mock.Call
}

// GetRandomWord is a helper method to define mock.On call
//   - category string
//   - difficulty string
func (_e *WordServiceInterface_Expecter) GetRandomWord(category interface{}, difficulty interface{}) *WordServiceInterface_GetRandomWord_Call {
	return &WordServiceInterface_GetRandomWord_Call{Call: _e.mock.On("GetRandomWord", category, difficulty)}
}

func (_c *WordServiceInterface_GetRandomWord_Call) Run(run func(category string, difficulty string)) *WordServiceInterface_GetRandomWord_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *WordServiceInterface_GetRandomWord_Call) Return(_a0 *domain.Word, _a1 error) *WordServiceInterface_GetRandomWord_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *WordServiceInterface_GetRandomWord_Call) RunAndReturn(run func(string, string) (*domain.Word, error)) *WordServiceInterface_GetRandomWord_Call {
	_c.Call.Return(run)
	return _c
}

// NewWordServiceInterface creates a new instance of WordServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWordServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *WordServiceInterface {
	mock := &WordServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
