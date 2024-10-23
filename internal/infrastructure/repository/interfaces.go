package repository

import (
	"github.com/CatLecter/yatt/internal/schemes"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type UserRepositoryInterface interface {
	CreateUser(user *schemes.UserRequest) (*schemes.UserResponse, error)
	GetUserByID(userID *uuid.UUID) (*schemes.UserResponse, error)
	CheckUserByEmail(email *string) (*bool, error)
}

type Repository struct{ UserRepositoryInterface }

func NewRepository(db *pgxpool.Pool, log *zerolog.Logger) *Repository {
	return &Repository{
		UserRepositoryInterface: NewUserRepository(db, log),
	}
}
