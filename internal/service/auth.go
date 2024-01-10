package service

import (
	"business-auth/internal/dto/request"
	"business-auth/internal/dto/response"
)

type AuthService interface {
	SignUp(signUpRequest request.SignUpRequest) error
	Login(loginRequest request.LoginRequest) (*response.LoginResponse, error)
}
