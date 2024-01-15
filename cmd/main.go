package main

import (
	//	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/julioshinoda/go-rest-exercise/internal/rules"
	"github.com/julioshinoda/go-rest-exercise/pkg/optii"
)

func main() {
	//	ctx := context.Background()

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	provider := optii.NewClient(os.Getenv("OPTII_USER"), os.Getenv("OPTII_SECRET"))
	rulesService := rules.NewService(provider)
	rh := rules.NewHandler(r, rulesService)
	rh.Evaluate()
	slog.Info(fmt.Sprintf("server running on port: %s", os.Getenv("SERVER_PORT")))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")), r); err != nil {
		slog.Error(err.Error())
	}
}
