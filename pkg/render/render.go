package render

import (
	"bytes"
	"github.com/awebisam/go-web/pkg/config"
	"github.com/awebisam/go-web/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

// TemplateRenderer TODO: Improve Error Handling
func TemplateRenderer(w http.ResponseWriter, tmpl string, templateData *models.TemplateData) {
	var templateCache map[string]*template.Template
	{
	}

	if app.Debug {
		templateCache, _ = CreateTemplateCache()

	} else {
		templateCache = app.TemplateCache
	}

	templateInstance, ok := templateCache[tmpl]

	if !ok {
		log.Fatalln("Trying to render a template that doesn't exist")
	}

	buf := new(bytes.Buffer)

	_ = templateInstance.Execute(buf, templateData)

	_, execError := buf.WriteTo(w)

	//ExecError := templateInstance.Execute(w, nil)
	if execError != nil {
		log.Print(execError)
		return
	}
}

// CreateTemplateCache creat
func CreateTemplateCache() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		log.Print(err)
		return templateCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			log.Print(err)
			return templateCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			log.Print(err)
			return templateCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				log.Print(err)
				return templateCache, err
			}
		}
		templateCache[name] = ts
	}
	return templateCache, err
}
