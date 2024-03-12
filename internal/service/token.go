package service

import (
	"business-auth/internal/dto"
	"business-auth/internal/repository"
	"business-auth/internal/repository/repository_impl"
	"business-auth/pkg/utils/token_utils"
	"gorm.io/gorm"
)

type TokenService interface {
	GetTrustedClientID() []string
	CreateToken(clientID string) (dto.AuthenticationToken, error)
	VerifyToken(clientID int64, token string)
}

type tokenService struct {
	repository repository.TokenRepo
}

func NewTokenService(db *gorm.DB) TokenService {
	return &tokenService{repository: repository_impl.NewTokenRepo(db)}
}

func (tokenService *tokenService) CreateToken(clientID string) (dto.AuthenticationToken, error) {
	authenticationToken, err := token_utils.GenerateJWT(clientID)
	if err != nil {
		return authenticationToken, err
	} else {
		return authenticationToken, nil
	}
}
func (tokenService *tokenService) VerifyToken(clientID int64, token string) {

}

func (tokenService *tokenService) GetTrustedClientID() []string {
	return []string{"123", "124"}
}
