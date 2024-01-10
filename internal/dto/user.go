package dto

type UserDto struct {
	User        string `json:"user"`
	Pwd         string `json:"pwd"`
	Fullname    string `json:"fullname"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	Gender      string `json:"gender"`
	Email       string `json:"email"`
	Status      string `json:"status"`
	CreatedDate string
	CreatedBy   string
	UpdatedDate string
	UpdatedBy   string
}

func NewUserDTO() *UserDto {
	return &UserDto{}
}
