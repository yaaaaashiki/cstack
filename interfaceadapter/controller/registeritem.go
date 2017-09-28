package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/yaaaaashiki/cstack/helper"
	"github.com/yaaaaashiki/cstack/usecase"
)

const ZeroValue = 0

type RegisterItemController struct {
	registerItemUseCase *usecase.RegisterItemUseCase
}

//input data by the front view
type InputItemField struct {
	Name        string `binding:"required" json:"name"`
	Price       int    `binding:"required" json:"price"`
	IconImage   string `binding:"required" json:"icon_image"`
	Description string `json:"description"`
}

type RegisterItemRequest struct {
	UserID      uint   `binding:"required" json:"user_id"`
	Name        string `binding:"required" json:"name"`
	Price       int    `binding:"required" json:"price" gorm:"column:price"`
	IconImage   string `binding:"required" json:"icon_image"`
	Description string `json:"description" gorm:"column:description"`
}

func NewRegisterItemController(registerItemUseCase *usecase.RegisterItemUseCase) *RegisterItemController {
	return &RegisterItemController{
		registerItemUseCase: registerItemUseCase,
	}
}

//Before running this function, should be judged new item or usual item.
//If new item input the data, run this function
func (s *RegisterItemController) Execute(c *gin.Context) {
	in := &InputItemField{}
	if err := c.MustBindWith(in, binding.JSON); err != nil {
		helper.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	_, userID, err := helper.GetSession(c)

	if err != nil {
		helper.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	if userID == ZeroValue {
		helper.ResponseErrorJSON(c, http.StatusBadRequest, errors.New("Cannot get session").Error())
		return
	}

	req := &usecase.RegisterItemRequest{
		UserID:      userID,
		Name:        in.Name,
		Price:       in.Price,
		IconImage:   in.IconImage,
		Description: in.Description,
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
