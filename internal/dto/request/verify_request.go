package request

type VerifyRequest struct {
	Data      string `json:"data"`
	Algorithm string `json:"algorithm"`
	Signature string `json:"signature"`
}
