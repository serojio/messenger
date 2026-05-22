package http

import (
	"encoding/json"
	"net/http"

	"users/internal/transport/http/handler"

	"github.com/go-chi/chi"
)

func NewRouter(userHandler *handler.UserHandler) http.Handler {
	r := chi.NewRouter()
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]string{
			"status": "ok",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.Create)
		r.Get("/", userHandler.List)

		r.Get("/{id}", userHandler.Get)
		r.Patch("/{id}", userHandler.Update)
		r.Delete("/{id}", userHandler.Delete)

	})

	return r
}
