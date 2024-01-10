package request

type HsmSignRequest struct {
	Data      string `json:"data"`
	Algorithm string `json:"algorithm"`
}
