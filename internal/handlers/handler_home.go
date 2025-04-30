package handlers

import (
	"context"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/VictorHRRios/catsnob/internal/database"
)

const layout = "templates/layout.html"

func (cfg *ApiConfig) HandlerIndex(w http.ResponseWriter, r *http.Request, u *database.User) {
	tmplPath := filepath.Join("templates", "home", "artists.html")

	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	artists, err := cfg.Queries.GetTop12Artists(context.Background())
	if err != nil {
		http.Error(w, "Error fetching artists", http.StatusInternalServerError)
		return
	}

	data := struct {
		Stylesheet *string
		Artists    []database.Artist
		User       *database.User
	}{
		Stylesheet: nil,
		Artists:    artists,
		User:       u,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}

}

func (cfg *ApiConfig) HandlerAlbums(w http.ResponseWriter, r *http.Request, u *database.User) {
	tmplPath := filepath.Join("templates", "home", "albums.html")

	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	albums, err := cfg.Queries.GetTop12Albums(context.Background())

	data := struct {
		Stylesheet *string
		Albums     []database.GetTop12AlbumsRow
		User       *database.User
	}{
		Stylesheet: nil,
		Albums:     albums,
		User:       u,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}

}

func (cfg *ApiConfig) HandlerTracks(w http.ResponseWriter, r *http.Request, u *database.User) {
	tmplPath := filepath.Join("templates", "home", "tracks.html")

	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	tracks, err := cfg.Queries.GetTop12Tracks(context.Background())
	data := struct {
		Stylesheet *string
		Tracks     []database.GetTop12TracksRow
		User       *database.User
	}{
		Stylesheet: nil,
		Tracks:     tracks,
		User:       u,
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}

}
