// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "job-application/entity/models"
)

// UserUsecase is an autogenerated mock type for the UserUsecase type
type UserUsecase struct {
	mock.Mock
}

// LoginUser provides a mock function with given fields: ctx, email, password
func (_m *UserUsecase) LoginUser(ctx context.Context, email string, password string) (string, error) {
	ret := _m.Called(ctx, email, password)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, email, password)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterUser provides a mock function with given fields: ctx, user, auth
func (_m *UserUsecase) RegisterUser(ctx context.Context, user *models.User, auth *models.Auth) (*models.Auth, error) {
	ret := _m.Called(ctx, user, auth)

	var r0 *models.Auth
	if rf, ok := ret.Get(0).(func(context.Context, *models.User, *models.Auth) *models.Auth); ok {
		r0 = rf(ctx, user, auth)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Auth)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.User, *models.Auth) error); ok {
		r1 = rf(ctx, user, auth)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserUsecase creates a new instance of UserUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserUsecase(t mockConstructorTestingTNewUserUsecase) *UserUsecase {
	mock := &UserUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
