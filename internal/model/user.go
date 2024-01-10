package model

import (
	"gorm.io/gorm"
	"time"
)

type TblUser struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey;size:16"`
	Username    string `gorm:"index,size:24,unique;type:text collate nocase"`
	Password    string `gorm:"size:200"`
	Status      string `gorm:"size:1"`
	Email       string `gorm:"index,size:24,unique;type:text collate nocase"`
	PhoneNumber string `gorm:"index,size:11,unique;type:text collate nocase"`
	Fullname    string `gorm:"size:200"`
	CreatedDate time.Time
	CreatedBy   string
	UpdatedDate time.Time
	UpdatedBy   string
}
