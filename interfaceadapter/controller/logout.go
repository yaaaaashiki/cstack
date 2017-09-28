package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yaaaaashiki/cstack/helper"
)

type LogoutController struct{}

func NewLogoutController() *LogoutController {
	return &LogoutController{}
}

func (s *LogoutController) Execute(c *gin.Context) {
	helper.ClearSession(c)
	c.Status(http.StatusNoContent)
}
