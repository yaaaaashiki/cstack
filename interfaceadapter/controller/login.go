package controller

import (
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/yaaaaashiki/cstack/helper"
	"github.com/yaaaaashiki/cstack/usecase"
)

type LoginController struct {
	loginUseCase *usecase.LoginUseCase
}

type LoginRequest struct {
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

func NewLoginController(loginUseCase *usecase.LoginUseCase) *LoginController {
	return &LoginController{
		loginUseCase: loginUseCase,
	}
}

func (s *LoginController) Execute(c *gin.Context) {
	payload := new(LoginRequest)
	if err := c.ShouldBindWith(payload, binding.JSON); err != nil {
		log.Println(err)
		helper.ResponseError(
			c,
			http.StatusBadRequest,
			usecase.ErrorInvalidRequestFormat,
			" invalid request format")
		return
	}

	res, err := s.loginUseCase.Execute(&usecase.LoginRequest{Email: payload.Email, Password: payload.Password})
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if !res.Success {
		helper.ResponseError(c, http.StatusUnauthorized, res.Error.ErrorCode, res.Error.Message)
		return
	}

	if err := helper.SaveSession(c, payload.Email, res.UserID); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": res.UserID})
}
