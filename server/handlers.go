package server

import (
	"job-application/configs"
	rest "job-application/handlers/rest"
	"job-application/interfaces"
	"job-application/usecase"
)

type (
	Handlers struct {
		JobHandler         interfaces.JobHandler
		UserHandler        interfaces.UserHandler
		ApplicationHandler interfaces.ApplicationHandler
	}
)

func NewHandlers(repo configs.Repository) (*Handlers, error) {
	usecase := usecase.NewUsecase(repo)

	return &Handlers{
		JobHandler:         rest.NewJobHandler(usecase.JobUsecase),
		UserHandler:        rest.NewUserHandler(usecase.UserUsecase),
		ApplicationHandler: rest.NewApplicationHandler(usecase.ApplicationUsecase),
	}, nil
}
