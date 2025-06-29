package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gamelog-backend-api/internal/database"
)

type apiConfig struct {
	dbQueries *database.Queries
	platform  string
	secret    string
}

func main() {
	log.Println("INFO: Gamelog Backend Server Started...")

	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Printf("%v", err)
	}

	dbQueries := database.New(db)

	apiCfg := apiConfig {
		dbQueries: dbQueries,
		platform: os.Getenv("PLATFORM"),
		secret: os.Getenv("SECRET"),
	}

	// TODO connect DB to query
	// TODO Setup Api Config

	mux := http.NewServeMux()
	mux := http.HandleFunc("POST /api/create-user", apiCfg.handlerUsersCreate)

	server := http.Server {
		Handler: mux,
		Addr: ":8080",
	}

	server.ListenAndServe()
}
