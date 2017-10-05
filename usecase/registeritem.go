package usecase

import (
	"errors"

	"github.com/yaaaaashiki/cstack/domain/model"
	"github.com/yaaaaashiki/cstack/domain/repository"
)

type RegisterItemUseCase struct {
	itemRepository *repository.ItemRepository
}

type RegisterItemRequest struct {
	UserID              uint   `json:"user_id" gorm:"column:user_id"`
	Name                string `json:"name" gorm:"column:name"`
	Price               int    `json:"price" gorm:"column:price"`
	CurrentPaymentPrice int    `json:"current_payment_price" gorm:"column:current_payment_price"`
	IconImage           string `json:"icon_image" gorm:"column:icon_image"`
	Description         string `json:"description" gorm:"column:description"`
}

type RegisterItemResponse struct {
	Item *model.Item
}

func NewRegisterItemUseCase(itemRepository *repository.ItemRepository) *RegisterItemUseCase {
	return &RegisterItemUseCase{
		itemRepository: itemRepository,
	}
}

//Before running this function, should be judged new item or usual item.
//If new item input the data, run this function
func (f *RegisterItemUseCase) Execute(req *RegisterItemRequest) (*RegisterItemResponse, error) { //TODO test

	exist, err := f.itemRepository.IsExistItem(req.UserID, req.Name)

	if err != nil {
		return nil, err
	}
	if exist {
		return nil, errors.New("Duplicate user cannnot create new user")
	}

	item, err := f.itemRepository.RegisterItem(req.UserID, req.Name, req.Price, req.IconImage, req.Description)
	if err != nil {
		return nil, err
	}

	return &RegisterItemResponse{Item: item}, nil
}
