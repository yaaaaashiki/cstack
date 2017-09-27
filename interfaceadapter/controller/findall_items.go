package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yaaaaashiki/cstack/helper"
	"github.com/yaaaaashiki/cstack/usecase"
)

type FindAllItemsController struct {
	findAllItemsUseCase *usecase.FindAllItemsUseCase
}

func NewFindAllItemsController(findAllItemsUseCase *usecase.FindAllItemsUseCase) *FindAllItemsController {
	return &FindAllItemsController{
		findAllItemsUseCase: findAllItemsUseCase,
	}
}

func (s *FindAllItemsController) Execute(c *gin.Context) {
	userID := c.Param("userId")

	rawRes, err := s.findAllItemsUseCase.Execute()
	if err != nil {
		helper.ResponseErrorJSON(c, http.StatusBadRequest, err.Error())
		return
	}

	//res := presenter.FindAllItems(rawRes)

	c.JSON(http.StatusOK, &rawRes)
}
