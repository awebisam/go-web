package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
)

// AppConfig TODO: Add more config options
// AppConfig Holds the application config items
type AppConfig struct {
	Debug         bool
	TemplateCache map[string]*template.Template
	Session       *scs.SessionManager
}
