package sqlkit_test

import (
	"database/sql"
	"strings"
	"testing"
	"time"

	test "github.com/gettimora/timora/pkg/sqlkit/testing"
)

func Test_Insert(t *testing.T) {
	tests := []struct {
		Name   string
		Setup  []test.Option
		Assert func(*testing.T, sql.Result, error)
	}{
		{
			Name:  "Insert without values",
			Setup: nil,
			Assert: func(t *testing.T, res sql.Result, err error) {
				if err != nil {
					t.Fatalf("expected no error, got %v", err)
				}
				rowsAffected, err := res.RowsAffected()
				if err != nil {
					t.Fatalf("failed to get rows affected: %v", err)
				}
				if rowsAffected != 1 {
					t.Fatalf("expected 1 row affected, got %d", rowsAffected)
				}
			},
		},
		{
			Name: "Insert with copy of unique value",
			Setup: []test.Option{
				test.WithExample(test.Example{
					ID:        1,
					Text:      "example1",
					Unique:    "unique1",
					Value:     42,
					CreatedAt: time.Now(),
				}),
			},
			Assert: func(t *testing.T, res sql.Result, err error) {
				if err == nil {
					t.Fatalf("expected error, got none")
				}
				if !strings.Contains(err.Error(), "UNIQUE constraint failed") {
					t.Fatalf("expected UNIQUE constraint failed error, got %v", err)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			DB := test.NewInMemoryTest(t, tt.Setup...)

			res, err := DB.
				Insert(test.Examples).
				Columns(
					test.Examples.ID,
					test.Examples.Text,
					test.Examples.Unique,
					test.Examples.Value,
					test.Examples.CreatedAt,
					test.Examples.IsActive,
				).
				Values(
					1,
					"example1",
					"unique1",
					42.4,
					time.Now(),
					true,
				).
				Exec(t.Context())

			tt.Assert(t, res, err)
		})
	}
}
