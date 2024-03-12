package service

import (
	"business-auth/internal/constants"
	"business-auth/internal/dto"
	"business-auth/internal/dto/request"
	"business-auth/internal/dto/response"
	"business-auth/pkg/utils/json_utils"
	"business-auth/pkg/utils/time_utils"
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AuthService interface {
	SignUp(signUpRequest request.SignUpRequest) error
	Login(loginRequest request.LoginRequest) (*response.LoginResponse, error)
}

type authService struct {
	userService UserService
}

func NewAuthService(db *gorm.DB) AuthService {
	return &authService{userService: NewUserService(db)}
}

func (authService *authService) SignUp(signUpRequest request.SignUpRequest) error {
	var userDto dto.UserDto
	userDto.User = signUpRequest.User
	userDto.Pwd = signUpRequest.Pwd
	userDto.Email = signUpRequest.Pwd
	userDto.PhoneNumber = signUpRequest.Pwd
	userDto.Address = signUpRequest.Pwd
	userDto.Fullname = signUpRequest.Pwd
	isAvailable, err := authService.userService.CheckAvailable(userDto)
	if !isAvailable {
		return err
	}
	userCreated, err := authService.userService.CreateNew(userDto)
	if err != nil {
		logrus.Error(err)
		return err
	} else {
		userDto.CreatedDate = time_utils.ToString(userCreated.CreatedDate, constants.DateTimestampPattern)
	}
	return nil
}
func (authService *authService) Login(loginRequest request.LoginRequest) (*response.LoginResponse, error) {
	var userDto dto.UserDto
	userDto.User = loginRequest.Username
	userDto.Pwd = loginRequest.Password
	userFound := authService.userService.GetByUsernameEmailOrPhoneNumber(userDto.User, userDto.Email, userDto.PhoneNumber)
	if userFound == nil {
		return nil, errors.New("user not found")
	}
	if loginRequest.Password != userFound.Password {
		return nil, errors.New("your password was wrong")
	} else {
		repUserDTO := dto.NewUserDTO()
		repUserDTO.User = userFound.Username
		repUserDTO.Email = userFound.Email
		repUserDTO.PhoneNumber = userFound.PhoneNumber
		repUserDTO.Fullname = userFound.Fullname
		loginResp := response.NewLoginResponse()
		userInfoIsonStr, err := json_utils.ConvertToString(loginResp)
		if err != nil {
			logrus.Error(err)
			return &loginResp, nil
		}
		loginResp.UserInfo = userInfoIsonStr

		return &loginResp, nil
	}
}
