package repository

import (
	"context"
	"job-application/entity/models"
	"job-application/interfaces"
	"job-application/query"

	"gorm.io/gorm"
)

type jobRepository struct {
	db *gorm.DB
}

// Find implements interfaces.jobRepository.
func (repo *jobRepository) Find(ctx context.Context, clauses []query.WhereClause) ([]models.Job, error) {
	var jobs []models.Job
	err := repo.db.Find(&jobs).Error
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

// Save implements interfaces.jobRepository.
func (repo *jobRepository) Save(ctx context.Context, job *models.Job) (*models.Job, error) {
	err := repo.db.Create(&job).Error
	if err != nil {
		return nil, err
	}
	return job, nil
}

func (repo *jobRepository) FindOpenJob(ctx context.Context, clauses []query.WhereClause) ([]models.Job, error) {
	var jobs []models.Job
	err := repo.db.Where("status = open").Find(&jobs).Error
	if err != nil {
		return nil, err
	}
	return jobs, nil

}

// UpdateJobQuota implements interfaces.jobRepository.
func (*jobRepository) UpdateJobQuota(ctx context.Context, id uint, quota uint) (*models.Job, error) {
	panic("unimplemented")
}

// UpdateJobStatus implements interfaces.jobRepository.
func (*jobRepository) UpdateJobStatus(ctx context.Context, id uint, status string) (*models.Job, error) {
	panic("unimplemented")
}

func NewJobRepository(db *gorm.DB) *jobRepository {
	return &jobRepository{db}
}

var _ interfaces.JobRepository = &jobRepository{}
