package response

type SignUpResponse struct {
	EffectDate string `json:"effectDate"`
}

func NewSignUpResponse() *SignUpResponse {
	return &SignUpResponse{}
}
