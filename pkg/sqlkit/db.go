package sqlkit

import (
	"context"
	"database/sql"
)

type DB struct {
	sql     *sql.DB
	dialect Dialect
}

func NewDB(raw *sql.DB, dialect Dialect) *DB {
	return &DB{sql: raw, dialect: dialect}
}

func (db *DB) Dialect() Dialect { return db.dialect }

func (db *DB) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return db.sql.ExecContext(ctx, query, args...)
}

func (db *DB) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return db.sql.QueryContext(ctx, query, args...)
}
