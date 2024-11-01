package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"time"
	domainUser "yatt/internal/domain/user"
)

type UserRepository struct {
	db  *pgxpool.Pool
	log *zerolog.Logger
}

func NewUserRepository(log *zerolog.Logger, db *pgxpool.Pool) UserStorageInterface {
	return &UserRepository{db: db, log: log}
}

func (repo *UserRepository) Save(ctx *context.Context, user *domainUser.UserModel) error {
	conn, err := repo.db.Acquire(*ctx)
	defer conn.Release()
	if err != nil {
		repo.log.Warn().Err(err).Msg("error acquiring connection")
		return err
	}
	query := `INSERT INTO users.user (username, full_name, email, password, last_login) 
			  VALUES ($1, $2, $3, $4, $5) 
			  RETURNING *`
	user.Mu.Lock()
	defer user.Mu.Unlock()
	err = conn.QueryRow(
		*ctx, query, &user.UserName, &user.FullName, &user.Email, &user.Password, &time.Time{},
	).Scan(
		&user.ID, &user.UserName, &user.FullName, &user.Email, &user.Password, &user.Active,
		&user.Hidden, &user.LastLogin, &user.CustomFields, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		repo.log.Warn().Err(err).Msg("failed to insert user")
		return err
	}
	return nil
}

func (repo *UserRepository) GetByID(ctx *context.Context, user *domainUser.UserModel) error {
	conn, err := repo.db.Acquire(*ctx)
	defer conn.Release()
	if err != nil {
		repo.log.Warn().Err(err).Msg("error acquiring connection")
		return err
	}
	user.Mu.Lock()
	defer user.Mu.Unlock()
	row := conn.QueryRow(*ctx, "SELECT * FROM users.user WHERE user_id = $1", &user.ID)
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

func (repo *UserRepository) GetByUserName(ctx *context.Context, user *domainUser.UserModel) error {
	conn, err := repo.db.Acquire(*ctx)
	defer conn.Release()
	if err != nil {
		repo.log.Warn().Err(err).Msg("error acquiring connection")
		return err
	}
	user.Mu.Lock()
	defer user.Mu.Unlock()
	row := conn.QueryRow(*ctx, "SELECT * FROM users.user WHERE username = $1", &user.UserName)
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

func (repo *UserRepository) Update(ctx *context.Context, user *domainUser.UserModel) error {
	conn, err := repo.db.Acquire(*ctx)
	defer conn.Release()
	if err != nil {
		repo.log.Warn().Err(err).Msg("error acquiring connection")
		return err
	}
	query := `UPDATE users.user SET 
                  username = $1,
                  full_name = $2,
                  email = $3,
                  custom_fields = custom_fields || $4::jsonb,
                  updated_at = current_timestamp
              WHERE user_id = $5 
              RETURNING *`
	user.Mu.Lock()
	defer user.Mu.Unlock()
	err = conn.QueryRow(*ctx, query, &user.UserName, &user.FullName, &user.Email, &user.CustomFields, &user.ID).Scan(
		&user.ID, &user.UserName, &user.FullName, &user.Email, &user.Password, &user.Active,
		&user.Hidden, &user.LastLogin, &user.CustomFields, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		repo.log.Warn().Err(err).Msg("failed to update user")
		return err
	}
	return nil
}

func (repo *UserRepository) CheckByEmail(ctx *context.Context, email string) (bool, error) {
	conn, err := repo.db.Acquire(*ctx)
	defer conn.Release()
	if err != nil {
		repo.log.Warn().Err(err).Msg("error acquiring connection")
		return true, err
	}
	row := conn.QueryRow(
		*ctx,
		"SELECT CASE WHEN EXISTS (SELECT user_id FROM users.user WHERE email = $1) THEN true ELSE false END",
		&email,
	)
	var result bool
	err = row.Scan(&result)
	if err != nil {
		repo.log.Warn().Err(err).Msg("failed to check user")
		return true, err
	}
	return result, nil
}

func (repo *UserRepository) CheckByUserName(ctx *context.Context, username string) (bool, error) {
	conn, err := repo.db.Acquire(*ctx)
	defer conn.Release()

	if err != nil {
		repo.log.Warn().Err(err).Msg("error acquiring connection")
		return true, err
	}
	row := conn.QueryRow(
		*ctx,
		"SELECT CASE WHEN EXISTS (SELECT user_id FROM users.user WHERE username = $1) THEN true ELSE false END",
		&username,
	)
	var result bool
	err = row.Scan(&result)
	if err != nil {
		repo.log.Warn().Err(err).Msg("failed to check user")
		return true, err
	}
	return result, nil
}
