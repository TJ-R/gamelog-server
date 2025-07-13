package models

import (
	"github.com/TJ-R/gamelog-backend/internal/database"
)

type ApiConfig struct {
	DbQueries *database.Queries
	Platform string
	Secret string
}
