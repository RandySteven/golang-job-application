package server

import "github.com/gin-gonic/gin"

func (h *Handlers) InitRouter(r *gin.RouterGroup) {

	jobRouter := r.Group("/jobs")
	jobRouter.GET("", h.JobHandler.GetAllJobs)
	jobRouter.POST("", h.JobHandler.CreateJob)

	applicationRouter := r.Group("/applications")
	applicationRouter.POST("", h.ApplicationHandler.ApplyApplication)
	applicationRouter.GET("/:id", h.ApplicationHandler.GetApplicationById)
}
