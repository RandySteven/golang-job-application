package usecase

import (
	"context"
	"job-application/apperror"
	"job-application/entity/models"
	"job-application/enums"
	"job-application/interfaces"
	"job-application/query"
)

type applicationUsecase struct {
	repo interfaces.ApplicationRepository
}

// GetAllApplications implements interfaces.ApplicationUsecase.
func (usecase *applicationUsecase) GetAllApplications(ctx context.Context, clauses []query.WhereClause) ([]models.Application, error) {
	return usecase.repo.Find(ctx, clauses)
}

// GetApplicationById implements interfaces.ApplicationUsecase.
func (usecase *applicationUsecase) GetApplicationById(ctx context.Context, id uint) (*models.Application, error) {
	return usecase.repo.GetApplicationById(ctx, id)
}

// ApplyApplication implements interfaces.ApplicationUsecase.
func (usecase *applicationUsecase) ApplyApplication(ctx context.Context, application *models.Application) (*models.Application, error) {
	applicationExists, _ := usecase.repo.GetApplicationByUserIdAndJobId(
		ctx,
		application.UserID,
		application.JobID,
	)
	if applicationExists != nil {
		return nil, apperror.NewErrorWrapper(enums.Conflict, "you can't apply this job")
	}

	application, err := usecase.repo.Save(ctx, application)
	if err != nil {
		return nil, err
	}

	applicationResponse, err := usecase.repo.GetApplicationById(ctx, application.ID)
	if err != nil {
		return nil, err
	}
	return applicationResponse, nil
}

func NewApplicationUsecase(repo interfaces.ApplicationRepository) *applicationUsecase {
	return &applicationUsecase{repo: repo}
}

var _ interfaces.ApplicationUsecase = &applicationUsecase{}
