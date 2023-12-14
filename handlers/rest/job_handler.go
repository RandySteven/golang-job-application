package handlers_rest

import (
	"context"
	"fmt"
	"job-application/entity/models"
	"job-application/entity/payload"
	"job-application/enums"
	"job-application/interfaces"
	"job-application/query"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type JobHandler struct {
	usecase interfaces.JobUsecase
}

func NewJobHandler(usecase interfaces.JobUsecase) *JobHandler {
	return &JobHandler{usecase: usecase}
}

var _ interfaces.JobHandler = &JobHandler{}

func (handler *JobHandler) CreateJob(c *gin.Context) {
	var request payload.JobRequest
	requestId := uuid.NewString()
	ctx := context.WithValue(c.Request.Context(), "request_id", requestId)

	if err := c.ShouldBind(&request); err != nil {
		c.Error(err)
		return
	}

	job := &models.Job{
		Name:        request.Name,
		Quota:       request.Quota,
		ExpiryDate:  request.ExpiryDate,
		JobPosterID: request.JobPosterID,
		Status:      enums.Open,
	}

	job, err := handler.usecase.CreateJob(ctx, job)
	if err != nil {
		c.Error(err)
		return
	}
	resp := payload.Response{
		Message: "Success created user",
		Data:    job,
	}
	c.JSON(http.StatusCreated, resp)
}

// GetAllJobs implements interfaces.JobHandler.
func (handler *JobHandler) GetAllJobs(c *gin.Context) {
	var search payload.SearchJob
	c.ShouldBindQuery(&search)
	requestId := uuid.NewString()
	ctx := context.WithValue(c.Request.Context(), "request_id", requestId)

	val := reflect.ValueOf(&search).Elem()
	var whereClauses []query.WhereClause
	for i := 0; i < val.NumField(); i++ {
		whereClause := &query.WhereClause{
			Field:     val.Type().Field(i).Name,
			Value:     fmt.Sprintf("%v", val.Field(i).Interface()),
			Condition: "ilike",
		}
		whereClauses = append(whereClauses, *whereClause)
	}

	jobs, err := handler.usecase.FindAllJobs(ctx, nil)
	if err != nil {
		c.Error(err)
		return
	}
	resp := payload.Response{
		Message: "Get all jobs",
		Data:    jobs,
	}
	c.JSON(http.StatusOK, resp)
}
