package server

import (
	"net/http"

	"github.com/TJ-R/gamelog-server/internal/models"
)

func addRoutes(
	mux *http.ServeMux,
	cfg models.Config,
) {
	// Put Handlers here
	
	mux.Handle("POST /users", handleCreateUsers(cfg))
	mux.Handle("POST /admin/reset", handleReset(cfg))
}
