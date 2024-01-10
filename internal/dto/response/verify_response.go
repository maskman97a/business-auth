package response

type VerifyResponse struct {
	Verified bool   `json:"verified"`
	Code     int    `json:"code"`
	Message  string `json:"message"`
}
