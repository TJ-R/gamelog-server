package admin

import (
	"context"
	"net/http"
	"github.com/TJ-R/gamelog-backend/internal/util/restutil"
)

func (h *handler) handleReset(w http.ResponseWriter, req *http.Request) {
	if h.platform != "dev" {
		http.Error(w, "Not Authorized", 403)
		return
	}

	err := h.dbQueries.DeleteUsers(context.Background())

	if err != nil {
		http.Error(w, "Error when deleting users", http.StatusInternalServerError)
	}

	restutil.RespondWithJSON(w, http.StatusOK, "OK")	

}
