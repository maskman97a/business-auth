package request

type SignUpRequest struct {
	User        string `json:"user"`
	Pwd         string `json:"pwd"`
	Fullname    string `json:"fullname"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	Gender      string `json:"gender"`
}
