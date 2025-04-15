package handlers

import (
	"context"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/VictorHRRios/catsnob/internal/database"
)

const layout = "templates/layout.html"

func (cfg *ApiConfig) HandlerIndex(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("templates", "index.html")

	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
	}

	artists, err := cfg.Queries.GetTop12Artists(context.Background())

	data := struct {
		Stylesheet *string
		Artists    []database.Artist
	}{
		Stylesheet: nil,
		Artists:    artists,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}

}
