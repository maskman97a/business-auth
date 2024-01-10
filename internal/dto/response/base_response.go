package response

import (
	"business-auth/internal/constants"
	"business-auth/pkg/utils/time_utils"
	"time"
)

type BaseResponse struct {
	Code         int    `json:"code"`
	Message      string `json:"message"`
	ResponseTime string `json:"responseTime"`
	Signature    string `json:"signature"`
	Data         string `json:"data"`
}

func NewBaseResponse(code int, msg string, signature string, data string) *BaseResponse {
	return &BaseResponse{Code: code, Message: msg, Signature: signature, Data: data, ResponseTime: time_utils.ToString(time.Now(), constants.DateTimestampPattern)}
}
