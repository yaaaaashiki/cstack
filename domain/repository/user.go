package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/yaaaaashiki/cstack/domain/model"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (f *UserRepository) FindAll() ([]model.User, error) {
	users := []model.User{}
	if err := f.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (f *UserRepository) FindByID(id int) (*model.User, error) {
	users := model.User{}
	if err := f.db.Find(&users, "id=?", id).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (f *UserRepository) FindByIDOrNil(id uint) (*model.User, error) {
	user := model.User{}
	res := f.db.Find(&user, "id=?", id)
	if res.RecordNotFound() {
		return nil, nil
	} else {
		if res.Error != nil {
			return nil, res.Error
		}
	}
	return &user, nil
}

func (f *UserRepository) FindByEmailOrNil(email string) (*model.User, error) {
	user := model.User{}
	res := f.db.Find(&user, "email=?", email)
	if res.RecordNotFound() {
		return nil, nil
	} else {
		if res.Error != nil {
			return nil, res.Error
		}
	}
	return &user, nil
}

func (f *UserRepository) RegisterUser(name string, email string, salt string, salted string, iconImage string) (*model.User, error) {
	newUser := model.User{}
	newUser.Name = name
	newUser.Email = email
	newUser.Salt = salt
	newUser.Salted = salted
	newUser.IconImage = iconImage

	if err := f.db.Create(&newUser).Error; err != nil {
		return nil, err
	}
	return &newUser, nil
}
