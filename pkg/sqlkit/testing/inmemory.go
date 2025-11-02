package testing

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gettimora/timora/pkg/sqlkit"
	"github.com/gettimora/timora/pkg/sqlkit/sqlite"
)

type (
	Option func(*testing.T, *sqlkit.DB)
)

func NewInMemoryTest(t *testing.T, opts ...Option) *sqlkit.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic("failed to open in-memory sqlite database: " + err.Error())
	}
	DB := sqlkit.NewDB(db, sqlite.NewDialect())

	_, err = DB.ExecContext(t.Context(), CreateSchema)
	if err != nil {
		t.Fatal("failed to create schema: " + err.Error())
	}

	for _, opt := range opts {
		opt(t, DB)
	}

	return DB
}

func WithExample(example Example) Option {
	return func(t *testing.T, DB *sqlkit.DB) {
		_, err := DB.
			Insert(Examples).
			Columns(
				Examples.ID,
				Examples.Text,
				Examples.Unique,
				Examples.Value,
				Examples.CreatedAt,
				Examples.IsActive,
			).
			Values(
				example.ID,
				example.Text,
				example.Unique,
				example.Value,
				example.CreatedAt,
				example.IsActive,
			).
			Exec(t.Context())
		if err != nil {
			t.Fatal("failed to insert example: " + err.Error())
		}
	}
}
