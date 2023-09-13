package config

import (
	"context"
	"flag"
	"fmt"
	"os"

	"dps/internal/themes"

	"github.com/muesli/termenv"
)

var (
	// Version contains the application version number. It's set via ldflags when building.
	Version = "0.1.0"

	all = flag.Bool("all", false, "include pseudo, duplicate, inaccessible file systems")
	// containers = flag.Bool("containers", false, "include containers")

	width    = flag.Uint("width", 160, "max output width")
	themeOpt = flag.String("theme", themes.DefaultThemeName(), "color themes: dark, light, ansi")
	styleOpt = flag.String("style", themes.DefaultStyleName(), "style: unicode, ascii")
)

type Config struct {
	Width uint
	theme themes.Theme
	term  termenv.Profile
}

// Ключ для доступа к конфигурации в контексте
type configContextKey struct{}

// Функция для установки конфигурации в контекст
func SetConfig(ctx context.Context, cfg *Config) context.Context {
	return context.WithValue(ctx, configContextKey{}, cfg)
}

// Функция для получения конфигурации из контекста
func GetConfig(ctx context.Context) *Config {
	if cfg, ok := ctx.Value(configContextKey{}).(*Config); ok {
		return cfg
	}
	return nil
}

func Init(ctx context.Context) *Config {
	var err error

	flag.Parse()

	// validate theme
	theme := themes.Theme{}
	term := termenv.EnvColorProfile()

	theme, err = themes.LoadTheme(*themeOpt, term)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if term == termenv.ANSI {
		// enforce ANSI theme for limited color support
		theme, err = themes.LoadTheme("ansi", term)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	return &Config{
		Width: *width,
		theme: theme,
		term:  term,
	}
}
