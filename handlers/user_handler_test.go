package handlers_test

import (
	"job-application/handlers"
	"job-application/middleware"
	"job-application/mocks"
	"net/http"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserHandlerTestSuite struct {
	suite.Suite
	userUseCase *mocks.UserUsecase
	userHandler *handlers.UserHandler
	router      *gin.Engine
}

func (suite *UserHandlerTestSuite) SetupSubTest() {
	suite.userUseCase = mocks.NewUserUsecase(suite.T())
	suite.userHandler = handlers.NewUserHandler(suite.userUseCase)
	suite.router = gin.Default()
	suite.router.Use(middleware.ErrorMiddleware())
}

func TestUserHandler(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}

func (suite *UserHandlerTestSuite) TestRegisterUser() {
	suite.Run("should return 201 success to register user", func() {
		requestBody := `{
			"name": "Randy Steven",
			"email": "randysteven12@gmail.com",
			"password": "test_1234"
		}`

		suite.userUseCase.On("RegisterUser",
			mock.Anything,
			mock.AnythingOfType("*models.User"),
			mock.AnythingOfType("*models.Auth"),
		).Return(mock.AnythingOfType("*models.Auth"), nil)

		req, _ := http.NewRequest(http.MethodPost, "/v1/register", strings.NewReader(requestBody))

		req.Header.Set("Content-Type", "application/json")
	})
}
