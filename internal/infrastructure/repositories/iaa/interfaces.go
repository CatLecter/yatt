package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	domain "yatt/internal/domain/user"
)

type UserStorageInterface interface {
	Save(ctx *context.Context, tx pgx.Tx, user *domain.UserModel) error
	GetByID(ctx *context.Context, tx pgx.Tx, user *domain.UserModel) error
	GetByUserName(ctx *context.Context, tx pgx.Tx, user *domain.UserModel) error
	Update(ctx *context.Context, tx pgx.Tx, user *domain.UserModel) error
	CheckByEmail(ctx *context.Context, tx pgx.Tx, email string) (bool, error)
	CheckByUserName(ctx *context.Context, tx pgx.Tx, username string) (bool, error)
}
