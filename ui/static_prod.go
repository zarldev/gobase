// build prod
//go:build prod
// +build prod

package ui

import (
	"embed"

	"net/http"
)

//go:embed static
var staticFS embed.FS

func StaticFileHandler() http.Handler {
	return http.FileServer(http.FS(staticFS))
}
