package create

import (
	"encoding/json"
	"gamelog-backend-api/internal/database"
	"log"
	"net/http"
	"time"
	"context"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

type UserRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}



func (h *handler) createUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	newUserReq := UserRequest{}

	err := decoder.Decode(&newUserReq)
	if err != nil {
		log.Println(err)
		respondWithError(w, "Error when decoding", http.StatusInternalServerError)
		return 
	}

	hashedPassword, err := auth.HashPassword(newUserReq.Password)
	if err != nil {
		log.Println(err)
		responsdWithError(w, "Error when hashing password", http.StatusInternalServerError)
		return
	}

	user, err := h.dbQueries.CreateUser(context.Background(), database.CreateUserParams{
		Email: newUserReq.Email,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		log.Println(err)
		respondWithError(w, "Error when creating user", http.StatusInternalServerError)
		return
	}

	newUser := User {
		ID: user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email: user.Email,
	}

	respondWithJson(w, newUser, http.StatusOK)
}
