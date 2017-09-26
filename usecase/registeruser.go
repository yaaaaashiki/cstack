package usecase

import (
	"errors"

	"github.com/yaaaaashiki/cstack/domain/model"
	"github.com/yaaaaashiki/cstack/domain/repository"
)

type RegisterUseCase struct {
	userRepository *repository.UserRepository
}

type RegisterRequest struct {
	ID        uint   `json:"id" gorm:"column:id"`
	Name      string `binding:"required" json:"name"`
	Email     string `binding:"required" json:"email"`
	Salt      string `json:"salt" gorm:"column:salt"`
	Salted    string `json:"salted" gorm:"column:salted"`
	IconImage string `binding:"required" json:"icon_image"`
}

type RegisterResponse struct {
	User *model.User
}

func NewRegisterUseCase(userRepository *repository.UserRepository) *RegisterUseCase {
	return &RegisterUseCase{
		userRepository: userRepository,
	}
}

//Before running this function, should be judged new user or usual user.
//If new user input the data, run this function
func (f *RegisterUseCase) Execute(req *RegisterRequest) (*RegisterResponse, error) { //TODO test

	isUser, err := f.userRepository.FindByEmailOrNil(req.Email)
	if err != nil {
		return nil, err
	}

	if isUser != nil {
		return nil, errors.New("Duplicate user cannnot create new user")
	}

	user, err := f.userRepository.RegisterUser(req.Name, req.Email, req.Salt, req.Salted, req.IconImage)
	if err != nil {
		return nil, err
	}

	return &RegisterResponse{User: user}, nil
}
