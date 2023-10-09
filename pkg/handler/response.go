package handler

import (
	"fr33d0mz/goMail/logger"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, status int, message string) {
	logger.Error.Println(message)
	c.AbortWithStatusJSON(status, errorResponse{
		Message: message,
	})
}
