// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "flamingo.me/flamingo/v3/core/security/domain"
	mock "github.com/stretchr/testify/mock"

	web "flamingo.me/flamingo/v3/framework/web"
)

// Provider is an autogenerated mock type for the Provider type
type Provider struct {
	mock.Mock
}

// All provides a mock function with given fields: _a0, _a1
func (_m *Provider) All(_a0 context.Context, _a1 *web.Session) []domain.Role {
	ret := _m.Called(_a0, _a1)

	var r0 []domain.Role
	if rf, ok := ret.Get(0).(func(context.Context, *web.Session) []domain.Role); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Role)
		}
	}

	return r0
}
