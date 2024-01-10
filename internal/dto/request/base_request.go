package request

type BaseRequest struct {
	RequestTime string `json:"requestTime"`
	IpAddress   string `json:"ipAddress"`
	Data        string `json:"data"`
	Signature   string `json:"signature"`
}
