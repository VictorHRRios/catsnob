package handlers

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/VictorHRRios/catsnob/internal/database"
	"github.com/google/uuid"
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
		http.Error(w, "error parsing files", http.StatusInternalServerError)
		return
	}

	rating := r.FormValue("rating")
	albumID, err := uuid.Parse(r.FormValue("albumid"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := tmpl.Execute(w, returnVals{Error: fmt.Sprintf("Error: %v", err)}); err != nil {
			http.Error(w, "error parsing files", http.StatusInternalServerError)
			return
		}
		return
	}
	album, err := cfg.Queries.GetAlbum(context.Background(), albumID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		if err := tmpl.Execute(w, returnVals{Error: "Album not found"}); err != nil {
			http.Error(w, "error parsing files", http.StatusInternalServerError)
			return
		}
		return
	}
	_, err = cfg.Queries.CreateReviewShort(context.Background(), database.CreateReviewShortParams{
		UserID:  u.ID,
		AlbumID: album.ID,
		Score:   rating,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := tmpl.Execute(w, returnVals{Error: fmt.Sprintf("%v", err)}); err != nil {
			http.Error(w, "error parsing files", http.StatusInternalServerError)
			return
		}
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/app/album/%v", albumID), http.StatusFound)
}
