// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package main

import mock "github.com/stretchr/testify/mock"

// mockUploader is an autogenerated mock type for the uploader type
type mockUploader struct {
	mock.Mock
}

// upload provides a mock function with given fields:
func (_m *mockUploader) upload() (string, string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0, ok = ret.Get(0).(string)
		if !ok {
			r0 = ""
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func() string); ok {
		r1 = rf()
	} else {
		r1, ok = ret.Get(1).(string)
		if !ok {
			r1 = ""
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}