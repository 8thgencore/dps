package themes

import (
	"fmt"

	"github.com/jedib0t/go-pretty/table"
	"github.com/mattn/go-runewidth"
)

func DefaultStyleName() string {
	/*
		Due to a bug in github.com/mattn/go-runewidth v0.0.9, the width of unicode rune(such as '╭') could not be correctly
		calculated.	Degrade to ascii to prevent broken table structure. Remove this once the bug is fixed.
	*/
	if runewidth.RuneWidth('╭') > 1 {
		return "ascii"
	}

	return "unicode"
}

// parseStyle converts user-provided style option into a table.Style.
func ParseStyle(styleOpt string) (table.Style, error) {
	switch styleOpt {
	case "unicode":
		return table.StyleRounded, nil
	case "ascii":
		return table.StyleDefault, nil
	default:
		return table.Style{}, fmt.Errorf("unknown style option: %s", styleOpt)
	}
}
