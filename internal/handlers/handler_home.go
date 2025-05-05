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
	type returnVals struct {
		Stylesheet *string
		Artists    []database.Artist
		User       *database.User
		Error      error
	}

	tmplPath := filepath.Join("templates", "home", "artists.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		tmpl.Execute(w, returnVals{User: u, Error: err})
		return
	}

	artists, err := cfg.Queries.GetTop12Artists(context.Background())
	respBody := returnVals{
		Stylesheet: nil,
		Artists:    artists,
		User:       u,
		Error:      err,
	}

	tmpl.Execute(w, respBody)
}

func (cfg *ApiConfig) HandlerAlbums(w http.ResponseWriter, r *http.Request, u *database.User) {
	type returnVals struct {
		Stylesheet *string
		Albums     []database.GetTop12AlbumsRow
		User       *database.User
		Error      error
	}
	tmplPath := filepath.Join("templates", "home", "albums.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		tmpl.Execute(w, returnVals{User: u, Error: err})
		return
	}

	albums, err := cfg.Queries.GetTop12Albums(context.Background())
	respBody := returnVals{
		Stylesheet: nil,
		Albums:     albums,
		User:       u,
		Error:      err,
	}
	err = tmpl.Execute(w, respBody)
}

func (cfg *ApiConfig) HandlerTracks(w http.ResponseWriter, r *http.Request, u *database.User) {
	type returnVals struct {
		Stylesheet *string
		Tracks     []database.GetTop12TracksRow
		User       *database.User
		Error      error
	}
	tmplPath := filepath.Join("templates", "home", "tracks.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		tmpl.Execute(w, returnVals{User: u, Error: err})
		return
	}

	tracks, err := cfg.Queries.GetTop12Tracks(context.Background())
	respBody := returnVals{
		Stylesheet: nil,
		Tracks:     tracks,
		User:       u,
		Error:      err,
	}
	tmpl.Execute(w, respBody)
}
