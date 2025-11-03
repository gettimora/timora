package sqlkit

import (
	"context"
	"database/sql"
	"log/slog"
)

var _ DB = (*SQLDatabase)(nil)

type (
	DB interface {
		Dialect() Dialect
		ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
		QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
		Close() error

		Insert(t Table) InsertQuery
	}

	SQLDatabase struct {
		sql     *sql.DB
		dialect Dialect
		config  SQLDatabaseConfig
	}

	SQLDatabaseConfig struct {
		Log       Logger
		LogConfig LogConfig
	}

	LogConfig struct {
		Enabled     bool
		IncludeArgs bool
		LogDuration bool
	}

	SQLDatabaseOption func(*SQLDatabaseConfig)

	Logger interface {
		Info(msg string, kv ...any)
	}
)

func New(raw *sql.DB, dialect Dialect, opts ...SQLDatabaseOption) DB {
	return NewSQLDatabase(raw, dialect, opts...)
}

func NewSQLDatabase(raw *sql.DB, dialect Dialect, opts ...SQLDatabaseOption) *SQLDatabase {
	cfg := defaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}
	db := &SQLDatabase{
		sql:     raw,
		dialect: dialect,
		config:  cfg,
	}
	return db
}

func WithLogger(logger Logger) SQLDatabaseOption {
	return func(cfg *SQLDatabaseConfig) {
		if logger == nil {
			panic("sqlkit: WithLogger(nil) is not allowed; use slog.Default() or a valid Logger")
		}
		cfg.Log = logger
	}
}

func WithLogConfig(logConfig LogConfig) SQLDatabaseOption {
	return func(cfg *SQLDatabaseConfig) {
		cfg.LogConfig = logConfig
	}
}

func defaultConfig() SQLDatabaseConfig {
	return SQLDatabaseConfig{
		Log: slog.Default(),
		LogConfig: LogConfig{
			Enabled:     false,
			IncludeArgs: false,
			LogDuration: false,
		},
	}
}

func (db *SQLDatabase) Dialect() Dialect { return db.dialect }

func (db *SQLDatabase) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	ctx = withOperationStart(ctx)
	res, err := db.sql.ExecContext(ctx, query, args...)
	db.logQuery(ctx, "Running DB Execution", query, args, err)
	return res, err
}

func (db *SQLDatabase) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	ctx = withOperationStart(ctx)
	rows, err := db.sql.QueryContext(ctx, query, args...)
	db.logQuery(ctx, "Running DB Query", query, args, err)
	return rows, err
}

func (db *SQLDatabase) Close() error {
	return db.sql.Close()
}
