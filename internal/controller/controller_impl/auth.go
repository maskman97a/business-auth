package controller_impl

import (
	"business-auth/internal/constants"
	"business-auth/internal/constants/error_code"
	"business-auth/internal/controller"
	"business-auth/internal/dto/request"
	"business-auth/internal/dto/response"
	"business-auth/internal/service"
	"business-auth/internal/service/service_impl"
	"business-auth/pkg/utils/json_utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type authController struct {
	authService  service.AuthService
	tokenService service.TokenService
	gin          *gin.Engine
}

func NewAuthController(gin *gin.Engine, db *gorm.DB) controller.AuthController {
	return &authController{gin: gin, authService: service_impl.NewAuthService(db), tokenService: service_impl.NewTokenService(db)}
}

func (authController *authController) SignUp(c *gin.Context) {
	logrus.Info(fmt.Sprintf("--Start %s--", "authController.SignUp"))
	var resp response.BaseResponse
	var respData response.SignUpResponse
	baseRequest, err := ValidateRequest(c)
	var respCode int
	var respDesc string
	if err != nil || baseRequest == nil {
		logrus.Error(err)
		respCode = error_code.InvalidRequest
		respDesc = error_code.GetErrorMsg(respCode)
	} else {
		var signUpRequest request.SignUpRequest
		err = json_utils.ConvertToObject(baseRequest.Data, &signUpRequest)
		if err != nil {
			logrus.Error(err)
			respCode = error_code.InvalidRequest
			respDesc = error_code.GetErrorMsg(respCode)
		} else {
			err = authController.authService.SignUp(signUpRequest)
			if err != nil {
				logrus.Error(err)
				respCode = error_code.Failed
				respDesc = error_code.GetErrorMsg(respCode)
			} else {
				respCode = error_code.Success
				respDesc = error_code.GetErrorMsg(respCode)

				respData.EffectDate = userDto.CreatedDate
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
	logrus.Info(fmt.Sprintf("--Finish %s--", "authController.SignUp"))
}

func (authController *authController) Login(c *gin.Context) {
	logrus.Info(fmt.Sprintf("--Start %s--", "authController.SignUp"))
	var resp response.BaseResponse
	var respData response.LoginResponse
	baseRequest, err := ValidateRequest(c)
	var respCode int
	var respDesc string
	if err != nil || baseRequest == nil {
		logrus.Error(err)
		respCode = error_code.InvalidRequest
		respDesc = error_code.GetErrorMsg(respCode)
	} else {
		var loginRequest request.LoginRequest
		err = json_utils.ConvertToObject(baseRequest.Data, &loginRequest)
		if err != nil {
			logrus.Error(err)
			respCode = error_code.InvalidRequest
			respDesc = error_code.GetErrorMsg(respCode)
		} else {
			userInfo, err := authController.authService.Login(userDto)
			if err != nil {
				logrus.Error(err)
				respCode = error_code.Failed
				respDesc = error_code.GetErrorMsg(respCode)
			} else {
				respCode = error_code.Success
				respDesc = error_code.GetErrorMsg(respCode)
				respData.EffectDate = userDto.CreatedDate
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
