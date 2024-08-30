package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

func NewHandler(db map[string]string) http.Handler {

	r := chi.NewMux()

	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Get("/api/users", HandleGetUsers(db))
	r.Get("/api/users/:id", HandleGetUsersWithID(db))
	r.Delete("/api/users/:id")
	r.Post("/api/users")
	r.Put("/api/users/:id")

	return r
}

type id uuid.UUID

type user struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	biography string
}

type application struct {
	data map[id]user
}

func HandleGetUsers(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		users, err := json.Marshal(db)
		if err != nil {
			http.Error(w, "failed to encode users", http.StatusInternalServerError)
			return
		}
		w.Write(users)
	}
}

func HandleGetUsersWithID(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		user, ok := db[id]
		if !ok {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(user))
	}
}
