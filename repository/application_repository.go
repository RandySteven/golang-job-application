package repository

import (
	"context"
	"job-application/apperror"
	"job-application/entity/models"
	"job-application/interfaces"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type applicationRepository struct {
	db *gorm.DB
}

// GetApplicationByUserIdAndJobId implements interfaces.ApplicationRepository.
func (repo *applicationRepository) GetApplicationByUserIdAndJobId(ctx context.Context, userId uint, jobId uint) (*models.Application, error) {
	var application *models.Application
	err := repo.db.WithContext(ctx).Model(&models.Application{}).
		Where("user_id = ?", userId).
		Where("job_id = ?", jobId).
		Find(&application).Error
	if err != nil {
		return nil, err
	}
	return application, nil
}

// GetApplicationById implements interfaces.ApplicationRepository.
func (repo *applicationRepository) GetApplicationById(ctx context.Context, id uint) (*models.Application, error) {
	var application *models.Application

	err := repo.db.WithContext(ctx).Model(&models.Application{}).
		Preload("User").
		Preload("Job").
		Where("id = ?", id).
		Find(&application).Error
	if err != nil {
		return nil, err
	}
	return application, nil
}

// Save implements interfaces.ApplicationRepository.
func (repo *applicationRepository) Save(ctx context.Context, application *models.Application) (*models.Application, error) {
	err := repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var job *models.Job
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Model(&models.Job{}).
			Where("id = ?", application.JobID).
			Scan(&job).Error
		if err != nil || job == nil {
			return apperror.NewErrJobIdNotFound()
		}

		if job.Quota == 0 {
			return apperror.NewErrJobQuotaZero()
		}

		err = tx.Table("jobs").
			Where("id = ?", job.ID).
			Update("quota", gorm.Expr("quota - ? ", 1)).
			Error
		if err != nil {
			return err
		}

		err = tx.Create(&application).Error
		if err != nil {
			return err
		}

		return nil
	})
	return application, err
}

func NewApplicationRepository(db *gorm.DB) *applicationRepository {
	return &applicationRepository{db: db}
}

var _ interfaces.ApplicationRepository = &applicationRepository{}