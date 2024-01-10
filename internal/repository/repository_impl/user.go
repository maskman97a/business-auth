package repository_impl

import (
	"business-auth/internal/constants"
	"business-auth/internal/dto"
	"business-auth/internal/model"
	"business-auth/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"log"
	"time"
)

type userRepo struct {
	gormDB *gorm.DB
}

func NewUserRepo(gormDB *gorm.DB) repository.UserRepo {
	return &userRepo{gormDB: gormDB}
}

func (userRepo *userRepo) Insert(dto dto.UserDto) (*model.TblUser, error) {
	user := &model.TblUser{
		Username:    dto.User,
		Password:    dto.Pwd,
		PhoneNumber: dto.PhoneNumber,
		Email:       dto.Email,
		Status:      constants.ActiveStatus,
		CreatedDate: time.Now(),
		CreatedBy:   constants.DefaultUser,
		UpdatedDate: time.Now(),
		UpdatedBy:   constants.DefaultUser,
	}

	err := userRepo.gormDB.AutoMigrate(&model.TblUser{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	result := userRepo.gormDB.Create(user)
	if result.Error != nil {
		logrus.Error(result.Error)
		return nil, result.Error
	}
	return user, nil
}

func (userRepo *userRepo) GetByUsernameEmailOrPhoneNumber(username string, email string, phoneNumber string) *model.TblUser {
	var tblUser model.TblUser
	userRepo.gormDB.Where("(username = ? OR email = ? OR phone_number) AND status = 1", username, email, phoneNumber).First(&tblUser)
	return &tblUser
}

func (*userRepo) Update() {

}

func (*userRepo) Delete() {

}

func (*userRepo) Select() {

}
