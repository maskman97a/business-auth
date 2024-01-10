package controller

import (
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	BaseController

	SignUp(c *gin.Context)
	Login(c *gin.Context)
}
