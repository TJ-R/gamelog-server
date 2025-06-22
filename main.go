package main

import _ "github.com/lib/pq"

import (
	"log"
	"net/http"
)

type apiConfig struct {
	platform string
	secret   string
}

func main() {
	log.Println("INFO: Gamelog Backend Server Started...")

	// TODO Load ENV variables
	// TODO connect DB to query
	// TODO Setup Api Config

	mux := http.NewServeMux()

	server := http.Server {
		Handler: mux,
		Addr: ":8080",
	}

	server.ListenAndServe()
}
