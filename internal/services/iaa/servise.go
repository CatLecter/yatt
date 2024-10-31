package service

import (
	repository "github.com/CatLecter/yatt/internal/infrastructure/repositories/iaa"
	"github.com/rs/zerolog"
)

type Service struct {
	UserServiceInterface
}

func New(log *zerolog.Logger, userStorage repository.UserStorageInterface) *Service {
	return &Service{
		UserServiceInterface: NewUserService(log, userStorage),
	}
}
