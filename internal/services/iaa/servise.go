package service

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	repository "yatt/internal/infrastructure/repositories/iaa"
)

type Service struct {
	UserServiceInterface
}

func New(log *zerolog.Logger, userStorage repository.UserStorageInterface, db *pgxpool.Pool) *Service {
	return &Service{
		UserServiceInterface: NewUserService(log, userStorage, db),
	}
}
