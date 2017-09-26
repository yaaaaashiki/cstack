package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/yaaaaashiki/cstack/helper"
	"github.com/yaaaaashiki/cstack/usecase"
)

const Rand = 100

type RegisterController struct {
	registerUseCase *usecase.RegisterUseCase
}

//input data by the front view
type InputField struct {
	Name          string `binding:"required" json:"name"`
	Email         string `binding:"required" json:"email"`
	InputPassword string `binding:"required" json:"password"`
	IconImage     string `binding:"required" json:"icon_image"`
}

type RegisterRequest struct {
	Name      string `binding:"required" json:"name"`
	Email     string `binding:"required" json:"email"`
	Salt      string `binding:"required" json:"salt"`
	Salted    string `binding:"required" json:"salted"`
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

	salt := helper.Salt(Rand)
	salted := helper.Stretch(in.InputPassword, salt)

	req := &usecase.RegisterRequest{
		Name:      in.Name,
		Email:     in.Email,
		Salt:      salt,
		Salted:    salted,
		IconImage: in.IconImage,
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
