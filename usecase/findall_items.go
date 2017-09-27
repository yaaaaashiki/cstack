package usecase

import (
	"github.com/yaaaaashiki/cstack/domain/repository"
	"strconv"
)

type FindAllItemsUseCase struct {
	itemRepository *repository.ItemRepository
}

type FindAllItemsResponse struct {
	Items []*itemDTO
}

type itemDTO struct {
	ID                  uint   `json:"id" gorm:"column:id"`
	UserID              uint   `json:"user_id" gorm:"column:user_id"`
	Name                string `json:"name" gorm:"column:name"`
	Price               string `json:"price" gorm:"column:price"`
	CurrentPaymentPrice string `json:"current_payment_price" gorm:"column:current_payment_price"`
	IconImage           string `json:"icon_image" gorm:"column:icon_image"`
	Description         string `json:"description" gorm:"column:description"`
}

func NewFindAllItemsUseCase(itemRepository *repository.ItemRepository) *FindAllItemsUseCase {
	return &FindAllItemsUseCase{
		itemRepository: itemRepository,
	}
}

//Before running this function, should be full with items table column
func (f *FindAllItemsUseCase) Execute(userID string) (*FindAllItemsResponse, error) {
	res, err := f.itemRepository.FindAllByUserIDOrNil(userID)
	if err != nil {
		return nil, err
	}

	c := &FindAllItemsResponse{
		Items: make([]*itemDTO, 0),
	}

	for _, v := range res {
		itemDTO := &itemDTO{}
		itemDTO.ID = uint(v.ID)
		itemDTO.UserID = uint(v.UserID)
		itemDTO.Name = v.Name
		itemDTO.Price = strconv.Itoa(v.Price)
		itemDTO.CurrentPaymentPrice = strconv.Itoa(v.CurrentPaymentPrice)
		itemDTO.IconImage = v.IconImage
		itemDTO.Description = v.Description
		c.Items = append(c.Items, itemDTO)
	}

	return c, nil
}
