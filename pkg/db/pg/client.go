package pg

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"

	"github.com/robertobadjio/platform-common/pkg/db"
)

type pgClient struct {
	masterDBC    db.DB
	queryTimeout time.Duration
}

func New(ctx context.Context, dsn string, queryTimeout time.Duration) (db.Client, error) {
	dbc, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, errors.Errorf("failed to connect to db: %v", err)
	}

	return &pgClient{
		masterDBC: &pg{
			dbc:          dbc,
			queryTimeout: queryTimeout,
		},
	}, nil
}

func (c *pgClient) DB() db.DB {
	return c.masterDBC
}

func (c *pgClient) Close() error {
	if c.masterDBC != nil {
		c.masterDBC.Close()
	}

	return nil
}
