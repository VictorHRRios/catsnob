package handlers

import (
	"context"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
)

func (cfg *ApiConfig) HandlerArtistProfile(w http.ResponseWriter, r *http.Request) {
	artistName := r.PathValue("artist")
	artistName = strings.ReplaceAll(artistName, "_", " ")
	artist, err := cfg.Queries.GetArtist(context.Background(), artistName)
	if err != nil {
		http.Error(w, "Error fetching artist", http.StatusInternalServerError)
		return
	}

	tmplPath := filepath.Join("templates", "artistProfile.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
	}
	data := struct {
		Stylesheet *string
		Name       string
		Img        string
		Biography  string
		FormedAt   string
		Genre      string
	}{
		Stylesheet: nil,
		Name:       artist.Name,
		Img:        artist.ImgUrl,
		Biography:  artist.Biography.String,
		FormedAt:   artist.FormedAt,
		Genre:      artist.Genre,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
