package usecase

import (
	"context"
	"job-application/configs"
	"job-application/entity/models"
	"job-application/interfaces"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type userUsecase struct {
	userRepo interfaces.UserRepository
	authRepo interfaces.AuthRepository
}

// LoginUser implements interfaces.UserUsecase.
func (usecase *userUsecase) LoginUser(ctx context.Context, email string, password string) (string, error) {
	user, err := usecase.authRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	if err != nil {
		return "", err
	}
	// log.Println(user.Password)
	// err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	// if err != nil {
	// 	return "", err
	// }
	expTime := time.Now().Add(time.Minute * 15)
	claims := &configs.JWTClaim{
		ID:    user.ID,
		Name:  user.User.Name,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "issuer",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return tokenAlgo.SignedString(configs.JWT_KEY)
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
