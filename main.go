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
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)
	apiCfg := handlers.ApiConfig{
		Queries: dbQueries,
	}

	mux := http.NewServeMux() // crea un multiplexer para las requests

	// directorio para archivos estaticos "css, imagenes, ..."
	fsHandler := http.StripPrefix("/app/assets/", http.FileServer(http.Dir(assetsDirectory)))
	mux.Handle("/app/assets/", fsHandler)

	// redirecciones a pagina principal
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/app/home", http.StatusSeeOther)
	})

	mux.HandleFunc("GET /app/home", apiCfg.HandlerIndex)
	mux.HandleFunc("GET /app/join", apiCfg.HandlerJoin)
	mux.HandleFunc("GET /app/user/{username}", apiCfg.HandlerUserProfile)
	//mux.HandleFunc("GET /app/artist/{artistId}", apiCfg.HandlerArtistDesc)
	mux.HandleFunc("GET /admin/createArtistDisc", apiCfg.HandlerFormArtistDisc)

	mux.HandleFunc("POST /app/join", apiCfg.HandlerCreateUser)
	mux.HandleFunc("POST /admin/createArtistDisc", apiCfg.HandlerCreateArtistDisc)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Fatal(srv.ListenAndServe())
}
