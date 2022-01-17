package main

import (
	"net/http"

	"github.com/anjotadena/projectX/pkg/config"
	"github.com/anjotadena/projectX/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// routes
	r.Get("/", handlers.Repo.Home)

	return r
}
