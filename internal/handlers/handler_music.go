package handlers

import (
	"context"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/VictorHRRios/catsnob/internal/database"
)

func (cfg *ApiConfig) HandlerArtistProfile(w http.ResponseWriter, r *http.Request, u *database.User) {
	artistName := r.PathValue("artist")
	artist, err := cfg.Queries.GetArtist(context.Background(), artistName)
	if err != nil {
		http.Error(w, "Error fetching artist", http.StatusInternalServerError)
		return
	}

	artistAlbums, err := cfg.Queries.GetArtistAlbums(context.Background(), artistName)
	if err != nil {
		http.Error(w, "Error fetching artist albums", http.StatusInternalServerError)
		return
	}

	tmplPath := filepath.Join("templates", "music", "artist.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
	}
	data := struct {
		Stylesheet *string
		Artist     database.Artist
		Albums     []database.GetArtistAlbumsRow
		User       *database.User
	}{
		Stylesheet: nil,
		Artist:     artist,
		Albums:     artistAlbums,
		User:       u,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func (cfg *ApiConfig) HandlerAlbum(w http.ResponseWriter, r *http.Request, u *database.User) {
	albumName := r.PathValue("album")
	album, err := cfg.Queries.GetAlbumTracks(context.Background(), albumName)
	if err != nil {
		http.Error(w, "Error fetching artist", http.StatusInternalServerError)
		return
	}

	tmplPath := filepath.Join("templates", "music", "album.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
	}
	data := struct {
		Stylesheet *string
		Tracks     []database.GetAlbumTracksRow
		User       *database.User
	}{
		Stylesheet: nil,
		Tracks:     album,
		User:       u,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
