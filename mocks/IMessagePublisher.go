// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	v1 "auth-service/internal/gen/otp/v1"

	mock "github.com/stretchr/testify/mock"
)

// IMessagePublisher is an autogenerated mock type for the IMessagePublisher type
type IMessagePublisher struct {
	mock.Mock
}

// Publish provides a mock function with given fields: request
func (_m *IMessagePublisher) Publish(request *v1.GenerateOTPRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1.GenerateOTPRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIMessagePublisher creates a new instance of IMessagePublisher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIMessagePublisher(t interface {
	mock.TestingT
	Cleanup(func())
}) *IMessagePublisher {
	mock := &IMessagePublisher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}