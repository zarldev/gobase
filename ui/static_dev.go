//go:build !prod
// +build !prod

// build !prod
package ui

import (
	"log/slog"
	"net/http"
)

func StaticFileHandler() http.Handler {
	slog.Info("static file handler", slog.Any("mode", "dev"))
	return http.FileServer(http.Dir("ui/static"))
}
