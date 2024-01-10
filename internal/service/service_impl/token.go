package service_impl

import (
	"business-auth/internal/repository"
	"business-auth/internal/repository/repository_impl"
	"business-auth/internal/service"
	"gorm.io/gorm"
)

type tokenService struct {
	repository repository.TokenRepo
}

func NewTokenService(db *gorm.DB) service.TokenService {
	return &tokenService{repository: repository_impl.NewTokenRepo(db)}
}
