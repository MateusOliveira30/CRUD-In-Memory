package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(db map[string]string) http.Handler {

	r := chi.NewMux()

	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Get("/api/users")
	r.Get("/api/users/:id")
	r.Delete("/api/users/:id")
	r.Post("/api/users")
	r.Put("/api/users/:id")

	return r
}
