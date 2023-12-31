// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "job-application/entity/models"

	query "job-application/query"
)

// ApplicationUsecase is an autogenerated mock type for the ApplicationUsecase type
type ApplicationUsecase struct {
	mock.Mock
}

// ApplyApplication provides a mock function with given fields: ctx, application
func (_m *ApplicationUsecase) ApplyApplication(ctx context.Context, application *models.Application) (*models.Application, error) {
	ret := _m.Called(ctx, application)

	var r0 *models.Application
	if rf, ok := ret.Get(0).(func(context.Context, *models.Application) *models.Application); ok {
		r0 = rf(ctx, application)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Application)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Application) error); ok {
		r1 = rf(ctx, application)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllApplications provides a mock function with given fields: ctx, clauses
func (_m *ApplicationUsecase) GetAllApplications(ctx context.Context, clauses []query.WhereClause) ([]models.Application, error) {
	ret := _m.Called(ctx, clauses)

	var r0 []models.Application
	if rf, ok := ret.Get(0).(func(context.Context, []query.WhereClause) []models.Application); ok {
		r0 = rf(ctx, clauses)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Application)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []query.WhereClause) error); ok {
		r1 = rf(ctx, clauses)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetApplicationById provides a mock function with given fields: ctx, id
func (_m *ApplicationUsecase) GetApplicationById(ctx context.Context, id uint) (*models.Application, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Application
	if rf, ok := ret.Get(0).(func(context.Context, uint) *models.Application); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Application)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewApplicationUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewApplicationUsecase creates a new instance of ApplicationUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewApplicationUsecase(t mockConstructorTestingTNewApplicationUsecase) *ApplicationUsecase {
	mock := &ApplicationUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
