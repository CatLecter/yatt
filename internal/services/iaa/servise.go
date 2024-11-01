package service

import (
	"github.com/rs/zerolog"
	repository "yatt/internal/infrastructure/repositories/iaa"
)

type Service struct {
	UserServiceInterface
}

func New(log *zerolog.Logger, userStorage repository.UserStorageInterface) *Service {
	return &Service{
		UserServiceInterface: NewUserService(log, userStorage),
	}
}
