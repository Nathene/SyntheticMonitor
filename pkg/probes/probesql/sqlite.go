package probesql

import (
	"context"
	"database/sql"
	"time"

	_ "modernc.org/sqlite"
)

type SQLite struct {
	*sql.DB
	ctx context.Context
}

func NewSQLite(ctx context.Context, dsn string) (*SQLite, error) {
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}
	return &SQLite{ctx: ctx, DB: db}, nil
}

func (s *SQLite) Connect() error {
	return s.DB.Ping()
}

func (s *SQLite) Query(query string) (*sql.Rows, error) {
	ctx, cancel := context.WithTimeout(s.ctx, 5*time.Second)
	defer cancel()
	s.PrepareContext(ctx, query)
	return s.DB.QueryContext(ctx, query)
}

func (s *SQLite) Close() error {
	return s.DB.Close()
}
