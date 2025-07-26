package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/TJ-R/gamelog-server/internal/auth"
	"github.com/TJ-R/gamelog-server/internal/database"
	"github.com/TJ-R/gamelog-server/internal/models"
	"github.com/TJ-R/gamelog-server/internal/util/restutil"
)





func handleCreateUsers(cfg models.Config) http.Handler {
	// Space to do prep

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			decoder := json.NewDecoder(r.Body)
			newUserReq := models.UserRequest{}

			err := decoder.Decode(&newUserReq)
			if err != nil {
				log.Println(err)
				http.Error(w, "Error when decoding", http.StatusInternalServerError)
				return 
			}

			hashedPassword, err := auth.HashPassword(newUserReq.Password)
			if err != nil {
				log.Println(err)
				http.Error(w, "Error when hashing password", http.StatusInternalServerError)
				return
			}

			db := database.New(cfg.DB)

			user, err := db.CreateUser(context.Background(), database.CreateUserParams{
				Email: newUserReq.Email,
				HashedPassword: hashedPassword,
			})
			if err != nil {
				log.Println(err)
				http.Error(w, "Error when creating user", http.StatusInternalServerError)
				return
			}

			newUser := models.User {
				ID: user.ID,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
				Email: user.Email,
			}

			restutil.RespondWithJSON(w, http.StatusOK, newUser)
		},
	)
}
