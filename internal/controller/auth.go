package controller

import (
	"business-auth/config/router"
	"business-auth/internal/constants"
	"business-auth/internal/constants/error_code"
	"business-auth/internal/dto/request"
	"business-auth/internal/dto/response"
	"business-auth/internal/service"
	"business-auth/pkg/utils/json_utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
	"net/http"
	"time"
)

type AuthController interface {
	BaseController
	Token(c *gin.Context)
	SignUp(c *gin.Context)
	Login(c *gin.Context)
}
type authController struct {
	authService  service.AuthService
	tokenService service.TokenService
}

func NewAuthController(db *gorm.DB) AuthController {
	return &authController{authService: service.NewAuthService(db), tokenService: service.NewTokenService(db)}
}

func (authController *authController) InitRouter(routerGroup *gin.RouterGroup) {
	api := routerGroup.Group("/auth")
	router.Post(api, "/token", authController.Token)
	router.Post(api, "/login", authController.Login)
	router.Post(api, "/signup", authController.SignUp)
}

func (authController *authController) Token(c *gin.Context) {
	var tokenResponse response.TokenResponse
	clientID := c.GetHeader("client-id")
	if clientID == "" {
		c.JSON(http.StatusUnauthorized, response.NewBaseResponse(
			error_code.InvalidRequest,
			error_code.GetErrorMsg(error_code.InvalidRequest),
			"", ""))
		return
	} else {
		trustedClientId := authController.tokenService.GetTrustedClientID()
		if !utils.Contains(trustedClientId, clientID) {
			c.JSON(http.StatusUnauthorized, response.BaseResponse{Code: error_code.InvalidRequest,
				Message: error_code.GetErrorMsg(error_code.InvalidRequest),
			})
			return
		}
	}
	clientSecret := c.GetHeader("client-secret")
	if clientSecret == "" {
		trustedClientId := authController.tokenService.GetTrustedClientID()
		if !utils.Contains(trustedClientId, clientID) {
			c.JSON(http.StatusUnauthorized, response.BaseResponse{Code: error_code.InvalidRequest,
				Message: error_code.GetErrorMsg(error_code.InvalidRequest),
			})
			return
		}

	}
	authenticationToken, err := authController.tokenService.CreateToken(clientID)
	if err != nil {
		c.JSON(http.StatusOK, response.BaseResponse{Code: error_code.SystemError,
			Message: error_code.GetErrorMsg(error_code.SystemError),
		})
		return
	} else {
		tokenResponse.Token = authenticationToken.Token
		tokenResponse.ExpiredDate = authenticationToken.ExpiredDate
		tokenRespStr, err := json_utils.ConvertToString(tokenResponse)
		if err != nil {
			c.JSON(http.StatusOK, response.BaseResponse{Code: error_code.SystemError,
				Message: error_code.GetErrorMsg(error_code.SystemError),
			})
		}
		c.JSON(http.StatusOK, response.BaseResponse{Code: error_code.Success,
			Data:    tokenRespStr,
			Message: error_code.GetErrorMsg(error_code.Success),
		})
	}
	return
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
			_, err := authController.authService.Login(loginRequest)
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
