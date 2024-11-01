package service

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	grpcclient "yatt/internal/infrastructure/clients/bff/grpc"
)

type UserService struct {
	log        *zerolog.Logger
	userClient *grpcclient.Client
}

func NewUserService(log *zerolog.Logger, userClient *grpcclient.Client) UserServiceInterface {
	return &UserService{log: log, userClient: userClient}
}

func (s *UserService) Create(ctx *gin.Context) { return }

func (s *UserService) Get(ctx *gin.Context) { return }

func (s *UserService) Update(ctx *gin.Context) { return }

func (s *UserService) Login(ctx *gin.Context) { return }
