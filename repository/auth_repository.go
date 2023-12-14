package repository

import (
	"context"
	"job-application/entity/models"
	"job-application/interfaces"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

// GetUserAuth implements interfaces.AuthRepository.
func (repo *AuthRepository) GetUserAuth(ctx context.Context, userId uint) (*models.Auth, error) {
	var auth *models.Auth
	err := repo.db.WithContext(ctx).Model(&models.Auth{}).
		Preload("User").
		Where("user_id = ?", userId).
		Find(&auth).Error
	if err != nil {
		return nil, err
	}
	return auth, nil
}

// GetByEmailAndPassword implements interfaces.AuthRepository.
func (repo *AuthRepository) GetByEmail(ctx context.Context, email string) (*models.Auth, error) {
	var auth *models.Auth
	err := repo.db.Model(&models.Auth{}).Where("email = ? ", email).Find(&auth).Error
	if err != nil {
		return nil, err
	}
	return auth, nil
}

// Save implements interfaces.AuthRepository.
func (repo *AuthRepository) Save(ctx context.Context, auth *models.Auth) (*models.Auth, error) {
	err := repo.db.Create(&auth).Error
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db}
}

var _ interfaces.AuthRepository = &AuthRepository{}
