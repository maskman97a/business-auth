package response

type HsmSignResponse struct {
	Signature string `json:"signature"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
}
