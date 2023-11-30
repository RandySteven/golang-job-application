package handlers

import (
	"context"
	"job-application/entity/models"
	"job-application/entity/payload"
	"job-application/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ApplicationHandler struct {
	usecase interfaces.ApplicationUsecase
}

// GetApplicationById implements interfaces.ApplicationHandler.
func (handler *ApplicationHandler) GetApplicationById(c *gin.Context) {
	id := c.Param("id")
	requestId := uuid.NewString()
	ctx := context.WithValue(c.Request.Context(), "request_id", requestId)

	intId, err := strconv.Atoi(id)
	if err != nil {
		c.Error(err)
		return
	}

	application, err := handler.usecase.GetApplicationById(ctx, uint(intId))
	if err != nil {
		c.Error(err)
		return
	}
	resp := payload.Response{
		Message: "Get applicaiton",
		Data:    application,
	}
	c.JSON(http.StatusOK, resp)
}

// ApplyApplication implements interfaces.ApplicationHandler.
func (handler *ApplicationHandler) ApplyApplication(c *gin.Context) {
	requestId := uuid.NewString()
	ctx := context.WithValue(c.Request.Context(), "request_id", requestId)
	var request *payload.ApplicationRequest

	if err := c.ShouldBind(&request); err != nil {
		c.Error(err)
		return
	}

	application := &models.Application{
		UserID: request.UserID,
		JobID:  request.JobID,
	}

	application, err := handler.usecase.ApplyApplication(ctx, application)
	if err != nil {
		c.Error(err)
		return
	}

	resp := payload.Response{
		Message: "Application success",
		Data:    application,
	}

	c.JSON(http.StatusOK, resp)
}

var _ interfaces.ApplicationHandler = &ApplicationHandler{}

func NewApplicationHandler(usecase interfaces.ApplicationUsecase) *ApplicationHandler {
	return &ApplicationHandler{usecase: usecase}
}