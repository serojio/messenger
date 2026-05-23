package http

import (
	"encoding/json"
	"net/http"

	"posts/internal/transport/http/handler"

	"github.com/go-chi/chi"
)

func NewRouter(postHandler *handler.PostHandler) http.Handler {
	r := chi.NewRouter()
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]string{
			"status": "ok",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	r.Route("/posts", func(r chi.Router) {
		r.Post("/", postHandler.Create)
		r.Get("/", postHandler.List)

		r.Get("/{id}", postHandler.Get)
		r.Patch("/{id}", postHandler.Update)
		r.Delete("/{id}", postHandler.Delete)

	})

	return r
}
