// Code generated by mockery v2.32.0. DO NOT EDIT.

package ports

import mock "github.com/stretchr/testify/mock"

// MockUserQueryRepository is an autogenerated mock type for the UserQueryRepository type
type MockUserQueryRepository struct {
	mock.Mock
}

// Login provides a mock function with given fields: email, password
func (_m *MockUserQueryRepository) Login(email string, password string) error {
	ret := _m.Called(email, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(email, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockUserQueryRepository creates a new instance of MockUserQueryRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserQueryRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserQueryRepository {
	mock := &MockUserQueryRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
