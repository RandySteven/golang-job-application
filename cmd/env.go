package cmd

import (
	"job-application/configs"
	"job-application/entity/models"
	"job-application/server"
	"log"
	"os"
)

func InitConfig() *models.Config {
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	return models.NewConfig(dbHost, dbPort, dbUser, dbPass, dbName)
}

func AppPort() string {
	return os.Getenv("APP_PORT")
}

func InitRepository() (*configs.Repository, error) {
	config := InitConfig()

	return configs.NewRepository(config)
}

func InitHandlers() *server.Handlers {

	repository, err := InitRepository()

	if err != nil {
		log.Println(err)
		return nil
	}

	// err = repository.Automigrate()
	if err != nil {
		return nil
	}

	handlers, err := server.NewHandlers(*repository)
	if err != nil {
		return nil
	}

	return handlers
}
