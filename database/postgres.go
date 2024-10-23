package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"time"
)

func NewPool(
	log *zerolog.Logger,
	uri *string,
	maxConnections *int32,
	minConnections *int32,
	maxConnLifetime *time.Duration,
	maxConnIdleTime *time.Duration,
) (*pgxpool.Pool, error) {
	poolConfig, err := pgxpool.ParseConfig(*uri)
	if err != nil {
		log.Err(err).Msg("Unable to parse connection string")
		return nil, err
	}
	poolConfig.MaxConns = *maxConnections
	poolConfig.MinConns = *minConnections
	poolConfig.MaxConnLifetime = *maxConnLifetime * time.Millisecond
	poolConfig.MaxConnIdleTime = *maxConnIdleTime * time.Millisecond
	ctx := context.Background()
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		log.Err(err).Msg("Unable to connect to database")

		return nil, err
	}
	return pool, nil
}
