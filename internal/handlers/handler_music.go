package handlers

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/VictorHRRios/catsnob/internal/database"
)

func (cfg *ApiConfig) HandlerArtistProfile(w http.ResponseWriter, r *http.Request, u *database.User) {
	type returnVals struct {
		Error      string
		Stylesheet *string
		Artist     database.Artist
		Albums     []database.GetArtistAlbumsRow
		User       *database.User
	}
	tmplPath := filepath.Join("templates", "music", "artist.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		tmpl.Execute(w, returnVals{Error: fmt.Sprintf("%v", err)})
		return
	}
	artistName := r.PathValue("artist")
	artist, err := cfg.Queries.GetArtist(context.Background(), artistName)
	if err != nil {
		tmpl.Execute(w, returnVals{Error: fmt.Sprintf("%v", err)})
		return
	}

	artistAlbums, err := cfg.Queries.GetArtistAlbums(context.Background(), artistName)
	if err != nil {
		tmpl.Execute(w, returnVals{Error: fmt.Sprintf("%v", err)})
		return
	}

	returnBody := returnVals{
		Stylesheet: nil,
		Artist:     artist,
		Albums:     artistAlbums,
		User:       u,
	}

	err = tmpl.Execute(w, returnBody)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func (cfg *ApiConfig) HandlerAlbum(w http.ResponseWriter, r *http.Request, u *database.User) {
	type returnVals struct {
		Error      string
		Stylesheet *string
		Tracks     []database.GetAlbumTracksRow
		User       *database.User
		Album      database.Album
		ArtistName string
	}
	tmplPath := filepath.Join("templates", "music", "album.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		if err := tmpl.Execute(w, returnVals{Error: fmt.Sprintf("%v", err)}); err != nil {
			log.Print(err)
		}
		return
	}
	albumName := r.PathValue("album")
	artistName := r.PathValue("artist")

	album, err := cfg.Queries.GetAlbum(context.Background(), albumName)
	if err != nil {
		if err := tmpl.Execute(w, returnVals{Error: fmt.Sprintf("Error fetching albums")}); err != nil {
			log.Print(err)
		}
		return
	}

	tracks, err := cfg.Queries.GetAlbumTracks(context.Background(), albumName)
	if err != nil {
		if err := tmpl.Execute(w, returnVals{Error: fmt.Sprintf("Error fetching tracks")}); err != nil {
			log.Print(err)
		}
		return
	}
	returnBody := returnVals{
		Stylesheet: nil,
		Tracks:     tracks,
		User:       u,
		Album:      album,
		ArtistName: artistName,
	}
	err = tmpl.Execute(w, returnBody)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func (cfg *ApiConfig) HandlerTrack(w http.ResponseWriter, r *http.Request, u *database.User) {
	type returnVals struct {
		Error      string
		Stylesheet *string
		Track      database.GetTrackRow
		User       *database.User
	}
	tmplPath := filepath.Join("templates", "music", "track.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		tmpl.Execute(w, returnVals{Error: fmt.Sprintf("%v", err)})
		return
	}
	trackName := r.PathValue("track")

	track, err := cfg.Queries.GetTrack(context.Background(), trackName)
	if err != nil {
		tmpl.Execute(w, returnVals{Error: fmt.Sprintf("Error fetching track")})
		return
	}

	returnBody := returnVals{
		Stylesheet: nil,
		Track:      track,
		User:       u,
	}

	err = tmpl.Execute(w, returnBody)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
