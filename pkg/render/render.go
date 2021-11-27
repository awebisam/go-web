package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

// TemplateRenderer TODO: Improve Error Handling
func TemplateRenderer(w http.ResponseWriter, tmpl string) {

	templateMap, err := GetTemplateMap()
	if err != nil {
		log.Fatalln("Error getting template map")
	}

	templateInstance, ok := templateMap[tmpl]

	if !ok {
		log.Fatalln("Trying to render a template that doesn't exist")
	}

	buf := new(bytes.Buffer)

	_ = templateInstance.Execute(buf, nil)

	_, execError := buf.WriteTo(w)

	//ExecError := templateInstance.Execute(w, nil)
	if execError != nil {
		log.Print(err)
		return
	}
}

func GetTemplateMap() (map[string]*template.Template, error) {
	/* A Function To Get Template Map With Template Name And type: template.Template instance  */
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
