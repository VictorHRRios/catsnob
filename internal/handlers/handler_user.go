package handlers

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/VictorHRRios/catsnob/internal/auth"
	"github.com/VictorHRRios/catsnob/internal/database"
)

func passwordChecker(s string) error {
	if len(s) < 4 {
		return fmt.Errorf("Password must be at least 4 characters long\n")
	}
	return nil
}

func HandlerJoin(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("templates", "user", "register.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "error parsing files", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "error rendering template", http.StatusInternalServerError)
		return
	}
}

func HandlerLogin(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("templates", "user", "login.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "error parsing files", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "error rendering template", http.StatusInternalServerError)
		return
	}
}

func (cfg *ApiConfig) HandlerAuthUser(w http.ResponseWriter, r *http.Request) {
	type returnVals struct {
		Error      string
		Stylesheet *string
		User       *database.User
	}
	tmplPath := filepath.Join("templates", "user", "login.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "error parsing files", http.StatusInternalServerError)
		return
	}
	name := r.FormValue("name")
	password := r.FormValue("password")
	user, err := cfg.Queries.GetUser(context.Background(), name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := tmpl.Execute(w, returnVals{Error: "User does not exist"})
		if err != nil {
			http.Error(w, "error rendering template", http.StatusInternalServerError)
			return
		}
		return
	}

	if !auth.CheckPasswordHash(password, user.HashedPassword) {
		w.WriteHeader(http.StatusInternalServerError)
		err := tmpl.Execute(w, returnVals{Error: "Error authenticating user"})
		if err != nil {
			http.Error(w, "error rendering template", http.StatusInternalServerError)
			return
		}
		return
	}

	token, err := auth.MakeJWT(user.ID, cfg.JWT, time.Hour)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := tmpl.Execute(w, returnVals{Error: "Error creating JWT"})
		if err != nil {
			http.Error(w, "error rendering template", http.StatusInternalServerError)
			return
		}
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
	type returnVals struct {
		Error      string
		Stylesheet *string
		User       *database.User
	}
	tmplPath := filepath.Join("templates", "user", "register.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		http.Error(w, "error parsing files", http.StatusInternalServerError)
		return
	}
	name := r.FormValue("name")
	password := r.FormValue("password")
	if err := passwordChecker(password); err != nil {
		if err := tmpl.Execute(w, returnVals{Error: err.Error()}); err != nil {
			http.Error(w, "error rendering template", http.StatusInternalServerError)
			return
		}
		return
	}

	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := tmpl.Execute(w, returnVals{Error: err.Error()}); err != nil {
			http.Error(w, "error rendering template", http.StatusInternalServerError)
			return
		}
		return
	}
	_, err = cfg.Queries.CreateUser(context.Background(), database.CreateUserParams{
		Name:           name,
		ImgUrl:         "/app/assets/images/profile.jpg",
		HashedPassword: hashedPassword,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := tmpl.Execute(w, returnVals{Error: fmt.Sprintf("Could not create user: %v", err)}); err != nil {
			http.Error(w, "error rendering template", http.StatusInternalServerError)
			return
		}
		return
	}

	http.Redirect(w, r, "/app/login", http.StatusFound)
}

func (cfg *ApiConfig) HandlerUserProfile(w http.ResponseWriter, r *http.Request, u *database.User) {
	type returnVals struct {
		Error       string
		Stylesheet  *string
		Name        string
		Img         string
		User        *database.User
		ProfileUser *database.User
		Reviews     []database.GetReviewByUserRow
	}
	tmplPath := filepath.Join("templates", "user", "profile.html")
	tmpl, err := template.ParseFiles(layout, tmplPath)
	if err != nil {
		log.Print(err)
		http.Error(w, "error parsing files", http.StatusInternalServerError)
		return
	}
	userName := r.PathValue("username")
	user, err := cfg.Queries.GetUser(context.Background(), userName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := tmpl.Execute(w, returnVals{Error: fmt.Sprintf("Could not fetch user: %v", err)})
		if err != nil {
			http.Error(w, "error rendering template", http.StatusInternalServerError)
			return
		}
		return
	}

	reviews, err := cfg.Queries.GetReviewByUser(context.Background(), user.ID)

	returnBody := returnVals{
		ProfileUser: &user,
		Reviews:     reviews,
		User:        u,
	}

	if err := tmpl.Execute(w, returnBody); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
