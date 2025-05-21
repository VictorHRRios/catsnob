package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/VictorHRRios/catsnob/internal/database"
	"github.com/VictorHRRios/catsnob/internal/handlers"
	"github.com/joho/godotenv" // Libreria para que funcione el .dotenv
	_ "github.com/lib/pq"      // Libreria para funcionalidad de base de datos
)

//
//	CATSNOB: Por VictorHRRios, LenikaMon, Cesar, Silvia
//

func main() {
	const assetsDirectory = "assets"
	const port = "8080" // Puerto local
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	jwt := os.Getenv("JWT")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)
	apiCfg := handlers.ApiConfig{
		Queries: dbQueries,
		JWT:     jwt,
	}

	mux := http.NewServeMux() // crea un multiplexer para las requests

	// directorio para archivos estaticos "css, imagenes, ..."
	fsHandler := http.StripPrefix("/app/assets/", http.FileServer(http.Dir(assetsDirectory)))
	mux.Handle("/app/assets/", fsHandler)

	// redirecciones a pagina principal
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/app/home", http.StatusSeeOther)
	})

	mux.HandleFunc("GET /app/we_are", apiCfg.AuthMiddleware(handleWeAre))
	mux.HandleFunc("GET /app/this_is", apiCfg.AuthMiddleware(handleThisIs))

	mux.HandleFunc("GET /app/home", apiCfg.AuthMiddleware(apiCfg.HandlerIndex))
	mux.HandleFunc("GET /app/home/albums", apiCfg.AuthMiddleware(apiCfg.HandlerAlbums))
	mux.HandleFunc("GET /app/home/tracks", apiCfg.AuthMiddleware(apiCfg.HandlerTracks))
	mux.HandleFunc("GET /app/home/lists", apiCfg.AuthMiddleware(apiCfg.HandlerLists))
	mux.HandleFunc("GET /app/join", handlers.HandlerJoin)
	mux.HandleFunc("GET /app/login", handlers.HandlerLogin)
	mux.HandleFunc("GET /app/user/{username}", apiCfg.AuthMiddleware(apiCfg.HandlerUserProfile))
	mux.HandleFunc("GET /app/user/{username}/review/{reviewid}", apiCfg.AuthMiddleware(apiCfg.HandlerUserReview))
	mux.HandleFunc("GET /app/artist/{artistid}", apiCfg.AuthMiddleware(apiCfg.HandlerArtistProfile))
	mux.HandleFunc("GET /app/album/{albumid}", apiCfg.AuthMiddleware(apiCfg.HandlerAlbum))
	mux.HandleFunc("GET /app/track/{trackid}", apiCfg.AuthMiddleware(apiCfg.HandlerTrack))
	mux.HandleFunc("GET /app/lists/create_list", apiCfg.AuthMiddleware(apiCfg.HandlerCreate_List))
	mux.HandleFunc("GET /app/lists/edit_list/{listid}", apiCfg.AuthMiddleware(apiCfg.HandlerEdit_List))
	mux.HandleFunc("GET /app/lists/add_albums/{listid}", apiCfg.AuthMiddleware(apiCfg.HandlerAdd_Items))
	mux.HandleFunc("GET /app/lists/delete_albums/{listid}", apiCfg.AuthMiddleware(apiCfg.HandlerDelete_Albums))
	mux.HandleFunc("GET /app/delete_list/{listid}", apiCfg.AuthMiddleware(apiCfg.HandlerDeleteList))
	mux.HandleFunc("GET /app/lists/create_list_tracks", apiCfg.AuthMiddleware(apiCfg.HandlerCreate_List))
	mux.HandleFunc("GET /app/lists/edit_list_tracks/{listid}", apiCfg.AuthMiddleware(apiCfg.HandlerEdit_List))
	mux.HandleFunc("GET /app/lists/add_tracks/{listid}", apiCfg.AuthMiddleware(apiCfg.HandlerAdd_Items))
	mux.HandleFunc("GET /app/lists/delete_tracks/{listid}", apiCfg.AuthMiddleware(apiCfg.HandlerDelete_Tracks))

	mux.HandleFunc("GET /admin/createArtistDisc", apiCfg.AuthMiddleware(apiCfg.HandlerFormArtistDisc))
	mux.HandleFunc("GET /admin", apiCfg.AuthMiddleware(apiCfg.HandlerAdminIndex))

	mux.HandleFunc("PUT /app/updateReview", apiCfg.AuthMiddleware(apiCfg.HandlerUpdateReview))
	mux.HandleFunc("PUT /app/updateShout", apiCfg.AuthMiddleware(apiCfg.HandlerUpdateShout))

	mux.HandleFunc("DELETE /app/deleteReview", apiCfg.AuthMiddleware(apiCfg.HandlerDeleteReview))
	mux.HandleFunc("DELETE /admin/deleteArtist", apiCfg.AuthMiddleware(apiCfg.HandlerDeleteArtist))
	mux.HandleFunc("DELETE /app/deleteShout", apiCfg.AuthMiddleware(apiCfg.HandlerDeleteShout))

	mux.HandleFunc("POST /app/createReview", apiCfg.AuthMiddleware(apiCfg.HandlerCreateReview))
	mux.HandleFunc("POST /app/createReviewLong", apiCfg.AuthMiddleware(apiCfg.HandlerCreateReviewLong))
	mux.HandleFunc("POST /app/join", apiCfg.HandlerCreateUser)
	mux.HandleFunc("POST /app/login", apiCfg.HandlerAuthUser)
	mux.HandleFunc("POST /app/logout", apiCfg.HandlerLogout)
	mux.HandleFunc("POST /admin/createArtistDisc", apiCfg.AuthMiddleware(apiCfg.HandlerCreateArtistDisc))
	mux.HandleFunc("POST /app/add_albums/{listid}", apiCfg.AuthMiddleware(apiCfg.HandlerAddItemsToList))
	mux.HandleFunc("POST /app/add_tracks/{listid}", apiCfg.AuthMiddleware(apiCfg.HandlerAddItemsToList))
	mux.HandleFunc("POST /app/delete_tracks/{listid}", apiCfg.AuthMiddleware(apiCfg.HandlerDeleteTracksFromList))
	mux.HandleFunc("POST /app/delete_albums/{listid}", apiCfg.AuthMiddleware(apiCfg.HandlerDeleteAlbumsFromList))
	mux.HandleFunc("POST /app/createAlbumList", apiCfg.AuthMiddleware(apiCfg.HandlerCreateList))
	mux.HandleFunc("POST /app/createTrackList", apiCfg.AuthMiddleware(apiCfg.HandlerCreateList))

	mux.HandleFunc("POST /app/createShout", apiCfg.AuthMiddleware(apiCfg.HandlerCreateShout))

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	log.Print("Server running on port 8080...")
	log.Fatal(srv.ListenAndServe())
}
