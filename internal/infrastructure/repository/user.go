package repository

import (
	"context"
	"github.com/CatLecter/yatt/internal/schemes"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type UserRepository struct {
	db  *pgxpool.Pool
	log *zerolog.Logger
}

func NewUserRepository(db *pgxpool.Pool, log *zerolog.Logger) UserRepositoryInterface {
	return &UserRepository{db: db, log: log}
}

func (repo *UserRepository) CreateUser(user *schemes.UserRequest) (*schemes.UserResponse, error) {
	ctx := context.Background()
	conn, err := repo.db.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		repo.log.Warn().Err(err).Msg("Error acquiring connection")
		return nil, err
	}
	row := conn.QueryRow(
		ctx, "INSERT INTO public.user (username, email) VALUES ($1, $2) RETURNING *", &user.UserName, &user.Email,
	)
	userResp := schemes.UserResponse{}
	err = row.Scan(
		&userResp.UserID,
		&userResp.UserName,
		&userResp.Email,
		&userResp.LastLoginDate,
		&userResp.IsActive,
		&userResp.CreatedAt,
		&userResp.UpdatedAt,
	)
	if err != nil {
		repo.log.Warn().Err(err).Msg("Failed to insert user")
		return nil, err
	}
	return &userResp, nil
}

func (repo *UserRepository) GetUserByID(userID *uuid.UUID) (*schemes.UserResponse, error) {
	ctx := context.Background()
	conn, err := repo.db.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		repo.log.Warn().Err(err).Msg("Error acquiring connection")
		return nil, err
	}
	row := conn.QueryRow(ctx, "SELECT * FROM public.user WHERE user_id = $1", &userID)
	userResp := schemes.UserResponse{}
	err = row.Scan(
		&userResp.UserID,
		&userResp.UserName,
		&userResp.Email,
		&userResp.LastLoginDate,
		&userResp.IsActive,
		&userResp.CreatedAt,
		&userResp.UpdatedAt,
	)
	if err != nil {
		repo.log.Warn().Err(err).Msg("Failed to get user")
		return nil, err
	}
	return &userResp, nil
}

func (repo *UserRepository) CheckUserByEmail(email *string) (*bool, error) {
	ctx := context.Background()
	conn, err := repo.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		repo.log.Warn().Err(err).Msg("Error acquiring connection")
		return nil, err
	}
	row := conn.QueryRow(
		ctx, "SELECT CASE WHEN EXISTS (SELECT user_id FROM public.user WHERE email = $1) THEN true ELSE false END", &email,
	)
	var result bool
	err = row.Scan(&result)
	if err != nil {
		repo.log.Warn().Err(err).Msg("Failed to check user")
		return nil, err
	}
	return &result, nil
}
