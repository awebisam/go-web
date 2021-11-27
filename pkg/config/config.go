package config

import "html/template"

// AppConfig TODO: Add more config options
// AppConfig Holds the application config items
type AppConfig struct {
	Debug bool
	/* When Debug is true:
	1. Templates are cached on every request that means that the application
		will be slower, but you don't need to restart the server everytime you
		make changes on template.
	2. TODO: Add more usage on Debug mode
	*/

	TemplateCache map[string]*template.Template
}
