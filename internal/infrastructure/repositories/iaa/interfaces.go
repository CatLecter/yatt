package repository

import (
	"context"
	domain "yatt/internal/domain/user"
)

type UserStorageInterface interface {
	Save(ctx *context.Context, user *domain.UserModel) error
	GetByID(ctx *context.Context, user *domain.UserModel) error
	GetByUserName(ctx *context.Context, user *domain.UserModel) error
	Update(ctx *context.Context, user *domain.UserModel) error
	CheckByEmail(ctx *context.Context, email string) (bool, error)
	CheckByUserName(ctx *context.Context, username string) (bool, error)
}
