package response

type LoginResponse struct {
	Token       string `json:"token"`
	ExpiredDate string `json:"expiredDate"`
	UserInfo    string `json:"userInfo"`
}

func NewLoginResponse() LoginResponse {
	return LoginResponse{}
}
