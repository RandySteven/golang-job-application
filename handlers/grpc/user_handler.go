package handler_grpc

import (
	"context"
	"job-application/entity/models"
	"job-application/interfaces"
	pb "job-application/proto"
	"time"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	usecase interfaces.UserUsecase
}

func (h *UserHandler) RegisterUser(ctx context.Context, req *pb.UserRequest) (*pb.UserProfile, error) {
	date, err := time.Parse("2006-01-02", req.Birthdate)
	if err != nil {
		return nil, err
	}
	user := &models.User{
		Name:        req.Name,
		DateOfBirth: date,
	}

	auth := &models.Auth{
		Email:    req.Email,
		Password: req.Password,
	}

	auth, err = h.usecase.RegisterUser(ctx, user, auth)
	if err != nil {
		return nil, err
	}

	return &pb.UserProfile{
		Name:     user.Name,
		Email:    auth.Email,
		Birthday: user.DateOfBirth.GoString(),
	}, nil
}

func (h *UserHandler) LoginUser(ctx context.Context, req *pb.UserLoginRequest) (*pb.LoginResponse, error) {
	token, err := h.usecase.LoginUser(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{Token: token}, nil
}

func NewUserHandler(usecase interfaces.UserUsecase) *UserHandler {
	return &UserHandler{
		usecase: usecase,
	}
}
