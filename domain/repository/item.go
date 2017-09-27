package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/yaaaaashiki/cstack/domain/model"
)

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{
		db: db,
	}
}

func (f *ItemRepository) FindAll() ([]model.Item, error) {
	items := []model.Item{}
	if err := f.db.Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (f *ItemRepository) FindByID(id int) (*model.Item, error) {
	items := model.Item{}
	if err := f.db.Find(&items, "id=?", id).Error; err != nil {
		return nil, err
	}
	return &items, nil
}

func (f *ItemRepository) FindByIDOrNil(id uint) (*model.Item, error) {
	item := model.Item{}
	res := f.db.Find(&item, "id=?", id)
	if res.RecordNotFound() {
		return nil, nil
	} else {
		if res.Error != nil {
			return nil, res.Error
		}
	}
	return &item, nil
}

func (f *ItemRepository) FindByEmailOrNil(email string) (*model.Item, error) {
	item := model.Item{}
	res := f.db.Find(&item, "email=?", email)
	if res.RecordNotFound() {
		return nil, nil
	} else {
		if res.Error != nil {
			return nil, res.Error
		}
	}
	return &item, nil
}

func (f *ItemRepository) FindAllByUserIDOrNil(userID string) ([]model.Item, error) {
	items := []model.Item{}
	res := f.db.Find(&items, "user_id=?", userID)
	if res.RecordNotFound() {
		return nil, nil
	} else {
		if res.Error != nil {
			return nil, res.Error
		}
	}
	return items, nil
}

func (f *ItemRepository) RegisterItem(name string, price int, iconImage string, description string) (*model.Item, error) {
	newItem := model.Item{}
	newItem.Name = name
	newItem.Price = price
	newItem.IconImage = iconImage
	newItem.Description = description

	if err := f.db.Create(&newItem).Error; err != nil {
		return nil, err
	}
	return &newItem, nil
}
