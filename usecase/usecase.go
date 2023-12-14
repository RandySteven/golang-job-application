package usecase

import (
	"job-application/configs"
	"job-application/interfaces"
)

type Usecase struct {
	interfaces.JobUsecase
	interfaces.UserUsecase
	interfaces.ApplicationUsecase
}

func NewUsecase(repo configs.Repository) *Usecase {
	return &Usecase{
		JobUsecase:         NewJobUsecase(repo.JobRepository),
		UserUsecase:        NewUserUsecase(repo.UserRepository, repo.AuthRepository),
		ApplicationUsecase: NewApplicationUsecase(repo.ApplicationRepository),
	}
}
