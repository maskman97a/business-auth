package model

import (
	"gorm.io/gorm"
	"time"
)

type TblToken struct {
	gorm.Model
	UserID      uint
	Token       string
	CreatedDate time.Time
	ExpiredDate time.Time
}

func NewTblToken() *TblToken {
	return &TblToken{}
}
