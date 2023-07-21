package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
)

type StoreConfig interface {
	GetDSN() string
}

type DB struct {
	*pgxpool.Pool
}

func NewDatabase(cfg StoreConfig) (*DB, error) {
	dsn := cfg.GetDSN()
	config, err := pgxpool.ParseConfig(dsn)

	if err != nil {
		return nil, err
	}

	config.ConnConfig.Tracer = &tracelog.TraceLog{
		LogLevel: tracelog.LogLevelTrace,
	}

	conn, err := pgxpool.NewWithConfig(context.Background(), config)

	if err != nil {
		return nil, err
	}

	return &DB{conn}, nil
}
