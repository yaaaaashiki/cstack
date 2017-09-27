package usecase

import (
	"github.com/yaaaaashiki/cstack/domain/repository"
	"github.com/yaaaaashiki/cstack/helper"
)

type LoginUseCase struct {
	userRepository *repository.UserRepository
}

type LoginRequest struct {
	Email    string
	Password string
}

type LoginResponse struct {
	Success bool
	Error   *helper.Error
	UserId  uint
}

func NewLoginUseCase(eventRepository *repository.UserRepository) *LoginUseCase {
	return &LoginUseCase{
		userRepository: eventRepository,
	}
}

func (f *LoginUseCase) Execute(request *LoginRequest) (response *LoginResponse, error error) { //TODO test
	userId, salt, salted, err := f.userRepository.GetPassword(request.Email)
	if err != nil {
		return nil, err
	}
	if salt == "" {
		return &LoginResponse{Success: false, Error: helper.NewError(ErrorUserNotFound, "user not found")}, nil
	}

	stretched := helper.Stretch(request.Password, salt)
	if stretched != salted {
		return &LoginResponse{Success: false, Error: helper.NewError(ErrorIncorrectPassword, "incorrect password")}, nil
	}

	return &LoginResponse{Success: true, Error: nil, UserId: userId}, nil
}
