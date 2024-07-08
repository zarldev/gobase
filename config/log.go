package config

import (
	"io"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

func loadLogger() *slog.Logger {
	var (
		w io.Writer
		h slog.Handler
	)
	switch ENVIRONMENT.ENV {
	case Dev:
		w = os.Stdout
		h = tint.NewHandler(w, &tint.Options{
			Level:      ENVIRONMENT.LOG_LEVEL.Level(),
			TimeFormat: time.Kitchen,
		})
	case Staging:
		w = os.Stderr
		h = tint.NewHandler(w, &tint.Options{
			Level:      ENVIRONMENT.LOG_LEVEL.Level(),
			TimeFormat: time.RFC3339Nano,
		})
	case Prod:
		w = os.Stderr
		h = slog.NewJSONHandler(w, &slog.HandlerOptions{
			Level: ENVIRONMENT.LOG_LEVEL.Level(),
		})
	default:
		w = os.Stdout
		h = tint.NewHandler(w, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		})
	}
	// create a new logger
	logger := slog.New(h)
	// set global logger with custom options
	slog.SetDefault(logger)
	return logger
}
