package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/yaaaaashiki/cstack/helper"
	"github.com/yaaaaashiki/cstack/usecase"
)

type RegisterController struct {
	registerUseCase *usecase.RegisterUseCase
}

//input data by the front view
type InputField struct {
	Biography string `binding:"required" json:"biography"`
	Sex       int    `binding:"required" json:"sex"`
}

type RegisterRequest struct {
	Name      string `binding:"required" json:"name"`
	Email     string `binding:"required" json:"email"`
	Salt      string `json:"salt" gorm:"column:salt"`
	Salted    string `json:"salted" gorm:"column:salted"`
	IconImage string `binding:"required" json:"icon_image"`
}

func NewRegisterController(registerUseCase *usecase.RegisterUseCase) *RegisterController {
	return &RegisterController{
		registerUseCase: registerUseCase,
	}
}

//Before running this function, should be judged new user or usual user.
//If new user input the data, run this function
func (s *RegisterController) Execute(c *gin.Context) {

	in := &InputField{}
	if err := c.MustBindWith(in, binding.JSON); err != nil {
		helper.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	//for debug
	helper.SaveSessionForDebug(c, uint(helper.CreateRand()), "helperSessionUser", "helperSessionUser.com", "")

	if err != nil {
		helper.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	req := &usecase.RegisterRequest{
		Name:      name,
		Email:     email,
		Salt:      salt,
		Salted:    salted,
		IconImage: userPicture,
	}

	res, err := s.registerUseCase.Execute(req)
	if err != nil {
		helper.ResponseErrorJSON(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user_id": res.User.ID,
	})
}
