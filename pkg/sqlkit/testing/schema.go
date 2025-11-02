package testing

import (
	"time"

	"github.com/gettimora/timora/pkg/sqlkit"
)

const (
	CreateSchema = `CREATE TABLE "examples" (
	"id" INTEGER PRIMARY KEY,
	"text" TEXT NOT NULL,
	"value" REAL,
	"unique" TEXT NOT NULL UNIQUE,
	"created_at" INTEGER NOT NULL,
	"is_active" INTEGER NOT NULL
);`
)

type (
	ExampleTable struct {
		ID        sqlkit.ColumnOf[int]
		Text      sqlkit.ColumnOf[string]
		Value     sqlkit.ColumnOf[float64]
		Unique    sqlkit.ColumnOf[string]
		CreatedAt sqlkit.ColumnOf[time.Time]
		IsActive  sqlkit.ColumnOf[bool]
	}

	Example struct {
		ID        int
		Text      string
		Value     float64
		Unique    string
		CreatedAt time.Time
		IsActive  bool
	}
)

var (
	Examples = NewExampleTable()
)

func NewExampleTable() *ExampleTable {
	t := &ExampleTable{}
	t.ID = sqlkit.NewIntColumn("id", t)
	t.Text = sqlkit.NewStringColumn("text", t)
	t.Value = sqlkit.NewFloatColumn("value", t)
	t.Unique = sqlkit.NewStringColumn("unique", t)
	t.CreatedAt = sqlkit.NewTimeColumn("created_at", t)
	t.IsActive = sqlkit.NewBoolColumn("is_active", t)
	return t
}

func (e *ExampleTable) Name() string {
	return "examples"
}
