package handlers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/VictorHRRios/catsnob/internal/database"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) HandlerArtistProfile(w http.ResponseWriter, r *http.Request, u *database.User) {
	type returnVals struct {
		Error      string
		Stylesheet *string
		Artist     database.Artist
		Albums     []database.GetArtistAlbumsRow
		User       *database.User
		Bio        string
	}
	tmplPath := filepath.Join("templates", "music", "artist.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "error parsing files", http.StatusInternalServerError)
		return
	}
	artistID, err := uuid.Parse(r.PathValue("artistid"))
	if err != nil {
		if err := tmpl.Execute(w, returnVals{Error: fmt.Sprintf("%v", err)}); err != nil {
			http.Error(w, "error rendering template", http.StatusInternalServerError)
			return
		}
		return
	}
	artist, err := cfg.Queries.GetArtist(context.Background(), artistID)
	if err != nil {
		if err := tmpl.Execute(w, returnVals{Error: fmt.Sprintf("%v", err)}); err != nil {
			http.Error(w, "error rendering template", http.StatusInternalServerError)
			return
		}
		return
	}

	artistAlbums, err := cfg.Queries.GetArtistAlbums(context.Background(), artistID)
	if err != nil {
		if err := tmpl.Execute(w, returnVals{Error: fmt.Sprintf("%v", err)}); err != nil {
			http.Error(w, "error rendering template", http.StatusInternalServerError)
			return
		}
		return
	}

	returnBody := returnVals{
		Artist: artist,
		Albums: artistAlbums,
		User:   u,
		Bio:    artist.Biography.String,
	}

	if err = tmpl.Execute(w, returnBody); err != nil {
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
		UserReview *database.AlbumReview
		Reviews    []database.GetReviewByAlbumRow
	}
	tmplPath := filepath.Join("templates", "music", "album.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "error parsing files", http.StatusInternalServerError)
		return
	}
	albumID, err := uuid.Parse(r.PathValue("albumid"))
	if err != nil {
		if err := tmpl.Execute(w, returnVals{Error: fmt.Sprintf("%v", err)}); err != nil {
			http.Error(w, "error rendering template", http.StatusInternalServerError)
			return
		}
		return
	}

	album, err := cfg.Queries.GetAlbum(context.Background(), albumID)
	if err != nil {
		if err := tmpl.Execute(w, returnVals{Error: fmt.Sprintf("Error fetching albums: %v", err)}); err != nil {
			http.Error(w, "error rendering template", http.StatusInternalServerError)
			return
		}
		return
	}

	tracks, err := cfg.Queries.GetAlbumTracks(context.Background(), albumID)
	if err != nil {
		if err := tmpl.Execute(w, returnVals{Error: fmt.Sprintf("Error fetching tracks: %v", err)}); err != nil {
			http.Error(w, "error rendering template", http.StatusInternalServerError)
			return
		}
		return
	}

	reviews, err := cfg.Queries.GetReviewByAlbum(context.Background(), albumID)
	if err != nil {
		if err := tmpl.Execute(w, returnVals{Error: fmt.Sprintf("Error fetching tracks: %v", err)}); err != nil {
			http.Error(w, "error rendering template", http.StatusInternalServerError)
			return
		}
		return
	}

	returnBody := returnVals{
		Tracks:  tracks,
		User:    u,
		Album:   album,
		Reviews: reviews,
	}
	if u == nil {
		if err = tmpl.Execute(w, returnBody); err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
		return
	}

	userReview, err := cfg.Queries.GetReviewByUserAlbum(context.Background(), database.GetReviewByUserAlbumParams{
		AlbumID: album.ID,
		UserID:  u.ID,
	})
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		if err := tmpl.Execute(w, returnVals{Error: fmt.Sprintf("Error fetching user review: %v", err)}); err != nil {
			http.Error(w, "error rendering template", http.StatusInternalServerError)
			return
		}
		return
	}
	if !errors.Is(err, sql.ErrNoRows) {
		returnBody.UserReview = &userReview
	}

	if err = tmpl.Execute(w, returnBody); err != nil {
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
		http.Error(w, "error parsing files", http.StatusInternalServerError)
		return
	}
	trackID, err := uuid.Parse(r.PathValue("trackid"))
	if err != nil {
		if err := tmpl.Execute(w, returnVals{Error: fmt.Sprintf("%v", err)}); err != nil {
			http.Error(w, "error rendering template", http.StatusInternalServerError)
			return
		}
		return
	}

	track, err := cfg.Queries.GetTrack(context.Background(), trackID)
	if err != nil {
		if err := tmpl.Execute(w, returnVals{Error: "Error fetching track"}); err != nil {
			http.Error(w, "error rendering template", http.StatusInternalServerError)
			return
		}
		return
	}

	returnBody := returnVals{
		Track: track,
		User:  u,
	}

	if err = tmpl.Execute(w, returnBody); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
