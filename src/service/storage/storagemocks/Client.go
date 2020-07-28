// Code generated by mockery v1.0.0. DO NOT EDIT.

package storagemocks

import (
	context "context"
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Name provides a mock function with given fields:
func (_m *Client) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Put provides a mock function with given fields: ctx, r, key
func (_m *Client) Put(ctx context.Context, r io.Reader, key string) error {
	ret := _m.Called(ctx, r, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, io.Reader, string) error); ok {
		r0 = rf(ctx, r, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
