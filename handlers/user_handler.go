package handlers

import (
	"context"
	"fmt"
	"job-application/apperror"
	"job-application/entity/models"
	"job-application/entity/payload"
	"job-application/interfaces"
	"job-application/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UserHandler struct {
	usecase interfaces.UserUsecase
}

func (handler *UserHandler) Logout(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})
	resp := payload.Response{
		Message: "Success to logout",
	}
	c.JSON(http.StatusOK, resp)
}

// LoginUser implements interfaces.UserHandler.
func (handler *UserHandler) LoginUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	requestId := uuid.NewString()
	ctx := context.WithValue(c.Request.Context(), "request_id", requestId)
	var request payload.UserLoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	pass, err := utils.HashPassword(request.Password)
	request.Password = pass
	if err != nil {
		c.Error(apperror.NewErrPasswordTooLong())
		return
	}

	token, err := handler.usecase.LoginUser(ctx, request.Email, request.Password)
	if err != nil {
		c.Error(err)
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	resp := payload.Response{
		Message: "Success login user",
		Data:    token,
	}
	c.JSON(http.StatusOK, resp)
}

var _ interfaces.UserHandler = &UserHandler{}

func NewUserHandler(usecase interfaces.UserUsecase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

func (handler *UserHandler) RegisterUser(c *gin.Context) {
	requestId := uuid.NewString()
	ctx := context.WithValue(c.Request.Context(), "request_id", requestId)
	var request *payload.UserRequest

	if err := c.ShouldBind(&request); err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			errMsg := fmt.Sprintf("%s field is %s", fieldErr.Field(), fieldErr.ActualTag())
			c.Error(fmt.Errorf("%s", errMsg))
			return
		}
	}

	user := &models.User{
		Name: request.Name,
	}

	pass, err := utils.HashPassword(request.Password)
	if err != nil {
		c.Error(err)
		return
	}
	auth := &models.Auth{
		Email:    request.Email,
		Password: pass,
	}

	auth, err = handler.usecase.RegisterUser(ctx, user, auth)
	if err != nil {
		c.Error(err)
		return
	}

	userProfile := payload.UserProfile{
		Name:     auth.User.Name,
		Email:    auth.Email,
		Birthday: auth.User.DateOfBirth,
	}

	resp := payload.Response{
		Message: "Register user success",
		Data:    userProfile,
	}
	c.JSON(http.StatusOK, resp)
}
