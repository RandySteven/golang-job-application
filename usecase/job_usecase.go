package usecase

import (
	"context"
	"job-application/entity/models"
	"job-application/interfaces"
	"job-application/query"
)

type jobUsecase struct {
	repo interfaces.JobRepository
}

// CreateJob implements interfaces.JobUsecase.
func (usecase *jobUsecase) CreateJob(ctx context.Context, job *models.Job) (*models.Job, error) {
	return usecase.repo.Save(ctx, job)
}

// FindAllJobs implements interfaces.JobUsecase.
func (usecase *jobUsecase) FindAllJobs(ctx context.Context, clauses []query.WhereClause) ([]models.Job, error) {
	return usecase.repo.Find(ctx, clauses)
}

func NewJobUsecase(repo interfaces.JobRepository) *jobUsecase {
	return &jobUsecase{repo: repo}
}

var _ interfaces.JobUsecase = &jobUsecase{}
