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
		tmpl.Execute(w, returnVals{PlaylistID: listID.String(), User: u, Error: "No albums found"})
		return
	}

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

	name := r.FormValue("titleList")
	description := r.FormValue("descriptionList")

	newList, err := cfg.Queries.CreateUserList(context.Background(), database.CreateUserListParams{
		Name:        sql.NullString{String: name, Valid: name != ""},
		Description: sql.NullString{String: description, Valid: description != ""},
		Type:        sql.NullString{String: "list", Valid: description != ""},
		UserID:      u.ID,
	})
	if err != nil {
		http.Error(w, "Error creating new user list", http.StatusBadRequest)
	}

	http.Redirect(w, r, fmt.Sprintf("/app/lists/edit_list/%v", newList.IDPlaylistA), http.StatusFound)
}

func (cfg *ApiConfig) HandlerAdd_Albums(w http.ResponseWriter, r *http.Request, u *database.User) {
	fmt.Println("HandlerAdd_Albums ejecutado")
	type returnVals struct {
		Stylesheet    *string
		Albums        []database.GetAlbumsNotInListRow
		Playlist_name string
		PlaylistID    string
		User          *database.User
		Error         string
	}
	tmplPath := filepath.Join("templates", "lists", "add_albums.html")
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

	list_name, _ := cfg.Queries.GetListName(context.Background(), listID)

	albums, err := cfg.Queries.GetAlbumsNotInList(context.Background(), listID)
	if err != nil || len(albums) == 0 {
		w.WriteHeader(http.StatusNotFound)
		tmpl.Execute(w, returnVals{PlaylistID: listID.String(), Error: "No albums found"})
		return
	}

	respBody := returnVals{
		Stylesheet:    nil,
		Albums:        albums,
		Playlist_name: list_name[0].String,
		PlaylistID:    listID.String(),
		User:          u,
	}
	if err := tmpl.Execute(w, respBody); err != nil {
		http.Error(w, "error rendering template", http.StatusInternalServerError)
		return
	}

}

func (cfg *ApiConfig) HandlerAddAlbumsToList(w http.ResponseWriter, r *http.Request, u *database.User) {
	fmt.Println("HandlerAddAlbumsToList ejecutado")

	err := r.ParseForm() // Analizo los datos del formulario
	if err != nil {
		http.Error(w, "Error form not complete", http.StatusBadRequest)
		return
	}

	listID, err := uuid.Parse(r.PathValue("listid"))

	albums := r.Form["album_ids"]

	if len(albums) == 0 {
		http.Error(w, "Error albums selecteds can not be null", http.StatusBadRequest)
		return
	}

	for _, albumID := range albums {
		aid, _ := uuid.Parse(albumID)
		a, err := cfg.Queries.AddAlbumToList(context.Background(), database.AddAlbumToListParams{
			UserListsID: listID,
			AlbumID:     aid,
		})
		if err == nil {
			fmt.Println("Album correctamente agregado: ", a)
		}
	}

	http.Redirect(w, r, fmt.Sprintf("/app/lists/edit_list/%v", listID), http.StatusFound)
}

func (cfg *ApiConfig) HandlerDelete_Albums(w http.ResponseWriter, r *http.Request, u *database.User) {
	fmt.Println("HandlerAdd_Albums ejecutado")
	type returnVals struct {
		Stylesheet    *string
		Albums        []database.GetAlbumsFromListRow
		Playlist_name string
		PlaylistID    string
		User          *database.User
		Error         string
	}
	tmplPath := filepath.Join("templates", "lists", "delete_albums.html")
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

	list_name, _ := cfg.Queries.GetListName(context.Background(), listID)

	albums, err := cfg.Queries.GetAlbumsFromList(context.Background(), listID)
	if err != nil || len(albums) == 0 {
		w.WriteHeader(http.StatusNotFound)
		tmpl.Execute(w, returnVals{PlaylistID: listID.String(), Error: "No albums found"})
		return
	}

	respBody := returnVals{
		Stylesheet:    nil,
		Albums:        albums,
		Playlist_name: list_name[0].String,
		PlaylistID:    listID.String(),
		User:          u,
	}
	if err := tmpl.Execute(w, respBody); err != nil {
		http.Error(w, "error rendering template", http.StatusInternalServerError)
		return
	}

}

func (cfg *ApiConfig) HandlerDeleteAlbumsFromList(w http.ResponseWriter, r *http.Request, u *database.User) {
	fmt.Println("HandlerDeleteAlbumsFromList ejecutado")

	err := r.ParseForm() // Analizo los datos del formulario
	if err != nil {
		http.Error(w, "Error form not complete", http.StatusBadRequest)
		return
	}

	listID, err := uuid.Parse(r.PathValue("listid"))

	albums := r.Form["album_ids"]

	if len(albums) == 0 {
		http.Error(w, "Error albums selecteds can not be null", http.StatusBadRequest)
		return
	}

	for _, albumID := range albums {
		aid, _ := uuid.Parse(albumID)
		err := cfg.Queries.DeleteAlbumFromList(context.Background(), database.DeleteAlbumFromListParams{
			UserListsID: listID,
			AlbumID:     aid,
		})
		if err == nil {
			fmt.Println("Album correctamente elimiando.")
		}
	}

	http.Redirect(w, r, fmt.Sprintf("/app/lists/edit_list/%v", listID), http.StatusFound)
}

func (cfg *ApiConfig) HandlerDeleteList(w http.ResponseWriter, r *http.Request, u *database.User) {
	fmt.Println("HandlerDeleteList ejecutado")

	listID, err := uuid.Parse(r.PathValue("listid"))

	if err != nil {
		http.Error(w, "Error list id can not be null", http.StatusBadRequest)
		return
	}

	err = cfg.Queries.DeleteList(context.Background(), listID)
	if err == nil {
		fmt.Println("Lista correctamente elimiando.")
	}

	http.Redirect(w, r, fmt.Sprintf("/app/home/lists"), http.StatusFound)
}
