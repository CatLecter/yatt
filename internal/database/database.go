package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

func MustNew(
	ctx context.Context,
	uri string,
	maxConnections int32,
	minConnections int32,
	maxConnLifetime time.Duration,
	maxConnIdleTime time.Duration,
) *pgxpool.Pool {
	poolConfig, err := pgxpool.ParseConfig(uri)
	if err != nil {
		panic(err)
	}
	poolConfig.MaxConns = maxConnections
	poolConfig.MinConns = minConnections
	poolConfig.MaxConnLifetime = maxConnLifetime
	poolConfig.MaxConnIdleTime = maxConnIdleTime
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		panic(err)
	}

	if err = pool.Ping(ctx); err != nil {
		panic(err)
	}
	return pool
}
