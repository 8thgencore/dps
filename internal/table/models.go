package table

import "github.com/jedib0t/go-pretty/table"

// TableOptions contains all options for the table.
type TableOptions struct {
	Columns []int
	SortBy  int
	Style   table.Style
}

// Column defines a column.
type Column struct {
	ID        string
	Name      string
	SortIndex int
	Width     int
}

type RowData struct {
	Data []string
}
