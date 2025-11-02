package sqlite

import (
	"fmt"

	"github.com/gettimora/timora/pkg/sqlkit"
)

var _ sqlkit.Dialect = (*Dialect)(nil)

type Dialect struct{}

func NewDialect() *Dialect {
	return &Dialect{}
}

func (d *Dialect) Placeholder(n int) string {
	return "?"
}

func (d *Dialect) QuoteIdent(ident string) string {
	return fmt.Sprintf(`"%s"`, ident)
}
