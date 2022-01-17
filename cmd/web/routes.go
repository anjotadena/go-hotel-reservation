package main

import (
	"net/http"

	"github.com/anjotadena/projectX/pkg/config"
	"github.com/anjotadena/projectX/pkg/handlers"
	"github.com/bmizerany/pat"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))

	return mux
}
