package sqlkit_test

import (
	"strings"
	"testing"
	"time"

	sqlkittesting "github.com/gettimora/timora/pkg/sqlkit/testing"
)

func TestInsert(t *testing.T) {
	t.Run("Insert single row", func(t *testing.T) {
		DB := sqlkittesting.NewInMemoryTestDB(t)

		res, err := DB.
			Insert(sqlkittesting.Examples).
			Columns(
				sqlkittesting.Examples.ID,
				sqlkittesting.Examples.Text,
				sqlkittesting.Examples.Value,
				sqlkittesting.Examples.Unique,
				sqlkittesting.Examples.CreatedAt,
				sqlkittesting.Examples.IsActive,
			).
			Values(
				1,
				"Example 1",
				1.1,
				"unique-1",
				time.Now().Unix(),
				true,
			).
			Exec(t.Context())

		if err != nil {
			t.Fatalf("failed to insert row: %v", err)
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			t.Fatalf("failed to get rows affected: %v", err)
		}

		if rowsAffected != 1 {
			t.Fatalf("expected 1 row affected, got %d", rowsAffected)
		}

		example, err := DB.GetExample(t, 1)
		if err != nil {
			t.Fatalf("failed to get inserted example: %v", err)
		}

		if example.ID != 1 ||
			example.Text != "Example 1" ||
			example.Value != 1.1 ||
			example.Unique != "unique-1" ||
			!example.IsActive {
			t.Fatalf("inserted example does not match expected values: %+v", example)
		}
	})

	t.Run("Insert multiple rows", func(t *testing.T) {
		DB := sqlkittesting.NewInMemoryTestDB(t)

		res, err := DB.
			Insert(sqlkittesting.Examples).
			Columns(
				sqlkittesting.Examples.ID,
				sqlkittesting.Examples.Text,
				sqlkittesting.Examples.Value,
				sqlkittesting.Examples.Unique,
				sqlkittesting.Examples.CreatedAt,
				sqlkittesting.Examples.IsActive,
			).
			Values(
				1,
				"Example 1",
				1.1,
				"unique-1",
				time.Now().Unix(),
				true,
			).
			Values(
				2,
				"Example 2",
				2.2,
				"unique-2",
				time.Now().Unix(),
				false,
			).
			Exec(t.Context())

		if err != nil {
			t.Fatalf("failed to insert multiple rows: %v", err)
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			t.Fatalf("failed to get rows affected: %v", err)
		}

		if rowsAffected != 2 {
			t.Fatalf("expected 2 rows affected, got %d", rowsAffected)
		}
	})

	t.Run("Insert duplicate primary key fails", func(t *testing.T) {
		DB := sqlkittesting.NewInMemoryTestDB(t, sqlkittesting.WithExampleData(
			&sqlkittesting.Example{
				ID:        1,
				Text:      "Example 1",
				Value:     1.1,
				Unique:    "unique-1",
				CreatedAt: time.Now(),
				IsActive:  true,
			},
		))

		res, err := DB.
			Insert(sqlkittesting.Examples).
			Columns(
				sqlkittesting.Examples.ID,
				sqlkittesting.Examples.Text,
				sqlkittesting.Examples.Value,
				sqlkittesting.Examples.Unique,
				sqlkittesting.Examples.CreatedAt,
				sqlkittesting.Examples.IsActive,
			).
			Values(
				1,
				"Example 1 Copy",
				1.1,
				"unique-1-copy",
				time.Now().Unix(),
				true,
			).
			Exec(t.Context())

		if err == nil {
			t.Fatalf("expected error when inserting duplicate primary key, got nil")
		}

		if !strings.Contains(err.Error(), "UNIQUE constraint failed") {
			t.Fatalf("expected UNIQUE constraint failed error, got: %v", err)
		}

		if res != nil {
			t.Fatalf("expected nil result when insert fails, got %v", res)
		}
	})
}
