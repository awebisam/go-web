package config

import "html/template"

// AppConfig Holds the application config items
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
