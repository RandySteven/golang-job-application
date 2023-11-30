package usecase

import (
	"context"
	"job-application/entity/models"
	"job-application/interfaces"
)

type userUsecase struct {
	userRepo interfaces.UserRepository
	authRepo interfaces.AuthRepository
}

// LoginUser implements interfaces.UserUsecase.
func (usecase *userUsecase) LoginUser(ctx context.Context, email string, password string) (string, error) {
	panic("unimplemented")
}

// RegisterUser implements interfaces.UserUsecase.
func (usecase *userUsecase) RegisterUser(ctx context.Context, user *models.User, auth *models.Auth) (*models.Auth, error) {
	user, err := usecase.userRepo.Save(ctx, user)
	if err != nil {
		return nil, err
	}
	auth.UserID = user.ID
	_, err = usecase.authRepo.Save(ctx, auth)
	if err != nil {
		return nil, err
	}
	return usecase.authRepo.GetUserAuth(ctx, auth.UserID)
}

func NewUserUsecase(userRepo interfaces.UserRepository,
	authRepo interfaces.AuthRepository) *userUsecase {
	return &userUsecase{
		userRepo: userRepo,
		authRepo: authRepo,
	}
}

var _ interfaces.UserUsecase = &userUsecase{}
