package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/VictorHRRios/catsnob/internal/api"
	"github.com/VictorHRRios/catsnob/internal/database"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) HandlerAdminIndex(w http.ResponseWriter, r *http.Request, u *database.User) {
	type returnVals struct {
		Error      string
		Stylesheet *string
		User       *database.User
		Artists    []database.Artist
	}
	tmplPath := filepath.Join("templates", "admin", "index.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "Error Parsing Files", http.StatusInternalServerError)
		return
	}
	if u == nil || !u.IsAdmin {
		w.WriteHeader(http.StatusForbidden)
		if err := tmpl.Execute(w, returnVals{Error: "Access denied for user"}); err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
		return
	}

	artists, err := cfg.Queries.GetArtists(context.Background())
	if err != nil {
		http.Error(w, "Error Retrieving Albums", http.StatusInternalServerError)
		return
	}

	returnBody := returnVals{
		User:    u,
		Artists: artists,
	}

	if err := tmpl.Execute(w, returnBody); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func (cfg *ApiConfig) HandlerDeleteArtist(w http.ResponseWriter, r *http.Request, u *database.User) {
	type DeleteRequest struct {
		ID uuid.UUID `json:"albumId"`
	}
	var req DeleteRequest

	if u == nil || !u.IsAdmin {
		http.Error(w, "Access Denied", http.StatusForbidden)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = cfg.Queries.DeleteArtist(context.Background(), req.ID)
	if err != nil {
		http.Error(w, "Could not delete artist", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"success"}`))
}

func (cfg *ApiConfig) HandlerFormArtistDisc(w http.ResponseWriter, r *http.Request, u *database.User) {
	type returnVals struct {
		Error      string
		Stylesheet *string
		User       *database.User
	}
	tmplPath := filepath.Join("templates", "admin", "registerArtist.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "Error Parsing Files", http.StatusInternalServerError)
		return
	}

	if u == nil {
		w.WriteHeader(http.StatusForbidden)
		if err := tmpl.Execute(w, returnVals{Error: "Access denied, no user logged in"}); err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
		return
	}
	if !u.IsAdmin {
		w.WriteHeader(http.StatusForbidden)
		if err := tmpl.Execute(w, returnVals{Error: "Access denied for user"}); err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
		return
	}

	returnBody := returnVals{
		User: u,
	}

	if err := tmpl.Execute(w, returnBody); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func (cfg *ApiConfig) HandlerCreateArtistDisc(w http.ResponseWriter, r *http.Request, u *database.User) {
	type returnVals struct {
		Error      string
		Stylesheet *string
		User       *database.User
	}
	tmplPath := filepath.Join("templates", "admin", "registerArtist.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "error parsing files", http.StatusInternalServerError)
		return
	}

	if u == nil || !u.IsAdmin {
		http.Error(w, "Access Denied", http.StatusForbidden)
		return
	}

	var validArtistBio bool

	name := r.FormValue("artist_id")
	retrArtist, err := api.GetArtist(&name)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		if err := tmpl.Execute(w, returnVals{Error: "Error searching for artist in api"}); err != nil {
			http.Error(w, "error rendering template", http.StatusInternalServerError)
			return
		}
		return
	}
	if retrArtist.Artists == nil {
		tmpl.Execute(w, returnVals{Error: "No artist found in https://www.theaudiodb.com/"})
		return
	}
	artist := retrArtist.Artists[0]

	if artist.StrBiographyEN == nil {
		validArtistBio = false
	} else {
		validArtistBio = true
	}

	artistDB, err := cfg.Queries.CreateArtist(context.Background(), database.CreateArtistParams{
		FormedAt:  artist.IntFormedYear,
		Name:      artist.StrArtist,
		Biography: sql.NullString{String: *artist.StrBiographyEN, Valid: validArtistBio},
		Genre:     artist.StrGenre,
		ImgUrl:    artist.StrArtistThumb,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tmpl.Execute(w, returnVals{Error: fmt.Sprintf("Error creating artist: %v", err)})
		return
	}

	_, err = cfg.handlerCreateArtistAlbums(name, artistDB.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tmpl.Execute(w, returnVals{Error: fmt.Sprintf("Error creating album: %v", err)})
		return
	}
	http.Redirect(w, r, "/admin/createArtistDisc", http.StatusFound)
}

func (cfg *ApiConfig) handlerCreateArtistAlbums(artistId string, artistDBId uuid.UUID) (database.Artist, error) {
	retrAlbums, err := api.GetAlbums(&artistId)
	if err != nil {
		return database.Artist{}, err
	}
	for _, album := range retrAlbums.Album {
		albumDB, err := cfg.Queries.CreateAlbum(context.Background(), database.CreateAlbumParams{
			Name:     album.StrAlbum,
			Genre:    album.StrGenre,
			ImgUrl:   album.StrAlbumThumb,
			ArtistID: artistDBId,
		})
		if err != nil {
			return database.Artist{}, err
		}
		_, err = cfg.handlerCreateAlbumTracks(album.IDAlbum, albumDB.ID, artistDBId)
		if err != nil {
			return database.Artist{}, err
		}
	}

	return database.Artist{}, nil

}

func (cfg *ApiConfig) handlerCreateAlbumTracks(albumId string, albumDBId uuid.UUID, artistDBId uuid.UUID) (database.Track, error) {
	retrAlbumTracks, err := api.GetAlbumSongs(&albumId)
	if err != nil {
		return database.Track{}, err
	}
	for _, track := range retrAlbumTracks.Track {
		trackDuration, err := strconv.Atoi(track.IntDuration)
		if err != nil {
			return database.Track{}, err
		}
		trackNumber, err := strconv.Atoi(track.IntTrackNumber)
		if err != nil {
			return database.Track{}, err
		}
		_, err = cfg.Queries.CreateAlbumTracks(context.Background(), database.CreateAlbumTracksParams{
			Name:             track.StrTrack,
			Duration:         int32(trackDuration),
			AlbumTrackNumber: int32(trackNumber),
			ArtistID:         artistDBId,
			AlbumID:          albumDBId,
		})
		if err != nil {
			return database.Track{}, err
		}
	}

	return database.Track{}, nil

}
