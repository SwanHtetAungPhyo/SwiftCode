package model

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Error      string `json:"error"`
}

func NewErrorMessage(c *gin.Context, status int, err, msg string) {
	c.JSON(status, ErrorResponse{
		StatusCode: status,
		Message:    msg,
		Error:      err,
	})
}
