package main

import (
	"github.com/google/uuid"
	"net/http"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}

// TODO Fix Reference to apiConfig so that the struct is usable in this file
// Not sure if it is an import / package issue since if this file is in the root
// This issue does not occur
// Look up some golang api endpoint structure 

func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {

}
