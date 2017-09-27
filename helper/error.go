package helper

import "github.com/gin-gonic/gin"

type Error struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}

func NewError(errorCode int, message string) *Error {
	return &Error{ErrorCode: errorCode, Message: message}
}

func ResponseErrorJSON(c *gin.Context, httpStatusCode int, message string) {
	ResponseErrorRaw(c, httpStatusCode, &Error{Message: message})
}

func ResponseErrorRaw(c *gin.Context, httpStatusCode int, error *Error) {
	c.JSON(httpStatusCode, gin.H{
		"errors": []*Error{error},
	})
}
