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
	tmplPath := filepath.Join("templates", "register.html")
	tmpl, err := template.ParseFiles(tmplPath)
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
