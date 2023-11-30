package handlers_test

import (
	"job-application/handlers"
	"job-application/middleware"
	"job-application/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type JobHandlerTestSuite struct {
	suite.Suite
	jobUsecase *mocks.JobUsecase
	jobHandler *handlers.JobHandler
	router     *gin.Engine
}

func (suite *JobHandlerTestSuite) SetupSubTest() {
	suite.jobUsecase = mocks.NewJobUsecase(suite.T())
	suite.jobHandler = handlers.NewJobHandler(suite.jobUsecase)
	suite.router = gin.Default()
	suite.router.Use(middleware.ErrorMiddleware())
}

func TestJobHandler(t *testing.T) {
	suite.Run(t, new(JobHandlerTestSuite))
}

func (suite *JobHandlerTestSuite) TestCreateJob() {

}
