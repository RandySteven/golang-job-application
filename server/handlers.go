package server

import (
	"job-application/configs"
	"job-application/handlers"
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
		JobHandler:         handlers.NewJobHandler(usecase.JobUsecase),
		UserHandler:        handlers.NewUserHandler(usecase.UserUsecase),
		ApplicationHandler: handlers.NewApplicationHandler(usecase.ApplicationUsecase),
	}, nil
}
