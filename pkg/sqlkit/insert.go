package sqlkit

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

var _ InsertQuery = (*SQLInsertQuery)(nil)

type (
	InsertQuery interface {
		Columns(cols ...Column) InsertQuery
		Values(vals ...any) InsertQuery
		BuildSQL() (string, []any)
		Exec(ctx context.Context) (sql.Result, error)
	}

	SQLInsertQuery struct {
		db      DB
		table   Table
		columns []Column
		rows    [][]any
	}
)

func (db *SQLDatabase) Insert(t Table) InsertQuery {
	return &SQLInsertQuery{db: db, table: t}
}

func (q *SQLInsertQuery) Columns(cols ...Column) InsertQuery {
	q.columns = cols
	return q
}

func (q *SQLInsertQuery) Values(vals ...any) InsertQuery {
	q.rows = append(q.rows, vals)
	return q
}

func (q *SQLInsertQuery) BuildSQL() (string, []any) {
	d := q.db.Dialect()

	colNames := make([]string, len(q.columns))
	for i, c := range q.columns {
		colNames[i] = d.QuoteIdent(c.Name())
	}

	args := make([]any, 0, len(q.rows)*len(q.columns))
	placeholders := make([]string, 0, len(q.rows))
	for _, row := range q.rows {
		start := len(args)
		rowPlaceholders := make([]string, len(row))
		for j := range row {
			rowPlaceholders[j] = d.Placeholder(start + j + 1)
		}
		placeholders = append(placeholders, fmt.Sprintf("(%s)", strings.Join(rowPlaceholders, ", ")))
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

func (q *SQLInsertQuery) Exec(ctx context.Context) (sql.Result, error) {
	ctx = withOperationStart(ctx)
	sqlStr, args := q.BuildSQL()
	return q.db.ExecContext(ctx, sqlStr, args...)
}
