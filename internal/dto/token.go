package dto

type TokenDto struct {
	UserID      uint   `json:"userID"`
	Token       string `json:"token"`
	CreatedDate string `json:"createdDate"`
	ExpiredDate string `json:"expiredDate"`
}

func NewTokenDto() TokenDto {
	return TokenDto{}
}
