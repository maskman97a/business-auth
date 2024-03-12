package controller

import (
	"business-auth/config/router"
	"business-auth/internal/constants"
	"business-auth/internal/constants/error_code"
	"business-auth/internal/dto/request"
	"business-auth/internal/dto/response"
	"business-auth/internal/service"
	"business-auth/pkg/utils/json_utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type BaseController interface {
	InitRouter(routerGroup *gin.RouterGroup)
}

type baseController struct {
}

func NewBaseController() BaseController {
	return &baseController{}
}

func (baseController *baseController) InitRouter(routerGroup *gin.RouterGroup) {
	api := routerGroup.Group("/")
	router.Get(api, "/", GetService)
}

func GetService(c *gin.Context) {
	c.JSON(http.StatusOK, response.NewBaseResponse(error_code.Success,
		"success",
		"", "",
	))
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

func SetSignature(data string) (string, error) {
	signResp, err := service.NewHsmService().Sign(request.SignRequest{Data: data})
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	return signResp.Signature, nil
}
