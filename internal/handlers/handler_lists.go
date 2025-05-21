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

// Handler para mostrar el formulario de creación de lista
func (cfg *ApiConfig) HandlerCreate_List(w http.ResponseWriter, r *http.Request, u *database.User) {
    tmplPath := filepath.Join("templates", "lists", "create_list.html")
    tmpl, err := template.ParseFiles(layout, tmplPath)
    if err != nil {
        http.Error(w, "error parsing files", http.StatusInternalServerError)
        return
    }
    type returnVals struct {
        Stylesheet *string
        User       *database.User
        Error      string
    }
    respBody := returnVals{User: u}
    tmpl.Execute(w, respBody)
}

// Handler para crear una lista (album o track)
func (cfg *ApiConfig) HandlerCreateList(w http.ResponseWriter, r *http.Request, u *database.User) {
    listType := r.FormValue("type") // "album" o "track"
    name := r.FormValue("titleList")
    description := r.FormValue("descriptionList")

    if listType != "album" && listType != "track" {
        http.Error(w, "Invalid list type", http.StatusBadRequest)
        return
    }

    newList, err := cfg.Queries.CreateUserList(context.Background(), database.CreateUserListParams{
        Name:        sql.NullString{String: name, Valid: name != ""},
        Description: sql.NullString{String: description, Valid: description != ""},
        Type:        sql.NullString{String: listType, Valid: true},
        UserID:      u.ID,
    })
    if err != nil {
        http.Error(w, "Error creating new user list", http.StatusBadRequest)
        return
    }

    if listType == "album" {
        http.Redirect(w, r, fmt.Sprintf("/app/lists/edit_list/%v", newList.IDPlaylistA), http.StatusFound)
    } else {
        http.Redirect(w, r, fmt.Sprintf("/app/lists/edit_list_tracks/%v", newList.IDPlaylistA), http.StatusFound)
    }
}

// Handler para editar una lista (album o track)
func (cfg *ApiConfig) HandlerEdit_List(w http.ResponseWriter, r *http.Request, u *database.User) {
    listID, err := uuid.Parse(r.PathValue("listid"))
    if err != nil {
        http.Error(w, "Invalid list ID", http.StatusBadRequest)
        return
    }

    // Detectar tipo de lista
    list, err := cfg.Queries.GetListByID(context.Background(), listID)
    if err != nil {
        http.Error(w, "List not found", http.StatusNotFound)
        return
    }

    if list.Type.String == "album" {
        // Editar lista de álbumes
        albums, err := cfg.Queries.GetAlbumsFromList(context.Background(), listID)
        if err != nil {
            http.Error(w, "Error fetching albums", http.StatusInternalServerError)
            return
        }
        tmplPath := filepath.Join("templates", "lists", "edit_list.html")
        tmpl, err := template.ParseFiles(layout, tmplPath)
        if err != nil {
            http.Error(w, "Error parsing template", http.StatusInternalServerError)
            return
        }
        type returnVals struct {
            Stylesheet *string
            Playlist   []database.GetAlbumsFromListRow
            PlaylistID string
            User       *database.User
            Error      string
        }
        respBody := returnVals{
            Playlist:   albums,
            PlaylistID: listID.String(),
            User:       u,
        }
        tmpl.Execute(w, respBody)
    } else {
        // Editar lista de tracks
        tracks, err := cfg.Queries.GetTracksFromList(context.Background(), listID)
        if err != nil {
            http.Error(w, "Error fetching tracks", http.StatusInternalServerError)
            return
        }
        tmplPath := filepath.Join("templates", "lists", "edit_list_tracks.html")
        tmpl, err := template.ParseFiles(layout, tmplPath)
        if err != nil {
            http.Error(w, "Error parsing template", http.StatusInternalServerError)
            return
        }
        type returnVals struct {
            Stylesheet *string
            Playlist   []database.GetTracksFromListRow
            PlaylistID string
            User       *database.User
            Error      string
        }
        respBody := returnVals{
            Playlist:   tracks,
            PlaylistID: listID.String(),
            User:       u,
        }
        tmpl.Execute(w, respBody)
    }
}
// Handler para agregar elementos a una lista (album o track)
func (cfg *ApiConfig) HandlerAdd_Items(w http.ResponseWriter, r *http.Request, u *database.User) {
    listID, err := uuid.Parse(r.PathValue("listid"))
    if err != nil {
        http.Error(w, "Invalid list ID", http.StatusBadRequest)
        return
    }
    list, err := cfg.Queries.GetListByID(context.Background(), listID)
    if err != nil {
        http.Error(w, "List not found", http.StatusNotFound)
        return
    }

    if list.Type.String == "album" {
        albums, _ := cfg.Queries.GetAlbumsNotInList(context.Background(), listID)
        list_name, _ := cfg.Queries.GetListName(context.Background(), listID)
        tmplPath := filepath.Join("templates", "lists", "add_albums.html")
        tmpl, _ := template.ParseFiles(layout, tmplPath)
        type returnVals struct {
            Stylesheet    *string
            Albums        []database.GetAlbumsNotInListRow
            Playlist_name string
            PlaylistID    string
            User          *database.User
            Error         string
        }
        respBody := returnVals{
            Albums:        albums,
            Playlist_name: list_name[0].String,
            PlaylistID:    listID.String(),
            User:          u,
        }
        tmpl.Execute(w, respBody)
    } else {
        tracks, _ := cfg.Queries.GetTracksNotInList(context.Background(), listID)
        list_name, _ := cfg.Queries.GetTrackListName(context.Background(), listID)
        tmplPath := filepath.Join("templates", "lists", "add_tracks.html")
        tmpl, _ := template.ParseFiles(layout, tmplPath)
        type returnVals struct {
            Stylesheet    *string
            Tracks        []database.GetTracksNotInListRow
            Playlist_name string
            PlaylistID    string
            User          *database.User
            Error         string
        }
        respBody := returnVals{
            Tracks:        tracks,
            Playlist_name: list_name.String,
            PlaylistID:    listID.String(),
            User:          u,
        }
        tmpl.Execute(w, respBody)
    }
}

