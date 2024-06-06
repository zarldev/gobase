//go:build !prod
// +build !prod

// build !prod
package ui

import (
	"fmt"
	"net/http"
)

func StaticFileHandler() http.Handler {
	fmt.Println("Running static file handler in dev mode.")
	return http.FileServer(http.Dir("ui/static"))
}
