package templates

import (
	"html/template"
	"path/filepath"
	"sync"
)

var (
	templates     *template.Template
	loadTemplates sync.Once
)

func LoadTemplates() (*template.Template, error) {
	var err error
	loadTemplates.Do(func() {
		pattern := filepath.Join("internal", "templates", "views", "*.tmpl")
		templates, err = template.ParseGlob(pattern)
	})
	return templates, err
}
