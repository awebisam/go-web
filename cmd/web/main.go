package main

import (
	"fmt"
	"github.com/awebisam/go-web/pkg/config"
	"github.com/awebisam/go-web/pkg/render"
	"log"
	"net/http"

	"github.com/awebisam/go-web/pkg/handlers"
)

const portNumber = ":8080"

func main() {

	var app config.AppConfig

	tc, createCacheError := render.CreateTemplateCache()
	if createCacheError != nil {
		log.Fatal("Unable to load templates into cache")
	}

	app.TemplateCache = tc
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about/", handlers.About)

	fmt.Printf("Starting server in port %s", portNumber)

	servingError := http.ListenAndServe(
		portNumber,
		nil,
	)
	if servingError != nil {
		return
	}
}
