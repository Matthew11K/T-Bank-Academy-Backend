// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// InputReader is an autogenerated mock type for the InputReader type
type InputReader struct {
	mock.Mock
}

type InputReader_Expecter struct {
	mock *mock.Mock
}

func (_m *InputReader) EXPECT() *InputReader_Expecter {
	return &InputReader_Expecter{mock: &_m.Mock}
}

// ReadInput provides a mock function with given fields:
func (_m *InputReader) ReadInput() (string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ReadInput")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InputReader_ReadInput_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadInput'
type InputReader_ReadInput_Call struct {
	*mock.Call
}

// ReadInput is a helper method to define mock.On call
func (_e *InputReader_Expecter) ReadInput() *InputReader_ReadInput_Call {
	return &InputReader_ReadInput_Call{Call: _e.mock.On("ReadInput")}
}

func (_c *InputReader_ReadInput_Call) Run(run func()) *InputReader_ReadInput_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *InputReader_ReadInput_Call) Return(_a0 string, _a1 error) *InputReader_ReadInput_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *InputReader_ReadInput_Call) RunAndReturn(run func() (string, error)) *InputReader_ReadInput_Call {
	_c.Call.Return(run)
	return _c
}

// NewInputReader creates a new instance of InputReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewInputReader(t interface {
	mock.TestingT
	Cleanup(func())
}) *InputReader {
	mock := &InputReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
