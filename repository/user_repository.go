package repository

import (
	"context"
	"job-application/entity/models"
	"job-application/entity/payload"
	"job-application/interfaces"
	"job-application/query"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// RegistertUser implements interfaces.UserRepository.
func (repo *UserRepository) RegistertUser(ctx context.Context, request *payload.UserRequest) (*models.User, error) {
	repo.db.Exec("SELECT pg_sleep(5)")
	user := &models.User{
		Name: request.Name,
	}

	auth := &models.Auth{
		Email:    request.Email,
		Password: request.Password,
	}

	err := repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		err := tx.Table("users").Create(&user).Error
		if err != nil {
			return err
		}

		auth.UserID = user.ID

		err = tx.Table("auths").Create(&auth).Error
		if err != nil {
			return err
		}

		return nil
	})
	return user, err
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

var _ interfaces.UserRepository = &UserRepository{}

// Find implements interfaces.UserRepository.
func (repo *UserRepository) Find(ctx context.Context, clauses []query.WhereClause) ([]models.User, error) {
	var users []models.User
	err := repo.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Save implements interfaces.UserRepository.
func (repo *UserRepository) Save(ctx context.Context, user *models.User) (*models.User, error) {
	err := repo.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUserById implements interfaces.UserRepository.
func (*UserRepository) UpdateUserById(ctx context.Context, id int, user *models.User) (*models.User, error) {
	return nil, nil
}
