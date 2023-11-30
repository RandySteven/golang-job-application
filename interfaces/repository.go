package interfaces

import (
	"context"
	"job-application/entity/models"
	"job-application/entity/payload"
	"job-application/query"
)

type (
	UserRepository interface {
		Save(ctx context.Context, user *models.User) (*models.User, error)
		Find(ctx context.Context, clauses []query.WhereClause) ([]models.User, error)
		UpdateUserById(ctx context.Context, id int, user *models.User) (*models.User, error)
		RegistertUser(ctx context.Context, request *payload.UserRequest) (*models.User, error)
	}

	AuthRepository interface {
		Save(ctx context.Context, auth *models.Auth) (*models.Auth, error)
		GetByEmail(ctx context.Context, email string) (*models.Auth, error)
		GetUserAuth(ctx context.Context, userId uint) (*models.Auth, error)
	}

	JobRepository interface {
		Save(ctx context.Context, job *models.Job) (*models.Job, error)
		Find(ctx context.Context, clauses []query.WhereClause) ([]models.Job, error)
		UpdateJobQuota(ctx context.Context, id uint, quota uint) (*models.Job, error)
		UpdateJobStatus(ctx context.Context, id uint, status string) (*models.Job, error)
	}

	ApplicationRepository interface {
		Save(ctx context.Context, application *models.Application) (*models.Application, error)
		GetApplicationById(ctx context.Context, id uint) (*models.Application, error)
		GetApplicationByUserIdAndJobId(ctx context.Context, userId uint, jobId uint) (*models.Application, error)
		Find(ctx context.Context, clauses []query.WhereClause) ([]models.Application, error)
	}
)
