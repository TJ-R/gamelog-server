package restutil

import (
	"encoding/json"
	"net/http"
	"log"
)

type ErrorResponse struct {
	Message string `json:"message"` 
	Code    int    `json:"code"`
}

func RespondWithJSON (w http.ResponseWriter, code int, payload any) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error: Encountered an error when marshalling json")
		errorRes := ErrorResponse{Message: "Internal server error", Code: http.StatusInternalServerError}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(errorRes.Code)
		json.NewEncoder(w).Encode(errorRes)
		return
	}

	w.Header().Set("Content-Type", "applicaiton/json")
	w.WriteHeader(code)
	w.Write(dat)
}
