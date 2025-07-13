package create

import (
	"github.com/TJ-R/gamelog-backend/internal/database"
	"net/http"
)

type handler struct {
	dbQueries *database.Queries
	platform  string
	secret    string
}

func NewHandler(dbQueries *database.Queries, platform, secret string) *handler {
	return &handler{
		dbQueries: dbQueries,
		platform:  platform,
		secret:    secret,
	}
}

func (h *handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /users", h.createUser)
}
