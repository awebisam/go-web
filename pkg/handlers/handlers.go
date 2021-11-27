package handlers

import (
	"github.com/awebisam/go-web/pkg/config"
	"net/http"

	"github.com/awebisam/go-web/pkg/render"
)

// Repo is a repository used by handlers
var Repo *Repository

// Repository TODO: Add database handlers and other things in repository
// Repository struct to store system wide used stuff
type Repository struct {
	App *config.AppConfig
}

// NewRepo returns a new repository with app config
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets a repository for handlers
func NewHandlers(repo *Repository) {
	Repo = repo
}

// About Handler handles about page
func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.TemplateRenderer(w, "about.page.tmpl")
}

// Home Handler handles home page
func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.TemplateRenderer(w, "home.page.tmpl")
}
