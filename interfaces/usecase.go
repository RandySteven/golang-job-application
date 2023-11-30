package interfaces

import (
	"context"
	"job-application/entity/models"
	"job-application/query"
)

type (
	UserUsecase interface {
		RegisterUser(ctx context.Context, user *models.User, auth *models.Auth) (*models.Auth, error)
		LoginUser(ctx context.Context, email, password string) (string, error)
	}

	JobUsecase interface {
		CreateJob(ctx context.Context, job *models.Job) (*models.Job, error)
		FindAllJobs(ctx context.Context, clauses []query.WhereClause) ([]models.Job, error)
	}

	ApplicationUsecase interface {
		ApplyApplication(ctx context.Context, application *models.Application) (*models.Application, error)
		GetApplicationById(ctx context.Context, id uint) (*models.Application, error)
		GetAllApplications(ctx context.Context, clauses []query.WhereClause) ([]models.Application, error)
	}
)
