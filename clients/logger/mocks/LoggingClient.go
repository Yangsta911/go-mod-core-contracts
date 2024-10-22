// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	errors "github.com/edgexfoundry/go-mod-core-contracts/v4/errors"

	mock "github.com/stretchr/testify/mock"
)

// LoggingClient is an autogenerated mock type for the LoggingClient type
type LoggingClient struct {
	mock.Mock
}

// Debug provides a mock function with given fields: msg, args
func (_m *LoggingClient) Debug(msg string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Debugf provides a mock function with given fields: msg, args
func (_m *LoggingClient) Debugf(msg string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Error provides a mock function with given fields: msg, args
func (_m *LoggingClient) Error(msg string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Errorf provides a mock function with given fields: msg, args
func (_m *LoggingClient) Errorf(msg string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Info provides a mock function with given fields: msg, args
func (_m *LoggingClient) Info(msg string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Infof provides a mock function with given fields: msg, args
func (_m *LoggingClient) Infof(msg string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// LogLevel provides a mock function with given fields:
func (_m *LoggingClient) LogLevel() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// SetLogLevel provides a mock function with given fields: logLevel
func (_m *LoggingClient) SetLogLevel(logLevel string) errors.EdgeX {
	ret := _m.Called(logLevel)

	var r0 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string) errors.EdgeX); ok {
		r0 = rf(logLevel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.EdgeX)
		}
	}

	return r0
}

// Trace provides a mock function with given fields: msg, args
func (_m *LoggingClient) Trace(msg string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Tracef provides a mock function with given fields: msg, args
func (_m *LoggingClient) Tracef(msg string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Warn provides a mock function with given fields: msg, args
func (_m *LoggingClient) Warn(msg string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Warnf provides a mock function with given fields: msg, args
func (_m *LoggingClient) Warnf(msg string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

type mockConstructorTestingTNewLoggingClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewLoggingClient creates a new instance of LoggingClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLoggingClient(t mockConstructorTestingTNewLoggingClient) *LoggingClient {
	mock := &LoggingClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
