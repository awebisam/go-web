package config

import "html/template"

// AppConfig Holds the application config items
type AppConfig struct {
	Debug         bool // Debug mode
	TemplateCache map[string]*template.Template
}
