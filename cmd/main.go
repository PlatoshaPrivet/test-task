package main

import (
	"log/slog"
	"net/http"
	"note_API/internal/auth"
	"note_API/internal/config"
	"note_API/internal/handlers"
	"note_API/pkg/logger"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	cfg := config.MustLoad()
	log := logger.New()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/notes", func(r chi.Router) {
		r.Use(auth.AuthMiddleware)
		r.Post("/", handlers.CreateNote)
		r.Get("/", handlers.GetNotes)
	})

	log.Info("starting server", slog.String("address", cfg.Address))

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      r,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}

	log.Error("server stopped")
}
