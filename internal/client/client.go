package client

import (
	"context"
	"ftx/internal/config"
)

type Client struct {
	Postgres *PostgresClient
	Redis    *RedisClient
}

func New(
	ctx context.Context,
	cfg *config.ClientConfig,
) (*Client, error) {
	pg, err := newPostgresClient(ctx, *cfg.Postgres)
	if err != nil {
		return nil, err
	}

	rds, err := newRedisClient(ctx, cfg.Redis)
	if err != nil {
		return nil, err
	}

	return &Client{
		Postgres: pg,
		Redis:    rds,
	}, nil
}

func (c *Client) Close() {
	c.Postgres.Close()
	c.Redis.Close()
}
