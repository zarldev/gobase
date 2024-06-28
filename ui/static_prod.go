// build prod
//go:build prod
// +build prod

package ui

import (
	"embed"
	"log/slog"

	"net/http"
)

//go:embed static
var staticFS embed.FS

func StaticFileHandler() http.Handler {
	slog.Info("static file handler", slog.Any("mode", "prod"))
	return http.FileServer(http.FS(staticFS))
}
