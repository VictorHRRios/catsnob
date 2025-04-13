package handlers

import (
	"context"
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/VictorHRRios/catsnob/internal/api"
	"github.com/VictorHRRios/catsnob/internal/database"
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

	_, err = cfg.Queries.CreateArtist(context.Background(), database.CreateArtistParams{
		FormedAt:  artist.IntFormedYear,
		Name:      artist.StrArtist,
		Biography: sql.NullString{String: *artist.StrBiographyEN, Valid: validArtistBio},
		Genre:     artist.StrGenre,
		ImgUrl:    artist.StrArtistThumb,
	})
	if err != nil {
		log.Print(err)
		http.Error(w, "Error creating artist", http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/admin/createArtistDisc", http.StatusFound)
}
