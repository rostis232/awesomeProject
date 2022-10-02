package main

import (
	"fmt"
	"github.com/rostis232/awesomeProject/pkg/config"
	"github.com/rostis232/awesomeProject/pkg/handlers"
	"github.com/rostis232/awesomeProject/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the main application.
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Starting application on port %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
