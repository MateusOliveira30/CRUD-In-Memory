package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

func NewHandler(db map[string]User) http.Handler {

	r := chi.NewMux()

	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Get("/api/users", HandleGetUsers(db))
	r.Get("/api/users/:id", HandleGetUsersWithID(db))
	r.Delete("/api/users/:id", HandleDelete(db))
	r.Post("/api/users", HandlePost(db))
	r.Put("/api/users/:id")

	return r
}

type id uuid.UUID

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Biography string `json:"biography"`
}

type application struct {
	data map[id]User
}

func HandleGetUsers(db map[string]User) http.HandlerFunc {
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

func HandleGetUsersWithID(db map[string]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		user, ok := db[id]
		if !ok {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

func HandlePost(db map[string]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body User
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "invalid input", http.StatusBadRequest)
			return
		}
		id := uuid.NewString()
		db[id] = body
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"id": id})
	}
}

func HandleDelete(db map[string]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		_, ok := db[id]
		if !ok {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}
		delete(db, id)
		w.WriteHeader(http.StatusNoContent)
	}
}
