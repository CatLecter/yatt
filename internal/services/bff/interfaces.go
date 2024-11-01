package service

import (
	"github.com/gin-gonic/gin"
)

type UserServiceInterface interface {
	Create(ctx *gin.Context)
	Get(ctx *gin.Context)
	Update(ctx *gin.Context)
	Login(ctx *gin.Context)
}
