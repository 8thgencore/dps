package table

import (
	"context"
	"dps/internal/config"
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	"github.com/muesli/termenv"
)

func RenderTable(ctx context.Context, columns []Column, rows []RowData) {
	// Get the configuration from the context
	config := config.GetConfig(ctx)

	// Create a new table writer
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetAllowedRowLength(int(config.Width))

	// Configure headers
	headerRow := make([]interface{}, len(columns))
	for i, col := range columns {
		headerRow[i] = col.Name
	}
	t.AppendHeader(headerRow)
	// Configure rows
	for _, row := range rows {
		dataRow := make([]interface{}, len(row.Data))
		for i, data := range row.Data {
			dataRow[i] = termenv.String(text.WrapSoft(data, columns[i].Width)).Foreground(config.Theme.ColorBlue)
		}
		t.AppendRow(dataRow)
	}

	// Create a new table writer
	t.SetStyle(config.StyleOpt)
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1},
	})

	// Create a new table writer
	t.Render()
}
