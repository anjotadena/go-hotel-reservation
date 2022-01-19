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
	// r.Use(WriteToConsole)
	r.Use(NoSurf)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(SessionLoad)

	// routes
	r.Get("/", handlers.Repo.Home)
	r.Get("/about", handlers.Repo.About)

	fileServer := http.FileServer(http.Dir("./static/"))

	r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return r
}
