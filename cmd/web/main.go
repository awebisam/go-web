package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/awebisam/go-web/pkg/config"
	"github.com/awebisam/go-web/pkg/handlers"
	"github.com/awebisam/go-web/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var (
	app     config.AppConfig
	session *scs.SessionManager
)

// main TODO: refactor and tidy up.
// main is the entry point for the application.
func main() {
	// change this to false in production.
	app.Debug = true

	loadSession()
	app.Session = session

	tc, createCacheError := render.CreateTemplateCache()
	if createCacheError != nil {
		log.Fatal("Unable to load templates into cache")
	}

	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Printf("Starting server on port %s", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

func loadSession() {
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.Debug
}
