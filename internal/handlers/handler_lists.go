package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/VictorHRRios/catsnob/internal/database"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) HandlerCreate_List(w http.ResponseWriter, r *http.Request, u *database.User) {
	fmt.Println("HandlerCreate_list ejecutado")
	type returnVals struct {
		Stylesheet *string
		User       *database.User
		Error      string
	}
	tmplPath := filepath.Join("templates", "lists", "create_list.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "error parsing files", http.StatusInternalServerError)
		return
	}

	respBody := returnVals{
		Stylesheet: nil,
		User:       u,
	}
	if err := tmpl.Execute(w, respBody); err != nil {
		http.Error(w, "error rendering template", http.StatusInternalServerError)
		return
	}
}

func (cfg *ApiConfig) HandlerEdit_List(w http.ResponseWriter, r *http.Request, u *database.User) {
	fmt.Println("HandlerEdit_list ejecutado")
	type returnVals struct {
		Stylesheet *string
		Playlist   []database.GetAlbumsFromListRow
		PlaylistID string
		User       *database.User
		Error      string
	}
	tmplPath := filepath.Join("templates", "lists", "edit_list.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "error parsing files", http.StatusInternalServerError)
		return
	}

	listID, err := uuid.Parse(r.PathValue("listid"))

	if err != nil {
		if err := tmpl.Execute(w, returnVals{Error: fmt.Sprintf("%v", err)}); err != nil {
			http.Error(w, "error rendering template", http.StatusInternalServerError)
			return
		}
		return
	}

	albums, err := cfg.Queries.GetAlbumsFromList(context.Background(), listID)
	if err != nil || len(albums) == 0 {
		w.WriteHeader(http.StatusNotFound)
		tmpl.Execute(w, returnVals{PlaylistID: listID.String(), Error: "No albums found"})
		return
	}
	fmt.Println("Lista de álbumes obtenidas en edit_list:", albums)

	respBody := returnVals{
		Stylesheet: nil,
		Playlist:   albums,
		PlaylistID: listID.String(),
		User:       u,
	}
	if err := tmpl.Execute(w, respBody); err != nil {
		http.Error(w, "error rendering template", http.StatusInternalServerError)
		return
	}

}

func (cfg *ApiConfig) HandlerCreateAlbumList(w http.ResponseWriter, r *http.Request, u *database.User) {
	type returnVals struct {
		Error      string
		Stylesheet *string
		User       *database.User
		List       *database.AlbumList
	}

	if u == nil {
		http.Error(w, "User is not authenticated", http.StatusUnauthorized)
		return
	}

	tmplPath := filepath.Join("templates", "home", "create_list.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "error parsing files", http.StatusInternalServerError)
		return
	}

	title := r.FormValue("titleList")

	if title == "" {
		w.WriteHeader(http.StatusBadRequest)
		tmpl.Execute(w, returnVals{Error: "Title cannot be empty", User: u})
		return
	}

	newList, err := cfg.Queries.CreateAlbumList(context.Background(), database.CreateAlbumListParams{
		UserID: u.ID,
		Title:  sql.NullString{String: title, Valid: title != ""},
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := tmpl.Execute(w, returnVals{Error: fmt.Sprintf("%v", err)}); err != nil {
			http.Error(w, "error parsing files", http.StatusInternalServerError)
			return
		}
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/app/edit_list/%v", newList.ID), http.StatusFound)
}

func (cfg *ApiConfig) HandlerAddAlbumList(w http.ResponseWriter, r *http.Request, u *database.User) {
	type returnVals struct {
		Error      string
		Stylesheet *string
		User       *database.User
		List       *database.AlbumList
	}

	if u == nil {
		http.Error(w, "User is not authenticated", http.StatusUnauthorized)
		return
	}

	tmplPath := filepath.Join("templates", "home", "edit_list.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "error parsing files", http.StatusInternalServerError)
		return
	}

	// albumid := r.FormValue("albumID")

	newList, err := cfg.Queries.CreateAlbumList(context.Background(), database.CreateAlbumListParams{
		UserID: u.ID,
		// Title:  sql.NullString{String: title, Valid: title != ""},
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := tmpl.Execute(w, returnVals{Error: fmt.Sprintf("%v", err)}); err != nil {
			http.Error(w, "error parsing files", http.StatusInternalServerError)
			return
		}
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/app/edit_list/%v/%v", newList.ID, true), http.StatusFound)
}
