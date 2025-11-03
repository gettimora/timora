package sqlkittesting

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gettimora/timora/pkg/sqlkit"
	"github.com/gettimora/timora/pkg/sqlkit/sqlite"
)

type (
	InMemoryTestDB struct {
		sqlkit.DB
	}

	Option func(*testing.T, sqlkit.DB)
)

func NewExampleTestDB(opts ...sqlkit.SQLDatabaseOption) *InMemoryTestDB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic("failed to open in-memory sqlite database: " + err.Error())
	}
	DB := sqlkit.New(db, sqlite.NewDialect(), opts...)
	imt := &InMemoryTestDB{DB: DB}

	_, err = DB.ExecContext(context.Background(), CreateSchema)
	if err != nil {
		panic("failed to create schema: " + err.Error())
	}

	return imt
}

func NewInMemoryTestDB(t *testing.T, opts ...Option) *InMemoryTestDB {
	t.Helper()

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic("failed to open in-memory sqlite database: " + err.Error())
	}
	DB := sqlkit.New(db, sqlite.NewDialect(), sqlkit.WithLogConfig(
		sqlkit.LogConfig{
			Enabled:     true,
			IncludeArgs: true,
			LogDuration: true,
		},
	))
	imt := &InMemoryTestDB{DB: DB}
	t.Cleanup(func() {
		imt.Close(t)
	})

	_, err = DB.ExecContext(t.Context(), CreateSchema)
	if err != nil {
		t.Fatal("failed to create schema: " + err.Error())
	}

	for _, opt := range opts {
		opt(t, DB)
	}

	return imt
}

func (it *InMemoryTestDB) GetExample(t *testing.T, id int) (*Example, error) {
	rows, err := it.QueryContext(
		context.Background(),
		`SELECT "id", "text", "value", "unique", "created_at", "is_active" FROM "examples" WHERE "id" = ?`,
		id,
	)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		return nil, err
	}
	if err != nil {
		t.Fatalf("failed to query example: %v", err)
	}
	defer rows.Close()

	var example Example
	var createdAtInt int64
	var isActiveInt int

	if !rows.Next() {
		return nil, sql.ErrNoRows
	}
	err = rows.Scan(
		&example.ID,
		&example.Text,
		&example.Value,
		&example.Unique,
		&createdAtInt,
		&isActiveInt,
	)
	if err != nil {
		return &example, err
	}

	example.CreatedAt = time.Unix(createdAtInt, 0)
	example.IsActive = isActiveInt != 0

	return &example, nil
}

func (it *InMemoryTestDB) AssertExample(t *testing.T, id int, fn func(t *testing.T, example Example)) {
	t.Helper()
	example, err := it.GetExample(t, id)
	if err != nil {
		t.Fatalf("failed to get example with id %d: %v", id, err)
	}
	fn(t, *example)
}

func (it *InMemoryTestDB) Close(t *testing.T) {
	t.Helper()

	if err := it.DB.Close(); err != nil {
		t.Fatalf("failed to close in-memory database: %v", err)
	}
}

func WithExampleData(examples ...*Example) Option {
	return func(t *testing.T, db sqlkit.DB) {
		t.Helper()

		insert := db.Insert(Examples).Columns(
			Examples.ID,
			Examples.Text,
			Examples.Value,
			Examples.Unique,
			Examples.CreatedAt,
			Examples.IsActive,
		)

		for _, ex := range examples {
			insert.Values(
				ex.ID,
				ex.Text,
				ex.Value,
				ex.Unique,
				ex.CreatedAt.Unix(),
				ex.IsActive,
			)
		}
		insert.Exec(t.Context())
	}
}
