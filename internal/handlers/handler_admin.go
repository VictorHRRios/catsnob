package handlers

import (
	"context"
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/VictorHRRios/catsnob/internal/api"
	"github.com/VictorHRRios/catsnob/internal/database"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) HandlerFormArtistDisc(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("templates", "registerArtist.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func (cfg *ApiConfig) HandlerCreateArtistDisc(w http.ResponseWriter, r *http.Request) {
	var validArtistBio bool

	name := r.FormValue("artist_id")
	retrArtist, err := api.GetArtist(&name)
	if err != nil {
		http.Error(w, "Error searching for artist in api", http.StatusInternalServerError)
	}
	artist := retrArtist.Artists[0]

	if artist.StrBiographyEN == nil {
		validArtistBio = false
	} else {
		validArtistBio = true
	}

	nameSlug := strings.ReplaceAll(strings.ToLower(artist.StrArtist), " ", "_")

	artistDB, err := cfg.Queries.CreateArtist(context.Background(), database.CreateArtistParams{
		FormedAt:  artist.IntFormedYear,
		Name:      artist.StrArtist,
		NameSlug:  nameSlug,
		Biography: sql.NullString{String: *artist.StrBiographyEN, Valid: validArtistBio},
		Genre:     artist.StrGenre,
		ImgUrl:    artist.StrArtistThumb,
	})
	if err != nil {
		log.Print(err)
		http.Error(w, "Error creating artist", http.StatusInternalServerError)
	}

	cfg.handlerCreateArtistAlbums(name, artistDB.ID)
	http.Redirect(w, r, "/admin/createArtistDisc", http.StatusFound)
}

func (cfg *ApiConfig) handlerCreateArtistAlbums(artistId string, artistDBId uuid.UUID) (database.Artist, error) {
	retrAlbums, err := api.GetAlbums(&artistId)
	if err != nil {
		return database.Artist{}, err
	}
	for _, album := range retrAlbums.Album {
		nameSlug := strings.ReplaceAll(strings.ToLower(album.StrAlbum), " ", "_")
		albumDB, err := cfg.Queries.CreateAlbum(context.Background(), database.CreateAlbumParams{
			Name:     album.StrAlbum,
			NameSlug: nameSlug,
			Genre:    album.StrGenre,
			ImgUrl:   album.StrAlbumThumb,
			ArtistID: artistDBId,
		})
		if err != nil {
			return database.Artist{}, err
		}
		cfg.handlerCreateAlbumTracks(album.IDAlbum, albumDB.ID, artistDBId)
	}

	return database.Artist{}, nil

}

func (cfg *ApiConfig) handlerCreateAlbumTracks(albumId string, albumDBId uuid.UUID, artistDBId uuid.UUID) (database.Track, error) {
	retrAlbumTracks, err := api.GetAlbumSongs(&albumId)
	if err != nil {
		return database.Track{}, err
	}
	for _, track := range retrAlbumTracks.Track {
		nameSlug := strings.ReplaceAll(strings.ToLower(track.StrTrack), " ", "_")
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
			NameSlug:         nameSlug,
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
