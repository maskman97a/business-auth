package repository

import (
	"business-auth/internal/dto"
	"business-auth/internal/model"
)

type UserRepo interface {
	Insert(dto dto.UserDto) (*model.TblUser, error)

	Update()

	Delete()

	Select()

	GetByUsernameEmailOrPhoneNumber(username string, email string, phoneNumber string) *model.TblUser
}
