package service

import (
	"github.com/rs/zerolog"
	grpcclient "yatt/internal/infrastructure/clients/bff/grpc"
)

type Service struct {
	UserServiceInterface
}

func New(log *zerolog.Logger, userClient *grpcclient.Client) *Service {
	return &Service{
		UserServiceInterface: NewUserService(log, userClient),
	}
}
