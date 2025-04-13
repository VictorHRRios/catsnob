package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

const layout = "templates/layout.html"

func (cfg *ApiConfig) HandlerIndex(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("templates", "index.html")

	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}

}
