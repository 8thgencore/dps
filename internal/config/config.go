package config

import (
	"context"

	"dps/internal/themes"

	"github.com/jedib0t/go-pretty/table"
	"github.com/muesli/termenv"
)

type Config struct {
	Width    uint
	Theme    themes.Theme
	StyleOpt table.Style
	term     termenv.Profile
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
