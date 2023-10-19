// Code generated by mockery v2.32.0. DO NOT EDIT.

package ports

import mock "github.com/stretchr/testify/mock"

// MockUserService is an autogenerated mock type for the UserService type
type MockUserService struct {
	mock.Mock
}

// Login provides a mock function with given fields: email, password
func (_m *MockUserService) Login(email string, password string) error {
	ret := _m.Called(email, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(email, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Register provides a mock function with given fields: email, password, passConfirm
func (_m *MockUserService) Register(email string, password string, passConfirm string) error {
	ret := _m.Called(email, password, passConfirm)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(email, password, passConfirm)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockUserService creates a new instance of MockUserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserService {
	mock := &MockUserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}