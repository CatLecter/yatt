package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type Repository struct {
	UserStorageInterface
}

func New(db *pgxpool.Pool, log *zerolog.Logger) *Repository {
	return &Repository{
		UserStorageInterface: NewUserRepository(log, db),
	}
}
