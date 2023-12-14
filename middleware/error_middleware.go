package middleware

import (
	"job-application/apperror"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, ginErr := range c.Errors {
			apperror.ErrorChecker(c, ginErr.Err)
			return
		}
	}
}
