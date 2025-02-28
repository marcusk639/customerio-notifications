// Package repository provides database access layer implementations
package repositories

import (
	"github.com/jmoiron/sqlx"
)

// Repository provides common functionality for all repositories
type Repository struct {
	db *sqlx.DB
}
