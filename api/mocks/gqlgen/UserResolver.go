// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	pg "github.com/leggettc18/grindlists/api/pg"
)

// UserResolver is an autogenerated mock type for the UserResolver type
type UserResolver struct {
	mock.Mock
}

// Lists provides a mock function with given fields: ctx, obj
func (_m *UserResolver) Lists(ctx context.Context, obj *pg.User) ([]pg.List, error) {
	ret := _m.Called(ctx, obj)

	var r0 []pg.List
	if rf, ok := ret.Get(0).(func(context.Context, *pg.User) []pg.List); ok {
		r0 = rf(ctx, obj)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pg.List)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pg.User) error); ok {
		r1 = rf(ctx, obj)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}