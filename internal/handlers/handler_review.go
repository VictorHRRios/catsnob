package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
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

func (cfg *ApiConfig) HandlerDeleteAlbumReview(w http.ResponseWriter, r *http.Request, u *database.User) {
	type DeleteRequest struct {
		ID uuid.UUID `json:"reviewId"`
	}
	var req DeleteRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = cfg.Queries.DeleteReview(context.Background(), req.ID)
	if err != nil {
		http.Error(w, "Could not delete request", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"success"}`))
}

func (cfg *ApiConfig) HandlerUpdateAlbumReview(w http.ResponseWriter, r *http.Request, u *database.User) {
	type UpdateRequest struct {
		ID      uuid.UUID `json:"id"`
		Title   string    `json:"title"`
		Content string    `json:"content"`
		Score   string    `json:"rating"`
	}
	var req UpdateRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	log.Print(req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = cfg.Queries.UpdateReview(context.Background(), database.UpdateReviewParams{
		Title:  sql.NullString{String: req.Title, Valid: true},
		Review: sql.NullString{String: req.Content, Valid: true},
		Score:  req.Score,
		ID:     req.ID,
	})
	if err != nil {
		http.Error(w, "Could not update request", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"success"}`))
}

func (cfg *ApiConfig) HandlerUserReview(w http.ResponseWriter, r *http.Request, u *database.User) {
	type returnVals struct {
		Error       string
		Stylesheet  *string
		User        *database.User
		ProfileUser *database.User
		Review      *database.GetReviewRow
		Title       string
		AlbumReview string
	}
	tmplPath := filepath.Join("templates", "review", "album.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "error parsing files", http.StatusInternalServerError)
		return
	}
	reviewID, err := uuid.Parse(r.PathValue("reviewid"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := tmpl.Execute(w, returnVals{Error: fmt.Sprintf("Error: %v", err)}); err != nil {
			http.Error(w, "error parsing files", http.StatusInternalServerError)
			return
		}
		return
	}

	review, err := cfg.Queries.GetReview(context.Background(), reviewID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := tmpl.Execute(w, returnVals{Error: fmt.Sprintf("Error: %v", err)}); err != nil {
			http.Error(w, "error parsing files", http.StatusInternalServerError)
			return
		}
		return
	}
	returnBody := returnVals{
		User:        u,
		Review:      &review,
		Title:       review.Title.String,
		AlbumReview: review.Review.String,
	}
	if err := tmpl.Execute(w, returnBody); err != nil {
		http.Error(w, "error rendering template", http.StatusInternalServerError)
	}
}
