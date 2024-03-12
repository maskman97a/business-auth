package service

import (
	"business-auth/config/hsm"
	"business-auth/internal/constants/error_code"
	"business-auth/internal/dto/request"
	"business-auth/internal/dto/response"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
)

type HsmService interface {
	Sign(req request.SignRequest) (response.SignResponse, error)

	Verify(req request.VerifyRequest) (response.VerifyResponse, error)
}

type hsmService struct {
	config *hsm.Config
}

func NewHsmService() HsmService {
	return &hsmService{config: hsm.NewHsmConfig()}
}

func (hsmService *hsmService) Sign(signRequest request.SignRequest) (response.SignResponse, error) {
	logrus.Info(fmt.Sprintf("--Start %s--", "hsmService.Sign"))
	resp, err := getSignReponse(hsmService.config, signRequest.Data)
	var signResponse response.SignResponse

	if err != nil {
		logrus.Error(err)
		return signResponse, err
	}
	signResponse.Signature = resp.Signature
	signResponse.Code = error_code.Success
	signResponse.Message = error_code.MapErrorCode[signResponse.Code]
	logrus.Info(fmt.Sprintf("--Finish %s--", "hsmService.Sign"))
	return signResponse, nil
}

func getSignReponse(config *hsm.Config, data string) (response.HsmSignResponse, error) {
	var hsmSignRequest request.HsmSignRequest
	hsmSignRequest.Data = data
	hsmSignRequest.Algorithm = "SHA256withRSA"
	hsmSignResponse, err := callHsmSign(*config, hsmSignRequest)
	if err != nil {
		return hsmSignResponse, err
	}
	return hsmSignResponse, nil
}

func callHsmSign(config hsm.Config, hsmSignRequest request.HsmSignRequest) (response.HsmSignResponse, error) {
	logrus.Info(fmt.Sprintf("--Start %s--", "hsmService.callHsmSign"))
	var hsmSignResponse response.HsmSignResponse
	url := config.Url + config.SignEndpoint
	err := CallApi(url, http.MethodPost, hsmSignRequest, &hsmSignResponse)
	if err != nil {
		logrus.Error(err)
		return hsmSignResponse, err
	}
	logrus.Info(fmt.Sprintf("--Finish %s--", "hsmService.callHsmSign"))
	return hsmSignResponse, nil
}

func (hsmService *hsmService) Verify(verifyRequest request.VerifyRequest) (response.VerifyResponse, error) {
	var verifyResponse response.VerifyResponse
	var hsmVerifyRequest request.HsmVerifyRequest
	hsmVerifyRequest.Signature = verifyRequest.Signature
	hsmVerifyRequest.Data = verifyRequest.Data
	hsmVerifyRequest.Algorithm = verifyRequest.Algorithm
	hsmVerifyResponse, err := callHsmVerify(hsmService.config, hsmVerifyRequest)
	if err != nil {
		logrus.Error(err)
		return verifyResponse, err
	}
	verifyResponse.Verified = hsmVerifyResponse.Verified
	verifyResponse.Code = hsmVerifyResponse.Code
	verifyResponse.Message = hsmVerifyResponse.Message
	return verifyResponse, nil
}

func callHsmVerify(config *hsm.Config, hsmSignRequest request.HsmVerifyRequest) (response.HsmVerifyResponse, error) {
	var hsmVerifyResponse response.HsmVerifyResponse
	url := config.Url + config.VerifyEnpoint
	method := "POST"

	reqJson, err := json.Marshal(hsmSignRequest)
	payload := strings.NewReader(string(reqJson))
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		logrus.Info(err)
		return hsmVerifyResponse, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		logrus.Info(err)
		return hsmVerifyResponse, err
	}
	defer func() {
		err = res.Body.Close()
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		logrus.Info(err)
		return hsmVerifyResponse, err
	}

	err = json.Unmarshal(body, &hsmVerifyResponse)
	if err != nil {
		logrus.Error(err)
		return hsmVerifyResponse, err
	}
	return hsmVerifyResponse, nil
}
