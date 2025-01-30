// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	domain "github.com/central-university-dev/backend_academy_2024_project_4-go-Matthew11K/internal/domain"
)

// ImageSaver is an autogenerated mock type for the ImageSaver type
type ImageSaver struct {
	mock.Mock
}

// Save provides a mock function with given fields: image, filename, gamma
func (_m *ImageSaver) Save(image *domain.FractalImage, filename string, gamma float64) error {
	ret := _m.Called(image, filename, gamma)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.FractalImage, string, float64) error); ok {
		r0 = rf(image, filename, gamma)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewImageSaver creates a new instance of ImageSaver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewImageSaver(t interface {
	mock.TestingT
	Cleanup(func())
}) *ImageSaver {
	mock := &ImageSaver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
