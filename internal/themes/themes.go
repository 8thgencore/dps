package themes

import (
	"fmt"

	"github.com/muesli/termenv"
)

// Theme defines a color theme used for printing tables.
type Theme struct {
	ColorRed     termenv.Color
	ColorYellow  termenv.Color
	ColorGreen   termenv.Color
	ColorBlue    termenv.Color
	ColorGray    termenv.Color
	ColorMagenta termenv.Color
	ColorCyan    termenv.Color
}

func DefaultThemeName() string {
	if !termenv.HasDarkBackground() {
		return "light"
	}
	return "dark"
}

// var env = termenv.EnvColorProfile()

func LoadTheme(theme string, term termenv.Profile) (Theme, error) {
	themes := make(map[string]Theme)

	themes["dark"] = Theme{
		ColorRed:     term.Color("#E88388"),
		ColorYellow:  term.Color("#DBAB79"),
		ColorGreen:   term.Color("#A8CC8C"),
		ColorBlue:    term.Color("#71BEF2"),
		ColorGray:    term.Color("#B9BFCA"),
		ColorMagenta: term.Color("#D290E4"),
		ColorCyan:    term.Color("#66C2CD"),
	}

	themes["light"] = Theme{
		ColorRed:     term.Color("#D70000"),
		ColorYellow:  term.Color("#FFAF00"),
		ColorGreen:   term.Color("#005F00"),
		ColorBlue:    term.Color("#000087"),
		ColorGray:    term.Color("#303030"),
		ColorMagenta: term.Color("#AF00FF"),
		ColorCyan:    term.Color("#0087FF"),
	}

	themes["ansi"] = Theme{
		ColorRed:     term.Color("9"),
		ColorYellow:  term.Color("11"),
		ColorGreen:   term.Color("10"),
		ColorBlue:    term.Color("12"),
		ColorGray:    term.Color("7"),
		ColorMagenta: term.Color("13"),
		ColorCyan:    term.Color("8"),
	}

	if _, ok := themes[theme]; !ok {
		return Theme{}, fmt.Errorf("unknown theme: %s", theme)
	}

	return themes[theme], nil
}
