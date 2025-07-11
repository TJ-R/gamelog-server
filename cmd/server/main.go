package main

import (
	"database/sql"
	"gamelog-backend-api/internal/create"
	"gamelog-backend-api/internal/database"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


func main() {
	log.Println("INFO: Gamelog Backend Server Started...")

	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Printf("%v", err)
	}

	dbQueries := database.New(db)
	mux := http.NewServeMux()

	createHandler := create.NewHandler(dbQueries, os.Getenv("PLATFORM"), os.Getenv("SECRET"))
	createHandler.RegisterRoutes(mux)

	server := http.Server {
		Handler: mux,
		Addr: ":8080",
	}

	server.ListenAndServe()
}
