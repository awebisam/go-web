package render

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

func Template(w http.ResponseWriter, tmpl string) {

	templateMap, err := GetTemplateMap(w)
	if err != nil {
		log.Println("Error getting template map")
	}

	templateInstance := templateMap[tmpl]

	executionError := templateInstance.Execute(w, nil)
	if executionError != nil {
		log.Print(err)
		return
	}
}

func GetTemplateMap(w http.ResponseWriter) (map[string]*template.Template, error) {
	/* A Function To Get  */
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
