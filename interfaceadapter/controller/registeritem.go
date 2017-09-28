package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/yaaaaashiki/cstack/helper"
	"github.com/yaaaaashiki/cstack/usecase"
)

type RegisterItemController struct {
	registerItemUseCase *usecase.RegisterItemUseCase
}

//input data by the front view
type InputItemField struct {
	Name          string `binding:"required" json:"name"`
	Email         string `binding:"required" json:"email"`
	InputPassword string `binding:"required" json:"password"`
	IconImage     string `binding:"required" json:"icon_image"`
}

type RegisterItemRequest struct {
	Name      string `binding:"required" json:"name"`
	Email     string `binding:"required" json:"email"`
	Salt      string `binding:"required" json:"salt"`
	Salted    string `binding:"required" json:"salted"`
	IconImage string `binding:"required" json:"icon_image"`
}

func NewRegisterItemController(registerItemUseCase *usecase.RegisterItemUseCase) *RegisterItemController {
	return &RegisterItemController{
		registerItemUseCase: registerItemUseCase,
	}
}

//Before running this function, should be judged new item or usual item.
//If new item input the data, run this function
func (s *RegisterItemController) Execute(c *gin.Context) {

	in := &InputField{}
	if err := c.MustBindWith(in, binding.JSON); err != nil {
		helper.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	salt := helper.Salt(Rand)
	salted := helper.Stretch(in.InputPassword, salt)

	req := &usecase.RegisterItemRequest{
		Name:      in.Name,
		Email:     in.Email,
		Salt:      salt,
		Salted:    salted,
		IconImage: in.IconImage,
	}

	res, err := s.registerItemUseCase.Execute(req)
	if err != nil {
		helper.ResponseErrorJSON(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"item_id": res.Item.ID,
	})
}
