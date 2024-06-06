package ui

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)

type Config struct {
	Port string
}

func StartUI(ctx context.Context, config Config) {
	r := chi.NewRouter()
	routes(r)
	server := &http.Server{
		Addr:         config.Port,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}

func Render(ctx context.Context, comp templ.Component, w io.Writer) {
	_ = comp.Render(ctx, w)
}
