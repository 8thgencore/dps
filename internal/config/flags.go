package config

import (
	"context"
	"fmt"
	"os"

	"dps/internal/themes"

	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
)

var (
	Version = "0.1.0"

	all      bool
	width    uint
	themeOpt string
	styleOpt string
)

var rootCmd = &cobra.Command{
	Use:   "dps",
	Short: "Docker Presentation System",
	Run: func(cmd *cobra.Command, args []string) {
		Init(cmd.Context())
	},
}

func init() {
	// Добавьте команды с псевдонимами
	rootCmd.AddCommand(containerCmd)
	rootCmd.AddCommand(imageCmd)
	rootCmd.AddCommand(networkCmd)
	rootCmd.AddCommand(volumeCmd)

	rootCmd.PersistentFlags().BoolVarP(&all, "all", "a", false, "include pseudo, duplicate, inaccessible file systems")
	rootCmd.PersistentFlags().UintVarP(&width, "width", "w", 160, "max output width")
	rootCmd.PersistentFlags().StringVar(&themeOpt, "theme", themes.DefaultThemeName(), "color themes: dark, light, ansi")
	rootCmd.PersistentFlags().StringVar(&styleOpt, "style", themes.DefaultStyleName(), "style: unicode, ascii")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// TODO:
var containerCmd = &cobra.Command{
	Use:     "container",
	Aliases: []string{"container"},
	Short:   "List containers",
}

var imageCmd = &cobra.Command{
	Use:     "image",
	Aliases: []string{"image"},
	Short:   "List images",
}

var networkCmd = &cobra.Command{
	Use:     "network",
	Aliases: []string{"network"},
	Short:   "List networks",
}

var volumeCmd = &cobra.Command{
	Use:     "volume",
	Aliases: []string{"volume"},
	Short:   "List volumes",
}

func Init(ctx context.Context) *Config {
	var err error

	// validate theme
	term := termenv.EnvColorProfile()

	theme, err := themes.LoadTheme(themeOpt, term)
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

	// validate style
	style, err := themes.ParseStyle(styleOpt)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return &Config{
		Width:    width,
		Theme:    theme,
		StyleOpt: style,
		term:     term,
	}
}
