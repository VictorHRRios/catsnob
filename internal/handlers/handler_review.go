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

func (cfg *ApiConfig) HandlerCreateAlbumReview(w http.ResponseWriter, r *http.Request, u *database.User) {
	type returnVals struct {
		Error      string
		Stylesheet *string
		User       *database.User
	}
	tmplPath := filepath.Join("templates", "music", "album.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tmpl.Execute(w, returnVals{Error: fmt.Sprintf("%v", err)})
		return
	}

	rating := r.FormValue("rating")
	albumName := r.FormValue("album")
	artistName := r.FormValue("artist")
	log.Print(albumName)
	album, err := cfg.Queries.GetAlbum(context.Background(), albumName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tmpl.Execute(w, returnVals{Error: fmt.Sprintf("Album does not exist: %v", err)})
		return
	}
	_, err = cfg.Queries.CreateReviewShort(context.Background(), database.CreateReviewShortParams{
		UserID:  u.ID,
		AlbumID: album.ID,
		Score:   rating,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tmpl.Execute(w, returnVals{Error: fmt.Sprintf("%v", err)})
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/app/music/%v/%v", artistName, albumName), http.StatusFound)
}
