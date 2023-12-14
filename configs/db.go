package configs

import (
	"fmt"
	"job-application/entity/models"
	"job-application/interfaces"
	"job-application/repository"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repository struct {
	JobRepository         interfaces.JobRepository
	AuthRepository        interfaces.AuthRepository
	UserRepository        interfaces.UserRepository
	ApplicationRepository interfaces.ApplicationRepository
	db                    *gorm.DB
}

func NewRepository(config *models.Config) (*Repository, error) {
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DbHost,
		config.DbPort,
		config.DbUser,
		config.DbPass,
		config.DbName,
	)
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	return &Repository{
		JobRepository:         repository.NewJobRepository(db),
		UserRepository:        repository.NewUserRepository(db),
		AuthRepository:        repository.NewAuthRepository(db),
		ApplicationRepository: repository.NewApplicationRepository(db),
		db:                    db,
	}, nil
}

func (r *Repository) Automigrate() error {
	r.db.Migrator().DropTable(&models.User{})
	r.db.Migrator().DropTable(&models.Auth{})
	r.db.Migrator().DropTable(&models.Job{})
	r.db.Migrator().DropTable(&models.UserJob{})
	r.db.Migrator().DropTable(&models.Application{})

	return r.db.AutoMigrate(
		&models.User{},
		&models.Auth{},
		&models.Job{},
		&models.UserJob{},
		&models.Application{},
	)
}

func (r *Repository) GetConnection() *gorm.DB {
	return r.db
}
