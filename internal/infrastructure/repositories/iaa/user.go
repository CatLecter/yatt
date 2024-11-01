package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
	"time"
	domainUser "yatt/internal/domain/user"
)

type UserRepository struct {
	log *zerolog.Logger
}

func NewUserRepository(log *zerolog.Logger) UserStorageInterface {
	return &UserRepository{log: log}
}

func (repo *UserRepository) Save(ctx *context.Context, tx pgx.Tx, user *domainUser.UserModel) error {
	query := `INSERT INTO users.user (username, full_name, email, password, last_login) 
			  VALUES ($1, $2, $3, $4, $5) 
			  RETURNING *`
	row := tx.QueryRow(
		*ctx, query, &user.UserName, &user.FullName, &user.Email, &user.Password, &time.Time{},
	)
	err := row.Scan(
		&user.ID, &user.UserName, &user.FullName, &user.Email, &user.Password, &user.Active,
		&user.Hidden, &user.LastLogin, &user.CustomFields, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		repo.log.Warn().Err(err).Msg("failed to insert user")
		return err
	}
	return nil
}

func (repo *UserRepository) GetByID(ctx *context.Context, tx pgx.Tx, user *domainUser.UserModel) error {
	query := `SELECT * FROM users.user WHERE user_id = $1`
	row := tx.QueryRow(*ctx, query, &user.ID)
	if err := row.Scan(
		&user.ID, &user.UserName, &user.FullName, &user.Email, &user.Password, &user.Active,
		&user.Hidden, &user.LastLogin, &user.CustomFields, &user.CreatedAt, &user.UpdatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errors.New("user not found")
		} else {
			repo.log.Warn().Err(err).Msg("failed to get user")
			return err
		}
	}
	return nil
}

func (repo *UserRepository) GetByUserName(ctx *context.Context, tx pgx.Tx, user *domainUser.UserModel) error {
	query := `SELECT * FROM users.user WHERE username = $1`
	row := tx.QueryRow(*ctx, query, &user.UserName)
	if err := row.Scan(
		&user.ID, &user.UserName, &user.FullName, &user.Email, &user.Password, &user.Active,
		&user.Hidden, &user.LastLogin, &user.CustomFields, &user.CreatedAt, &user.UpdatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errors.New("user not found")
		} else {
			repo.log.Warn().Err(err).Msg("failed to get user")
			return err
		}
	}
	return nil
}

func (repo *UserRepository) Update(ctx *context.Context, tx pgx.Tx, user *domainUser.UserModel) error {
	query := `UPDATE users.user SET 
                  username = $1,
                  full_name = $2,
                  email = $3,
                  custom_fields = custom_fields || $4::jsonb,
                  updated_at = current_timestamp
              WHERE user_id = $5 
              RETURNING *`
	err := tx.QueryRow(*ctx, query, &user.UserName, &user.FullName, &user.Email, &user.CustomFields, &user.ID).Scan(
		&user.ID, &user.UserName, &user.FullName, &user.Email, &user.Password, &user.Active,
		&user.Hidden, &user.LastLogin, &user.CustomFields, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		repo.log.Warn().Err(err).Msg("failed to update user")
		return err
	}
	return nil
}

func (repo *UserRepository) CheckByEmail(ctx *context.Context, tx pgx.Tx, email string) (bool, error) {
	query := `SELECT CASE WHEN EXISTS (SELECT user_id FROM users.user WHERE email = $1) THEN true ELSE false END`
	row := tx.QueryRow(*ctx, query, &email)
	var result bool
	err := row.Scan(&result)
	if err != nil {
		repo.log.Warn().Err(err).Msg("failed to check user")
		return true, err
	}
	return result, nil
}

func (repo *UserRepository) CheckByUserName(ctx *context.Context, tx pgx.Tx, username string) (bool, error) {
	query := `SELECT CASE WHEN EXISTS (SELECT user_id FROM users.user WHERE username = $1) THEN true ELSE false END`
	row := tx.QueryRow(*ctx, query, &username)
	var result bool
	err := row.Scan(&result)
	if err != nil {
		repo.log.Warn().Err(err).Msg("failed to check user")
		return true, err
	}
	return result, nil
}
