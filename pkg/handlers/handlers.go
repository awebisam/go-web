package handlers

import (
	"net/http"

	"github.com/awebisam/go-web/pkg/render"
)

func About(w http.ResponseWriter, r *http.Request) {
	render.TemplateRenderer(w, "about.page.tmpl")
}

func Home(w http.ResponseWriter, r *http.Request) {
	render.TemplateRenderer(w, "home.page.tmpl")
}
