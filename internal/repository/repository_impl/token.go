package repository_impl

import (
	"business-auth/internal/dto"
	"business-auth/internal/model"
	"business-auth/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type tokenRepo struct {
	gormDB *gorm.DB
}

func NewTokenRepo(gormDB *gorm.DB) repository.TokenRepo {
	return &tokenRepo{gormDB: gormDB}
}

func (tokenRepo *tokenRepo) Insert(tokenDto dto.TokenDto) (*model.TblToken, error) {
	user := &model.TblToken{
		UserID:      tokenDto.UserID,
		Token:       tokenDto.Token,
		CreatedDate: time.Now(),
		ExpiredDate: time.Now().Add(time.Hour * 24),
	}

	err := tokenRepo.gormDB.AutoMigrate(&model.TblToken{})
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	result := tokenRepo.gormDB.Create(user)
	if result.Error != nil {
		logrus.Error(result.Error)
		return nil, result.Error
	}
	return user, nil
}

func (*tokenRepo) Update(dto dto.TokenDto) (*model.TblToken, error) {
	return nil, nil
}
