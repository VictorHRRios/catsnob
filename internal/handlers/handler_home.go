package handlers

import (
	"context"
	"fmt"
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
		Error      string
	}

	tmplPath := filepath.Join("templates", "home", "artists.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "error parsing files", http.StatusInternalServerError)
		return
	}

	artists, err := cfg.Queries.GetArtists(context.Background())
	if err != nil {
		if err := tmpl.Execute(w, returnVals{Error: "Could not fetch artists"}); err != nil {
			http.Error(w, "error rendering template", http.StatusInternalServerError)
			return
		}
		return
	}
	respBody := returnVals{
		Artists: artists,
		User:    u,
	}

	if err := tmpl.Execute(w, respBody); err != nil {
		http.Error(w, "error rendering template", http.StatusInternalServerError)
		return
	}
}

func (cfg *ApiConfig) HandlerAlbums(w http.ResponseWriter, r *http.Request, u *database.User) {
	type returnVals struct {
		Stylesheet *string
		Albums     []database.Album
		User       *database.User
		Error      string
	}
	tmplPath := filepath.Join("templates", "home", "albums.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "error parsing files", http.StatusInternalServerError)
		return
	}

	albums, err := cfg.Queries.GetAlbums(context.Background())
	if err != nil {
		if err := tmpl.Execute(w, returnVals{Error: "Could not fetch albums"}); err != nil {
			http.Error(w, "error rendering template", http.StatusInternalServerError)
			return
		}
		return
	}
	respBody := returnVals{
		Stylesheet: nil,
		Albums:     albums,
		User:       u,
	}
	if err := tmpl.Execute(w, respBody); err != nil {
		http.Error(w, "error rendering template", http.StatusInternalServerError)
		return
	}
}

func (cfg *ApiConfig) HandlerTracks(w http.ResponseWriter, r *http.Request, u *database.User) {
	type returnVals struct {
		Stylesheet *string
		Tracks     []database.GetTracksRow
		User       *database.User
		Error      string
	}
	tmplPath := filepath.Join("templates", "home", "tracks.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "error parsing files", http.StatusInternalServerError)
		return
	}

	tracks, err := cfg.Queries.GetTracks(context.Background())
	if err != nil {
		if err := tmpl.Execute(w, returnVals{Error: "Could not fetch tracks"}); err != nil {
			http.Error(w, "error rendering template", http.StatusInternalServerError)
			return
		}
		return
	}
	respBody := returnVals{
		Stylesheet: nil,
		Tracks:     tracks,
		User:       u,
	}
	if err := tmpl.Execute(w, respBody); err != nil {
		http.Error(w, "error rendering template", http.StatusInternalServerError)
		return
	}
}

func (cfg *ApiConfig) HandlerLists(w http.ResponseWriter, r *http.Request, u *database.User) {
	fmt.Println("HandlerList ejecutado con el user", u)
	type returnVals struct {
		Stylesheet *string
		UserLists  []database.GetUserListsRow
		User       *database.User
		Error      string
	}
	tmplPath := filepath.Join("templates", "home", "lists.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "error parsing files", http.StatusInternalServerError)
		return
	}

	// Esto lo voy a modificar, es un placeholder al igual que la queri
	uLists, err := cfg.Queries.GetUserLists(context.Background(), u.ID)
	if err != nil || len(uLists) == 0 {
		w.WriteHeader(http.StatusNotFound)
		if err := tmpl.Execute(w, returnVals{Error: "No album lists found"}); err != nil {
			fmt.Println("Error al renderizar la plantilla:", err)
		}
		return
	}
	fmt.Println("Listas de álbumes obtenidas:", uLists)

	respBody := returnVals{
		Stylesheet: nil,
		UserLists:  uLists,
		User:       u,
	}
	if err := tmpl.Execute(w, respBody); err != nil {
		http.Error(w, "error rendering template", http.StatusInternalServerError)
		return
	}
}
