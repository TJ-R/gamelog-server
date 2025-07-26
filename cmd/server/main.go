package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
	"github.com/TJ-R/gamelog-server/internal/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"database/sql"
	"github.com/TJ-R/gamelog-server/internal/models"
)

func run(ctx context.Context, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)	
	defer cancel()

	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening db connection")
	}
	
	// start the server
	cfg := models.Config {
		DB: db,
		Platform: os.Getenv("PLATFORM"),
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
	}
	srv := server.NewServer(&cfg)

	httpServer := &http.Server{
		Addr: net.JoinHostPort(cfg.Host, cfg.Port),
		Handler: srv,
	}

	go func() {
		log.Printf("listening on %s\n", httpServer.Addr)

		err := httpServer.ListenAndServe()

		if err != nil && err != http.ErrServerClosed{
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10 * time.Second)
		defer cancel()
		err := httpServer.Shutdown(shutdownCtx)

		if err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()
	return nil


}

func main() {
	ctx := context.Background()
	err := run(ctx, nil)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
