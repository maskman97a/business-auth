package dto

type AuthenticationToken struct {
	Token       string `json:"token"`
	ExpiredDate string `json:"expiredDate"`
}
