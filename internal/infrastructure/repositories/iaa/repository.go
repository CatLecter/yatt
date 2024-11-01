package repository

import (
	"github.com/rs/zerolog"
)

type Repository struct {
	UserStorage UserStorageInterface
}

func New(log *zerolog.Logger) *Repository {
	return &Repository{
		UserStorage: NewUserRepository(log),
	}
}
