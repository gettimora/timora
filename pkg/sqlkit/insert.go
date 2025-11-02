package sqlkit

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type InsertQuery struct {
	db      *DB
	table   Table
	columns []Column
	rows    [][]any
}

func (db *DB) Insert(t Table) *InsertQuery {
	return &InsertQuery{db: db, table: t}
}

func (q *InsertQuery) Columns(cols ...Column) *InsertQuery {
	q.columns = cols
	return q
}

func (q *InsertQuery) Values(vals ...any) *InsertQuery {
	q.rows = append(q.rows, vals)
	return q
}

func (q *InsertQuery) BuildSQL() (string, []any) {
	d := q.db.Dialect()

	colNames := make([]string, len(q.columns))
	for i, c := range q.columns {
		colNames[i] = d.QuoteIdent(c.Name())
	}

	args := make([]any, 0, len(q.rows)*len(q.columns))
	placeholders := make([]string, 0, len(q.rows))
	for _, row := range q.rows {
		rowPlaceholders := make([]string, len(row))
		for j := range row {
			rowPlaceholders[j] = d.Placeholder(len(args) + j + 1)
		}
		args = append(args, row...)
		placeholders = append(placeholders, fmt.Sprintf("(%s)", strings.Join(rowPlaceholders, ", ")))
		args = append(args, row...)
		args = args[:len(args)-len(row)]
		args = append(args, row...)
	}

	sql := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES %s",
		d.QuoteIdent(q.table.Name()),
		strings.Join(colNames, ", "),
		strings.Join(placeholders, ", "),
	)

	return sql, args
}

func (q *InsertQuery) Exec(ctx context.Context) (sql.Result, error) {
	sqlStr, args := q.BuildSQL()
	return q.db.ExecContext(ctx, sqlStr, args...)
}
