package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/VictorHRRios/catsnob/internal/auth"
	"github.com/VictorHRRios/catsnob/internal/database"
)

type contextKey string

const userIDKey contextKey = "userID"

func (cfg ApiConfig) AuthMiddleware(next func(w http.ResponseWriter, r *http.Request, u *database.User)) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth_token")
		if err != nil {
			next(w, r, nil)
			log.Print(err)
			return
		}
		userID, err := auth.ValidateJWT(cookie.Value, cfg.JWT)
		if err != nil {
			next(w, r, nil)
			log.Print(err)
			return
		}
		user, err := cfg.Queries.GetUserFromID(context.Background(), userID)
		if err != nil {
			log.Print(err)
			next(w, r, nil)
		}
		next(w, r, &user)
	})
}
