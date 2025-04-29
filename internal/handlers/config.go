package handlers

import "github.com/VictorHRRios/catsnob/internal/database"

// Esta estructura se usara para hacer las consultas de SQL
type ApiConfig struct {
	Queries *database.Queries
	JWT     string
}
