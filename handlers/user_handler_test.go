package handlers_test

import (
	"errors"
	"job-application/entity/models"
	"job-application/handlers"
	"job-application/middleware"
	"job-application/mocks"
	"net/http"
	"net/http/httptest"
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

var auths = []models.Auth{
	{
		ID:       1,
		Email:    "randysteven12@gmail.com",
		Password: "test_1234",
		User: models.User{
			ID:   1,
			Name: "Randy Steven",
		},
	},
	{
		ID:       2,
		Email:    "randysteven13@gmail.com",
		Password: "test_1234",
		User: models.User{
			ID:   2,
			Name: "Randy Steven",
		},
	},
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
		req, _ := http.NewRequest(http.MethodPost, "/v1/register", strings.NewReader(requestBody))
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		suite.userUseCase.On("RegisterUser",
			mock.Anything,
			mock.AnythingOfType("*models.User"),
			mock.AnythingOfType("*models.Auth"),
		).Return(&auths[0], nil)

		suite.userHandler.RegisterUser(ctx)
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusCreated, w.Code)
	})

	suite.Run("should return 500 failed to register user", func() {
		requestBody := `{
			"name": "Randy Steven",
			"email": "randysteven12@gmail.com",
			"password": "test_1234"
		}`

		suite.userUseCase.On("RegisterUser",
			mock.Anything,
			mock.AnythingOfType("*models.User"),
			mock.AnythingOfType("*models.Auth"),
		).Return(nil, errors.New("mock error"))

		req, _ := http.NewRequest(http.MethodPost, "/v1/register", strings.NewReader(requestBody))

		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		suite.router.POST("/v1/register", suite.userHandler.RegisterUser)
		suite.router.ServeHTTP(w, req)

		suite.T().Log(w.Body)

		suite.Equal(http.StatusInternalServerError, w.Code)
	})
}
