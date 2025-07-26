package models

import (
	"database/sql"
)

type Config struct {
	DB *sql.DB 
	Platform string	
	Port string
	Host string
}
