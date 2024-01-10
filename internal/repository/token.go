package repository

import (
	"business-auth/internal/dto"
	"business-auth/internal/model"
)

type TokenRepo interface {
	Insert(dto dto.TokenDto) (*model.TblToken, error)

	Update(dto dto.TokenDto) (*model.TblToken, error)
}
