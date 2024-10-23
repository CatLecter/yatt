package services

import (
	identitygrpc "github.com/CatLecter/yatt/internal/infrastructure/clients/identity/grpc"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type AuthServiceInterface interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type UserServiceInterface interface {
	Create(ctx *gin.Context)
	Get(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type Service struct {
	AuthServiceInterface
	UserServiceInterface
}

func New(identityClient *identitygrpc.Client, log *zerolog.Logger) *Service {
	return &Service{
		AuthServiceInterface: NewAuthService(identityClient, log),
		UserServiceInterface: NewUserService(identityClient, log),
	}
}
