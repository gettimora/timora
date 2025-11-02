package sqlkit

import "time"

type (
	Table interface {
		Name() string
	}

	Column interface {
		Name() string
		Table() Table
	}

	ColumnOf[T any] struct {
		ColName string
		Tab     Table
	}
)

var _ Column = ColumnOf[any]{}

func (c ColumnOf[T]) Name() string { return c.ColName }
func (c ColumnOf[T]) Table() Table { return c.Tab }

func NewStringColumn(name string, table Table) ColumnOf[string] {
	return ColumnOf[string]{ColName: name, Tab: table}
}

func NewIntColumn(name string, table Table) ColumnOf[int] {
	return ColumnOf[int]{ColName: name, Tab: table}
}

func NewTimeColumn(name string, table Table) ColumnOf[time.Time] {
	return ColumnOf[time.Time]{ColName: name, Tab: table}
}

func NewFloatColumn(name string, table Table) ColumnOf[float64] {
	return ColumnOf[float64]{ColName: name, Tab: table}
}

func NewBoolColumn(name string, table Table) ColumnOf[bool] {
	return ColumnOf[bool]{ColName: name, Tab: table}
}
