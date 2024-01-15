package rules

import (
	"encoding/json"
	"net/http"

	chi "github.com/go-chi/chi/v5"
	"github.com/julioshinoda/go-rest-exercise/internal/model"
)

type Handler struct {
	r *chi.Mux
	s *Service
}

func NewHandler(r *chi.Mux, srv *Service) *Handler {
	return &Handler{r: r, s: srv}
}

func (h *Handler) Evaluate() {
	h.r.Post("/evaluate", func(w http.ResponseWriter, r *http.Request) {
		var request model.Request
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err := h.s.Evaluate(request)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Write([]byte("ok"))
		return
	})
}
