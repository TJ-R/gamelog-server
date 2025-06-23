package main

import 	_ "github.com/lib/pq"

import (
	"log"
	"net/http"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	platform string
	secret   string
}

func main() {
	log.Println("INFO: Gamelog Backend Server Started...")

	godotenv.Load()
	//dbURL := os.Getenv("DB_URL")
	//db, err := sql.Open("postgres", dbURL)
	//if err != nil {
	//	log.Printf("%v", err)
	//}

	// TODO connect DB to query
	// TODO Setup Api Config

	mux := http.NewServeMux()

	server := http.Server {
		Handler: mux,
		Addr: ":8080",
	}

	server.ListenAndServe()
}
