package ui

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)

type Config struct {
	Port string
}

const (
	timeout5sec  = 5 * time.Second
	timeout10sec = 10 * time.Second
	timeout15sec = 15 * time.Second
)

func StartUI(ctx context.Context, config Config) {
	r := chi.NewRouter()
	routes(r)
	server := &http.Server{
		Addr:         config.Port,
		Handler:      r,
		ReadTimeout:  timeout5sec,
		WriteTimeout: timeout10sec,
		IdleTimeout:  timeout15sec,
	}
	go func() {
		slog.InfoContext(ctx, "starting server", slog.String("port", config.Port))
		if err := server.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				slog.InfoContext(ctx, "server closed")
				return
			}
			slog.ErrorContext(ctx, "failed to start server", slog.Any("error", err))
		}
	}()
	go func() {
		slog.InfoContext(ctx, "listening for shutdown signal")
		<-ctx.Done()
		slog.InfoContext(ctx, "shutting down ui server")
		if err := server.Shutdown(ctx); err != nil {
			slog.ErrorContext(ctx, "failed to shutdown server", slog.Any("error", err))
		}
	}()
}

func Render(ctx context.Context, comp templ.Component, w io.Writer) {
	_ = comp.Render(ctx, w)
}
