package service

import (
	"business-auth/internal/dto/request"
	"business-auth/internal/dto/response"
)

type HsmService interface {
	Sign(req request.SignRequest) (response.SignResponse, error)

	Verify(req request.VerifyRequest) (response.VerifyResponse, error)
}
