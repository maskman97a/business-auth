package response

type TokenResponse struct {
	Token       string `json:"token"`
	ExpiredDate string `json:"expiredDate"`
}