// Handler para agregar elementos seleccionados a la lista
func (cfg *ApiConfig) HandlerAddItemsToList(w http.ResponseWriter, r *http.Request, u *database.User) {
    listID, err := uuid.Parse(r.PathValue("listid"))
    if err != nil {
        http.Error(w, "Invalid list ID", http.StatusBadRequest)
        return
    }
    list, err := cfg.Queries.GetListByID(context.Background(), listID)
    if err != nil {
        http.Error(w, "List not found", http.StatusNotFound)
        return
    }
    err = r.ParseForm()
    if err != nil {
        http.Error(w, "Error form not complete", http.StatusBadRequest)
        return
    }

    if list.Type.String == "album" {
        albums := r.Form["album_ids"]
        for _, albumID := range albums {
            aid, _ := uuid.Parse(albumID)
            cfg.Queries.AddAlbumToList(context.Background(), database.AddAlbumToListParams{
                UserListsID: listID,
                AlbumID:     aid,
            })
        }
        http.Redirect(w, r, fmt.Sprintf("/app/lists/edit_list/%v", listID), http.StatusFound)
    } else {
        tracks := r.Form["track_ids"]
        for _, trackID := range tracks {
            tid, _ := uuid.Parse(trackID)
            cfg.Queries.AddTrackToList(context.Background(), database.AddTrackToListParams{
                UserListsID: listID,
                TrackID:     tid,
            })
        }
        http.Redirect(w, r, fmt.Sprintf("/app/lists/edit_list_tracks/%v", listID), http.StatusFound)
    }
}

// Handler para eliminar elementos de la lista (album o track)
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
// Handler para eliminar la lista completa
func (cfg *ApiConfig) HandlerDeleteList(w http.ResponseWriter, r *http.Request, u *database.User) {
    listID, err := uuid.Parse(r.PathValue("listid"))
    if err != nil {
        http.Error(w, "Error list id can not be null", http.StatusBadRequest)
        return
    }
    err = cfg.Queries.DeleteList(context.Background(), listID)
    http.Redirect(w, r, fmt.Sprintf("/app/home/lists"), http.StatusFound)
}

// Handler para eliminar elementos de la lista (track)


func (cfg *ApiConfig) HandlerDelete_Tracks(w http.ResponseWriter, r *http.Request, u *database.User) {
	fmt.Println("HandlerAdd_Tracks ejecutado")
	type returnVals struct {
		Stylesheet    *string
		Tracks        []database.GetTracksFromListRow
		Playlist_name string
		PlaylistID    string
		User          *database.User
		Error         string
	}
	tmplPath := filepath.Join("templates", "lists", "delete_tracks.html")
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

	list_name, _ := cfg.Queries.GetTrackListName(context.Background(), listID)

	tracks, err := cfg.Queries.GetTracksFromList(context.Background(), listID)
	if err != nil || len(tracks) == 0 {
		w.WriteHeader(http.StatusNotFound)
		tmpl.Execute(w, returnVals{PlaylistID: listID.String(), Error: "No tracks found"})
		return
	}

	respBody := returnVals{
		Stylesheet:    nil,
		Tracks:        tracks,
		Playlist_name: list_name.String,
		PlaylistID:    listID.String(),
		User:          u,
	}
	if err := tmpl.Execute(w, respBody); err != nil {
		http.Error(w, "error rendering template", http.StatusInternalServerError)
		return
	}

}
func (cfg *ApiConfig) HandlerDeleteTracksFromList(w http.ResponseWriter, r *http.Request, u *database.User) {
	fmt.Println("HandlerDeleteTracksFromList ejecutado")

	err := r.ParseForm() // Analizo los datos del formulario
	if err != nil {
		http.Error(w, "Error form not complete", http.StatusBadRequest)
		return
	}

	listID, err := uuid.Parse(r.PathValue("listid"))

	tracks := r.Form["track_ids"]

	if len(tracks) == 0 {
		http.Error(w, "Error tracks selecteds can not be null", http.StatusBadRequest)
		return
	}

	for _, trackID := range tracks {
		tid, _ := uuid.Parse(trackID)
		err := cfg.Queries.DeleteTrackFromList(context.Background(), database.DeleteTrackFromListParams{
			UserListsID: listID,
			TrackID:     tid,
		})
		if err == nil {
			fmt.Println("Track correctamente elimiando.")
		}
	}

	http.Redirect(w, r, fmt.Sprintf("/app/lists/edit_list_tracks/%v", listID), http.StatusFound)
}	