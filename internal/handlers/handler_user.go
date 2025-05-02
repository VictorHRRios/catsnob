package handlers

import (
	"context"
	"html/template"
	"net/http"
	"path/filepath"
	"time"

	"github.com/VictorHRRios/catsnob/internal/auth"
	"github.com/VictorHRRios/catsnob/internal/database"
)

func (cfg *ApiConfig) HandlerJoin(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("templates", "user", "register.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func (cfg *ApiConfig) HandlerLogin(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("templates", "user", "login.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func (cfg *ApiConfig) HandlerAuthUser(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	password := r.FormValue("password")
	user, err := cfg.Queries.GetUser(context.Background(), name)

	if err != nil {
		http.Error(w, "Error fetching user", http.StatusInternalServerError)
		return
	}

	if !auth.CheckPasswordHash(password, user.HashedPassword) {
		http.Error(w, "Error authenticating user", http.StatusUnauthorized)
		return
	}

	token, err := auth.MakeJWT(user.ID, cfg.JWT, time.Hour)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})
	http.Redirect(w, r, "/app/home", http.StatusFound)
}

func (cfg *ApiConfig) HandlerLogout(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{
		Name:     "auth_token",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   false,
	}

	http.SetCookie(w, c)

	http.Redirect(w, r, "/app/home", http.StatusFound)
}

func (cfg *ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	password := r.FormValue("password")
	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	_, err = cfg.Queries.CreateUser(context.Background(), database.CreateUserParams{
		Name:           name,
		ImgUrl:         "https://go.dev/doc/gopher/runningsquare.jpg",
		HashedPassword: hashedPassword,
	})
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/app/login", http.StatusFound)
}

func (cfg *ApiConfig) HandlerUserProfile(w http.ResponseWriter, r *http.Request, u *database.User) {
	userName := r.PathValue("username")
	user, err := cfg.Queries.GetUser(context.Background(), userName)
	if err != nil {
		http.Error(w, "Error fetching user", http.StatusInternalServerError)
		return
	}

	tmplPath := filepath.Join("templates", "user", "profile.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	data := struct {
		Stylesheet *string
		Name       string
		Img        string
		User       *database.User
	}{
		Stylesheet: nil,
		Name:       user.Name,
		Img:        user.ImgUrl,
		User:       u,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
