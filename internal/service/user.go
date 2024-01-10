package service

import (
	"business-auth/internal/dto"
	"business-auth/internal/model"
)

type UserService interface {
	CheckAvailable(dto dto.UserDto) (bool, error)
	GetByUsernameEmailOrPhoneNumber(username string, email string, phoneNumber string) *model.TblUser
	CreateNew(dto dto.UserDto) (*model.TblUser, error)
}
