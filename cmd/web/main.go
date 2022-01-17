package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anjotadena/projectX/pkg/config"
	"github.com/anjotadena/projectX/pkg/handlers"
	"github.com/anjotadena/projectX/pkg/render"
)

const portNumber = ":8000"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	// Create repository variable
	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
