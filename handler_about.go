package main

import (
	"html/template"
	"net/http"

	"github.com/VictorHRRios/catsnob/internal/database"
)

func handleWeAre(w http.ResponseWriter, r *http.Request, u *database.User) {
	type returnVals struct {
		Error      string
		Stylesheet *string
		User       *database.User
	}
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/about/we_are.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, returnVals{User: u})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleThisIs(w http.ResponseWriter, r *http.Request, u *database.User) {
	type returnVals struct {
		Error      string
		Stylesheet *string
		User       *database.User
	}
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/about/this_is.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, returnVals{User: u})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
