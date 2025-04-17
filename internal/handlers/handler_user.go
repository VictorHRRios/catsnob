package handlers

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/VictorHRRios/catsnob/internal/database"
)

func (cfg *ApiConfig) HandlerJoin(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("templates", "user", "register.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func (cfg *ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	_, err := cfg.Queries.CreateUser(context.Background(), database.CreateUserParams{
		Name:   name,
		ImgUrl: "https://go.dev/doc/gopher/runningsquare.jpg",
	})
	if err != nil {
		log.Print(err)
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/app/home", http.StatusFound)
}

func (cfg *ApiConfig) HandlerUserProfile(w http.ResponseWriter, r *http.Request) {
	userName := r.PathValue("username")
	user, err := cfg.Queries.GetUser(context.Background(), userName)
	if err != nil {
		log.Print(err)
		http.Error(w, "Error fetching user", http.StatusInternalServerError)
		return
	}

	tmplPath := filepath.Join("templates", "user", "profile.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
	}

	data := struct {
		Stylesheet *string
		Name       string
		Img        string
	}{
		Stylesheet: nil,
		Name:       user.Name,
		Img:        user.ImgUrl,
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Print(err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
