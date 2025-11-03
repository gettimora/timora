package sqlkit

import "time"

type (
	Table interface {
		Name() string
	}

	Column interface {
		Name() string
	}

	ColumnOf[T any] struct {
		ColName string
	}
)

var _ Column = ColumnOf[any]{}

func (c ColumnOf[T]) Name() string { return c.ColName }

func NewStringColumn(name string, table Table) ColumnOf[string] {
	return ColumnOf[string]{ColName: name}
}

func NewIntColumn(name string, table Table) ColumnOf[int] {
	return ColumnOf[int]{ColName: name}
}

func NewTimeColumn(name string, table Table) ColumnOf[time.Time] {
	return ColumnOf[time.Time]{ColName: name}
}

func NewFloatColumn(name string, table Table) ColumnOf[float64] {
	return ColumnOf[float64]{ColName: name}
}

func NewBoolColumn(name string, table Table) ColumnOf[bool] {
	return ColumnOf[bool]{ColName: name}
}
