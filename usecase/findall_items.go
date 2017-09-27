package usecase

import (
	"github.com/yaaaaashiki/cstack/domain/repository"
)

type FindAllCompaniesUseCase struct {
	itemRepository *repository.ItemRepository
}

type FindAllCompaniesResponse struct {
	Companies []*itemDTO
}

type itemDTO struct {
	ID         uint   `json:"id" gorm:"column:id"`
	FacebookID uint   `json:"facebook_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	URL        string `json:"url" gorm:"column:url"`
	Biography  string `json:"biography"`
	Zipcode    string `json:"zipcode" gorm:"column:zipcode"`
	Address1   string `json:"address1" gorm:"address1"`
	Address2   string `json:"address2" gorm:"address2"`
	Tel1       string `json:"tel1" gorm:"tel1"`
	Tel2       string `json:"tel2" gorm:"tel2"`
	Tel3       string `json:"tel3" gorm:"tel3"`
	IconImage  string `json:"icon_image"`
}

func NewFindAllCompanaisUseCase(itemRepository *repository.ItemRepository) *FindAllCompaniesUseCase {
	return &FindAllCompaniesUseCase{
		itemRepository: itemRepository,
	}
}

//Before running this function, should be full with companies table column
func (f *FindAllCompaniesUseCase) Execute() (*FindAllCompaniesResponse, error) {
	res, err := f.itemRepository.FindAll()
	if err != nil {
		return nil, err
	}

	c := &FindAllCompaniesResponse{
		Companies: make([]*itemDTO, 0),
	}

	for _, v := range res {
		itemDTO := &itemDTO{}
		itemDTO.ID = uint(v.ID)
		itemDTO.FacebookID = v.FacebookID
		itemDTO.Name = v.Name
		itemDTO.Email = v.Email
		itemDTO.URL = v.URL
		itemDTO.Biography = v.Biography
		itemDTO.Zipcode = v.Zipcode
		itemDTO.Address1 = v.Address1
		itemDTO.Address2 = v.Address2
		itemDTO.Tel1 = v.Tel1
		itemDTO.Tel2 = v.Tel2
		itemDTO.Tel3 = v.Tel3
		itemDTO.IconImage = v.IconImage
		c.Companies = append(c.Companies, itemDTO)
	}

	return c, nil
}
