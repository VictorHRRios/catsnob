package handlers

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/VictorHRRios/catsnob/internal/database"
)

func (cfg *ApiConfig) HandlerArtistProfile(w http.ResponseWriter, r *http.Request) {
	artistName := r.PathValue("artist")
	artist, err := cfg.Queries.GetArtist(context.Background(), artistName)
	if err != nil {
		http.Error(w, "Error fetching artist", http.StatusInternalServerError)
		return
	}

	artistAlbums, err := cfg.Queries.GetArtistAlbums(context.Background(), artistName)
	fmt.Printf("artistAlbums: %v\n", artistAlbums)
	if err != nil {
		http.Error(w, "Error fetching artist albums", http.StatusInternalServerError)
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
		Albums     []database.GetArtistAlbumsRow
	}{
		Stylesheet: nil,
		Name:       artist.Name,
		Img:        artist.ImgUrl,
		Biography:  artist.Biography.String,
		FormedAt:   artist.FormedAt,
		Genre:      artist.Genre,
		Albums:     artistAlbums,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func (cfg *ApiConfig) HandlerAlbum(w http.ResponseWriter, r *http.Request) {
	albumName := r.PathValue("album")
	album, err := cfg.Queries.GetAlbumTracks(context.Background(), albumName)
	if err != nil {
		http.Error(w, "Error fetching artist", http.StatusInternalServerError)
		return
	}

	tmplPath := filepath.Join("templates", "album.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
	}
	data := struct {
		Stylesheet *string
		Tracks     []database.GetAlbumTracksRow
	}{
		Stylesheet: nil,
		Tracks:     album,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
