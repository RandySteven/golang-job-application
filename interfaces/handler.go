package interfaces

import "github.com/gin-gonic/gin"

type (
	JobHandler interface {
		CreateJob(c *gin.Context)
		GetAllJobs(c *gin.Context)
	}

	UserHandler interface {
		RegisterUser(c *gin.Context)
		LoginUser(c *gin.Context)
		Logout(c *gin.Context)
	}

	ApplicationHandler interface {
		ApplyApplication(c *gin.Context)
		GetApplicationById(c *gin.Context)
		GetAllApplications(c *gin.Context)
	}
)
