package helper

import "github.com/gin-gonic/gin"

type Error struct {
	Message string `json:"message"`
}

func ResponseErrorJSON(c *gin.Context, httpStatusCode int, message string) {
	ResponseErrorRaw(c, httpStatusCode, &Error{Message: message})
}

func ResponseErrorRaw(c *gin.Context, httpStatusCode int, error *Error) {
	c.JSON(httpStatusCode, gin.H{
		"errors": []*Error{error},
	})
}
