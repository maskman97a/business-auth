package controller_impl

import (
	"business-auth/conf/router"
	"business-auth/internal/constants"
	"business-auth/internal/constants/error_code"
	"business-auth/internal/controller"
	"business-auth/internal/dto/request"
	"business-auth/internal/dto/response"
	"business-auth/internal/service/service_impl"
	"business-auth/pkg/utils/json_utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type baseController struct {
}

func NewBaseController() controller.BaseController {
	return &baseController{}
}

func ValidateRequest(c *gin.Context) (*request.BaseRequest, error) {
	logrus.Info(fmt.Sprintf("--Start %s--", "baseController.ValidateRequest"))
	var baseRequest request.BaseRequest
	err := c.BindJSON(&baseRequest)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	if baseRequest.RequestTime == "" {
		return nil, errors.New("requestTime is empty")
	}
	_, err = time.Parse(constants.DateTimestampPattern, baseRequest.RequestTime)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	if baseRequest.IpAddress == "" {
		return nil, errors.New("ipAddress is empty")
	}
	if baseRequest.Signature == "" {
		return nil, errors.New("signature is empty")
	}
	if baseRequest.Data == "" {
		return nil, errors.New("data is empty")
	}
	if !json_utils.IsJSON(baseRequest.Data) {
		return nil, errors.New("data is invalid")
	}
	logrus.Info(fmt.Sprintf("--Finish %s--", "baseController.ValidateRequest"))
	return &baseRequest, nil
}

func InitRouter(gin *gin.Engine, mainGroup string, group string, path string, handler func(c *gin.Context), method string) {
	api := gin.Group(mainGroup).Group(group)
	if method == http.MethodPost {
		router.Post(api, path, handler)
	} else {
		router.Get(api, path, handler)
	}

}

func SetSignature(data string) (string, error) {
	signResp, err := service_impl.NewHsmService().Sign(request.SignRequest{Data: data})
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	return signResp.Signature, nil
}

type ServiceFunc func(any) error

func ExecuteFunc(f ServiceFunc, a any) error {
	return f(a)
}

func ExecuteService(c *gin.Context, respData any, reqData any, f ServiceFunc) {
	logrus.Info(fmt.Sprintf("--Start %s--", "authController.Login"))
	var resp response.BaseResponse
	baseRequest, err := ValidateRequest(c)
	var respCode int
	var respDesc string
	if err != nil || baseRequest == nil {
		logrus.Error(err)
		respCode = error_code.InvalidRequest
		respDesc = error_code.GetErrorMsg(respCode)
	} else {
		err = json_utils.ConvertToObject(baseRequest.Data, &reqData)
		if err != nil {
			logrus.Error(err)
			respCode = error_code.InvalidRequest
			respDesc = error_code.GetErrorMsg(respCode)
		} else {
			err = f(reqData)
			if err != nil {
				logrus.Error(err)
				respCode = error_code.Failed
				respDesc = error_code.GetErrorMsg(respCode)
			} else {
				respCode = error_code.Success
				respDesc = error_code.GetErrorMsg(respCode)
			}
		}
	}
	resp.Data, _ = json_utils.ConvertToString(respData)

	resp.ResponseTime = time.Now().Format(constants.DateTimestampPattern)
	signature, err := SetSignature(resp.Data)
	if err != nil {
		respCode = error_code.SystemError
		respDesc = error_code.GetErrorMsg(respCode)
	}
	c.JSON(http.StatusOK, response.NewBaseResponse(respCode, respDesc, signature, resp.Data))
	logrus.Info(fmt.Sprintf("--Finish %s--", "authController.Login"))
}
