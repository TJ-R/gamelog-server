package server

import (
	"net/http"
	"github.com/TJ-R/gamelog-server/internal/models"
)


func NewServer(
	config *models.Config,
) http.Handler {
	mux := http.NewServeMux()
	addRoutes(
		mux,
		*config,
	)

	var handler http.Handler = mux

	// Could add middleware on handler here
	// example middleware to handler CORS / auth middleware
	return handler
}
