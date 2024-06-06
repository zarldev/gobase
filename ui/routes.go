package ui

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/zarldev/go-base/ui/pages/home"
	"github.com/zarldev/go-base/ui/pages/landing"
	"github.com/zarldev/go-base/ui/pages/profile"
	"github.com/zarldev/go-base/ui/pages/settings"
)

func routes(r *chi.Mux) {
	r.Get("/", LandingHandler())
	r.Get("/user", HomePageHandler())
	r.Get("/home", HomeHandler())
	r.Get("/profile", ProfileHandler())
	r.Get("/settings", SettingsHandler())
	r.Post("/settings", SettingsSaveHandler())
	r.Handle("/static/*", http.StripPrefix("/static/", StaticFileHandler()))
}

func LandingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Render(r.Context(), landing.Page(), w)
	}
}

func HomePageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Render(r.Context(), home.Page(), w)
	}
}

func HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("HomeHandler")
		Render(r.Context(), home.Content(), w)
	}
}

func ProfileHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("ProfileHandler")
		Render(r.Context(), profile.Page(), w)
	}
}

func SettingsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("SettingsHandler")
		Render(r.Context(), settings.Page(
			"Bruno", "bruno@bruno.com",
			false), w)
	}
}

func SettingsSaveHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("SettingsSaveHandler")
		newName := r.FormValue("name")
		newEmail := r.FormValue("email")
		Render(r.Context(), settings.Page(
			newName, newEmail,
			true), w)
	}
}

func AuthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("AuthHandler")
		http.Redirect(w, r, "/user", http.StatusFound)
		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    "123456789",
			Expires:  time.Now().Add(time.Hour * 24 * 365),
			HttpOnly: true,
		})
		http.SetCookie(w, &http.Cookie{
			Name:     "user",
			Value:    "bruno",
			Expires:  time.Now().Add(time.Hour * 24 * 365),
			HttpOnly: true,
		})
	}
}
