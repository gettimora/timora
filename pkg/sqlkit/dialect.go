package sqlkit

// Dialect defines SQL dialect-specific behaviors.
type Dialect interface {
	// Placeholder returns the placeholder string for the n-th parameter.
	// For example, in Postgres, Placeholder(1) returns "$1".
	Placeholder(n int) string
	// QuoteIdent returns the quoted identifier for the given name.
	// For example, in Postgres, QuoteIdent("user") returns `"user"`.
	QuoteIdent(name string) string
}
