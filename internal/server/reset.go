package server

import (
	"context"
	"net/http"

	"github.com/TJ-R/gamelog-server/internal/database"
	"github.com/TJ-R/gamelog-server/internal/models"
	"github.com/TJ-R/gamelog-server/internal/util/restutil"
)

func handleReset(cfg models.Config) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			if cfg.Platform != "dev" {
				http.Error(w, "Not Authorized", 403)
				return
			}

			db := database.New(cfg.DB)

			err := db.DeleteUsers(context.Background())

			if err != nil {
				http.Error(w, "Error when deleting users", http.StatusInternalServerError)
			}

			restutil.RespondWithJSON(w, http.StatusOK, "OK")	

		},
	)
}
