package repository

import (
	"backend/internal/models"
	"database/sql"
)

//we Creted A Repository Interfcae and To Satisfy the Interface
type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies() ([]*models.Movie, error)
}