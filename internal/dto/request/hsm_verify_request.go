package request

type HsmVerifyRequest struct {
	Data      string `json:"data"`
	Algorithm string `json:"algorithm"`
	Signature string `json:"signature"`
}
