package payload

import (
	"job-application/entity/models"
	"time"
)

type UserRequest struct {
	Name     string    `json:"name" binding:"required,min=3,max=32"`
	Birthday time.Time `json:"birthday"`
	Email    string    `json:"email" binding:"required,email"`
	Password string    `json:"password"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserProfile struct {
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Birthday time.Time `json:"birthday"`
}

type UserDetail struct {
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	CurrentJob models.Job `json:"current_job"`
}
