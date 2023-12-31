// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "job-application/entity/models"

	query "job-application/query"
)

// JobRepository is an autogenerated mock type for the JobRepository type
type JobRepository struct {
	mock.Mock
}

// Find provides a mock function with given fields: ctx, clauses
func (_m *JobRepository) Find(ctx context.Context, clauses []query.WhereClause) ([]models.Job, error) {
	ret := _m.Called(ctx, clauses)

	var r0 []models.Job
	if rf, ok := ret.Get(0).(func(context.Context, []query.WhereClause) []models.Job); ok {
		r0 = rf(ctx, clauses)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Job)
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

// Save provides a mock function with given fields: ctx, job
func (_m *JobRepository) Save(ctx context.Context, job *models.Job) (*models.Job, error) {
	ret := _m.Called(ctx, job)

	var r0 *models.Job
	if rf, ok := ret.Get(0).(func(context.Context, *models.Job) *models.Job); ok {
		r0 = rf(ctx, job)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Job)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Job) error); ok {
		r1 = rf(ctx, job)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateJobQuota provides a mock function with given fields: ctx, id, quota
func (_m *JobRepository) UpdateJobQuota(ctx context.Context, id uint, quota uint) (*models.Job, error) {
	ret := _m.Called(ctx, id, quota)

	var r0 *models.Job
	if rf, ok := ret.Get(0).(func(context.Context, uint, uint) *models.Job); ok {
		r0 = rf(ctx, id, quota)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Job)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint, uint) error); ok {
		r1 = rf(ctx, id, quota)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateJobStatus provides a mock function with given fields: ctx, id, status
func (_m *JobRepository) UpdateJobStatus(ctx context.Context, id uint, status string) (*models.Job, error) {
	ret := _m.Called(ctx, id, status)

	var r0 *models.Job
	if rf, ok := ret.Get(0).(func(context.Context, uint, string) *models.Job); ok {
		r0 = rf(ctx, id, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Job)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint, string) error); ok {
		r1 = rf(ctx, id, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewJobRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewJobRepository creates a new instance of JobRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewJobRepository(t mockConstructorTestingTNewJobRepository) *JobRepository {
	mock := &JobRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
