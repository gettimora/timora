package sqlkit_test

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/gettimora/timora/pkg/sqlkit"
	sqlkittesting "github.com/gettimora/timora/pkg/sqlkit/testing"
)

func ExampleDB_Insert() {
	ctx := context.Background()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db := sqlkittesting.NewExampleTestDB(sqlkit.WithLogger(logger),
		sqlkit.WithLogConfig(sqlkit.LogConfig{
			Enabled:     true,
			IncludeArgs: true,
			LogDuration: false,
		}),
		sqlkit.WithLogger(
			slog.New(slog.NewTextHandler(
				os.Stdout,
				&slog.HandlerOptions{

					AddSource: false,
					ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
						if a.Key == slog.TimeKey {
							a.Value = slog.StringValue("2025-11-02T00:00:00Z")
						}
						return a
					},
				})),
		),
	)

	res, err := db.Insert(sqlkittesting.Examples).
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
			time.Date(2023, 10, 11, 0, 0, 0, 0, time.UTC).Unix(),
			true,
		).
		Exec(ctx)

	if err != nil {
		panic("failed to insert row: " + err.Error())
	}

	fmt.Println(res.LastInsertId())
	fmt.Println(res.RowsAffected())

	// Output:
	// time=2025-11-02T00:00:00Z level=INFO msg="Running DB Execution" query="CREATE TABLE \"examples\" (\n\t\"id\" INTEGER PRIMARY KEY,\n\t\"text\" TEXT NOT NULL,\n\t\"value\" REAL,\n\t\"unique\" TEXT NOT NULL UNIQUE,\n\t\"created_at\" INTEGER NOT NULL,\n\t\"is_active\" INTEGER NOT NULL\n);" args=[]
	// time=2025-11-02T00:00:00Z level=INFO msg="Running DB Execution" query="INSERT INTO \"examples\" (\"id\", \"text\", \"value\", \"unique\", \"created_at\", \"is_active\") VALUES (?, ?, ?, ?, ?, ?)" args="[1 Example 1 1.1 unique-1 1696982400 true]"
	// 1 <nil>
	// 1 <nil>
}
