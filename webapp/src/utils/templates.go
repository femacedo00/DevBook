package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// LoadTemplates parses all HTML templates and stores them in the templates variable
func LoadTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
	templates = template.Must(templates.ParseGlob("views/templates/*.html"))
}

// ExecuteHtmlTemplate renders an html page to the HTTP response
func ExecuteHtmlTemplate(w http.ResponseWriter, template string, data interface{}) {
	templates.ExecuteTemplate(w, template, data)
}
