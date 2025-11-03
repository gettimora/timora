package sqlkit

import (
	"context"
	"time"
)

type contextKeyOperationStart struct{}

func withOperationStart(ctx context.Context) context.Context {
	if _, ok := getOperationStart(ctx); ok {
		return ctx
	}
	return context.WithValue(ctx, contextKeyOperationStart{}, time.Now())
}

func getOperationStart(ctx context.Context) (time.Time, bool) {
	t, ok := ctx.Value(contextKeyOperationStart{}).(time.Time)
	return t, ok
}

func (db *SQLDatabase) logQuery(ctx context.Context, op string, query string, args []any, err error) {
	cfg := db.config.LogConfig
	if !cfg.Enabled {
		return
	}

	fields := []any{"query", query}

	if cfg.IncludeArgs {
		fields = append(fields, "args", args)
	}

	if cfg.LogDuration {
		if opStart, ok := getOperationStart(ctx); ok {
			fields = append(fields, "operation_duration", time.Since(opStart))
		}
	}

	if err != nil {
		fields = append(fields, "error", err)
	}

	db.config.Log.Info(op, fields...)
}
