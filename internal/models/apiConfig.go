package models

import (
	"gamelog-backend-api/internal/database"
)

type ApiConfig struct {
	DbQueries *database.Queries
	Platform string
	Secret string
}
