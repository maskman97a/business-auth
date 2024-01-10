package service_impl

import (
	"business-auth/internal/dto"
	"business-auth/internal/model"
	"business-auth/internal/repository"
	"business-auth/internal/repository/repository_impl"
	"business-auth/internal/service"
	"errors"
	"gorm.io/gorm"
	"strings"
)

type userService struct {
	userRepo repository.UserRepo
}

func NewUserService(db *gorm.DB) service.UserService {
	return &userService{userRepo: repository_impl.NewUserRepo(db)}
}

func (userService *userService) CreateNew(dto dto.UserDto) (*model.TblUser, error) {
	return userService.userRepo.Insert(dto)
}
func (userService *userService) CheckAvailable(dto dto.UserDto) (bool, error) {
	userFound := userService.GetByUsernameEmailOrPhoneNumber(dto.User, dto.Email, dto.PhoneNumber)
	if userFound != nil {
		if strings.ToUpper(userFound.Username) == strings.ToUpper(dto.User) {
			return false, errors.New("username already exists")
		} else if strings.ToUpper(userFound.Email) == strings.ToUpper(dto.Email) {
			return false, errors.New("email already exists")
		} else if strings.ToUpper(userFound.PhoneNumber) == strings.ToUpper(dto.PhoneNumber) {
			return false, errors.New("phoneNumber already exists")
		}
	}
	return true, nil
}

func (userService *userService) GetByUsernameEmailOrPhoneNumber(username string, email string, phoneNumber string) *model.TblUser {
	userFound := userService.userRepo.GetByUsernameEmailOrPhoneNumber(username, email, phoneNumber)
	return userFound
}
