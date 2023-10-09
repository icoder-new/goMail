package handler

import (
	"fr33d0mz/goMail/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Subject   string `json:"subject" binding:"required"`
	Message   string `json:"message" binding:"required"`
}

func (h *Handler) SendMessage(c *gin.Context) {
	var req UserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	logger.Debug.Println(h.service.IsUserExistByEmail(req.Email))
	if !h.service.IsUserExistByEmail(req.Email) {
		logger.Debug.Println("User does not exist, process of registering is started")
		_, err := h.service.CreateUser(req.Firstname, req.Lastname, req.Email)
		if err != nil {
			logger.Debug.Println("we got error")
			logger.Error.Println(err)
			newErrorResponse(c, http.StatusInternalServerError, "internal server error")
			return
		}
	}

	err := h.service.SendMessage(req.Email, req.Subject, req.Message)
	if err != nil {
		logger.Debug.Println("Message was not sent")
		logger.Error.Println(err.Error())
		newErrorResponse(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(200, gin.H{
		"to":      req.Email,
		"subject": req.Subject,
		"message": "sent",
	})

}
