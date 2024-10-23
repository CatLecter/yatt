package services

import (
	identitygrpc "github.com/CatLecter/yatt/internal/infrastructure/clients/identity/grpc"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type AuthService struct {
	client *identitygrpc.Client
	log    *zerolog.Logger
}

func NewAuthService(client *identitygrpc.Client, log *zerolog.Logger) AuthServiceInterface {
	return &AuthService{client: client, log: log}
}

func (s *AuthService) Register(ctx *gin.Context) {
	panic("implement me")
}

func (s *AuthService) Login(ctx *gin.Context) {
	panic("implement me")
}
