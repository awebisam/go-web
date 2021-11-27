package handlers

import (
	"net/http"

	"github.com/awebisam/go-web/pkg/render"
)

func About(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "about.page.tmpl")
}

func Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "home.page.tmpl")
}
