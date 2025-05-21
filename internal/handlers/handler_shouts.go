package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/VictorHRRios/catsnob/internal/database"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) HandlerCreateShout(w http.ResponseWriter, r *http.Request, u *database.User) {
	type returnVals struct {
		Error      string
		User       *database.User
		Stylesheet *string
	}

	tmplPath := filepath.Join("templates", "review", "album.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "error parsing files", http.StatusInternalServerError)
		return
	}

	reviewIDStr := r.FormValue("reviewid")
	reviewID, err := uuid.Parse(reviewIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		tmpl.Execute(w, returnVals{Error: "ID ", User: u})
		return
	}

	title := r.FormValue("title")
	shoutText := r.FormValue("shout_text")

	_, err = cfg.Queries.CreateShouts(r.Context(), database.CreateShoutsParams{
		UserID:    u.ID,
		ReviewID:  reviewID,
		Title:     title,
		ShoutText: shoutText,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tmpl.Execute(w, returnVals{Error: fmt.Sprintf("Error al crear shout: %v", err), User: u})
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/app/user/%s/review/%s", u.ID, reviewID), http.StatusFound)
}

// DELETE SHOUT
func (cfg *ApiConfig) HandlerDeleteShout(w http.ResponseWriter, r *http.Request, u *database.User) {
	type DeleteRequest struct {
		ID uuid.UUID `json:"shoutId"`
	}
	var req DeleteRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	log.Print(req)

	err = cfg.Queries.DeleteShout(context.Background(), req.ID)
	if err != nil {
		http.Error(w, "Could not delete request", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"success"}`))
}

// UPDATE SHOUT
func (cfg *ApiConfig) HandlerUpdateShout(w http.ResponseWriter, r *http.Request, u *database.User) {
	type UpdateRequest struct {
		ID        uuid.UUID `json:"id"`
		Title     string    `json:"title"`
		ShoutText string    `json:"shout_text"`
	}
	var req UpdateRequest

	// parsear JSON
	err := json.NewDecoder(r.Body).Decode(&req)
	log.Print(req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = cfg.Queries.UpdateShout(context.Background(), database.UpdateShoutParams{
		ID:        req.ID,
		Title:     req.Title,
		ShoutText: req.ShoutText,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		http.Error(w, "Could not update shout", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"success"}`))
}
